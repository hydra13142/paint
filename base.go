package paint

import (
	"image"
	"math"
	"image/color"
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

func Max(x float64, r ...float64) float64 {
	for _, o := range r {
		if x < o {
			x = o
		}
	}
	return x
}

func Min(x float64, r ...float64) float64 {
	for _, o := range r {
		if x > o {
			x = o
		}
	}
	return x
}

func Round(x float64) int {
	if x >= 0 {
		if x-math.Floor(x) >= 0.5 {
			return int(x) + 1
		} else {
			return int(x)
		}
	} else {
		if x-math.Floor(x) >= 0.5 {
			return int(x)
		} else {
			return int(x) - 1
		}
	}
}
