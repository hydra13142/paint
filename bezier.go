package paint

import (
	"errors"
	"image"
	"image/color"
	"math"
)

func BezierSlice(img Image, num int, clr color.Color, spt []image.Point) error {
	l := len(spt)
	if l < 2 {
		return errors.New("need at least two points")
	}
	c := spt[0]
	for i := 1; i < num; i++ {
		t := float64(i) / float64(num)
		k, v := 1, math.Pow(1-t, float64(l-1))
		x, y := 0.0, 0.0
		for j, o := range spt {
			x += v * float64(k*o.X)
			y += v * float64(k*o.Y)
			k = k * (l - j - 1) / (j + 1)
			v = v * t / (1 - t)
		}
		a, b := Round(x), Round(y)
		Line(img, c.X, c.Y, a, b, clr)
		c.X, c.Y = a, b
	}
	Line(img, c.X, c.Y, spt[l-1].X, spt[l-1].Y, clr)
	return nil
}

func Bezier(img Image, num int, clr color.Color, spt ...image.Point) error {
	return BezierSlice(img, num, clr, spt)
}
