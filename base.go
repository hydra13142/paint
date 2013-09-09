package paint

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}

type NetType int

const (
	Level NetType = 1 << iota
	Plumb
	Slant
	Twill
)

func RGB(r, g, b uint8) color.RGBA {
	return color.RGBA{r, g, b, 255}
}

func RGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA{r, g, b, a}
}

func NewCanvas(x, y int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, x-1, y-1))
}

func LoadCanvas(name string) (Image, error) {
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
		return nil, errors.New("Cannot load png file")
	}
	return y, nil
}

func SaveCanvas(name string, img Image) error {
	f, e := os.Create(name)
	if e != nil {
		return e
	}
	defer f.Close()
	return png.Encode(f, img)
}
