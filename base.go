package paint

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
)

var ErrorCannotPaint = errors.New("Cannot paint image of this format")

type NetType int

type Image interface {
	image.Image
	Set(x, y int, c color.Color)
	SubImage(r image.Rectangle) image.Image
}

type Canvas struct {
	Image
	Fore color.RGBA
	Back color.RGBA
}

const (
	Level NetType = 1 << iota
	Plumb
	Slant
	Twill
)

func (c *Canvas) SetForeColor(i color.Color) {
	c.Fore = color.RGBAModel.Convert(i).(color.RGBA)
}

func (c *Canvas) SetBackColor(i color.Color) {
	c.Back = color.RGBAModel.Convert(i).(color.RGBA)
}

func (c *Canvas) PSet(x, y int) {
	c.Image.Set(x, y, c.Fore)
}

func (c *Canvas) SubCanvas(r image.Rectangle) (*Canvas, error) {
	u, ok := c.Image.SubImage(r).(Image)
	if !ok {
		return nil, ErrorCannotPaint
	}
	return &Canvas{Image: u}, nil
}

func (c *Canvas) Save(name string) error {
	f, e := os.Create(name)
	if e != nil {
		return e
	}
	defer f.Close()
	return png.Encode(f, c.Image)
}

func RGB(r, g, b uint8) color.RGBA {
	return color.RGBA{r, g, b, 255}
}

func RGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

func ToRGBA(i color.Color) color.RGBA {
	return color.RGBAModel.Convert(i).(color.RGBA)
}

func NewCanvas(x, y int) *Canvas {
	return &Canvas{image.NewRGBA(image.Rect(0, 0, x-1, y-1)), RGB(0, 0, 0), RGB(255, 255, 255)}
}

func LoadCanvas(name string) (*Canvas, error) {
	f, e := os.Open(name)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	x, _, e := image.Decode(f)
	if e != nil {
		return nil, e
	}
	y, ok := x.(Image)
	if !ok {
		return nil, ErrorCannotPaint
	}
	return &Canvas{y, RGB(0, 0, 0), RGB(255, 255, 255)}, nil
}
