package paint

import (
	"errors"
	"image"
	"math"
)

// 绘制贝塞尔曲线，根据提供的校准点绘制，num为采样点个数
func (img *Image) Bezier(num int, spt ...image.Point) error {
	return img.BezierSlice(num, spt)
}

// 类似BezierSlice但接受校准点的切片
func (img *Image) BezierSlice(num int, spt []image.Point) error {
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
		a, b := round(x), round(y)
		img.Line(c.X, c.Y, a, b)
		c.X, c.Y = a, b
	}
	img.Line(c.X, c.Y, spt[l-1].X, spt[l-1].Y)
	return nil
}
