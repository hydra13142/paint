package paint

import (
	"image"
	"image/color"
	"math"
)

func Project(dst Image, src image.Image, prj func(x, y float64) (x1, y1 float64)) {
	type medi struct{ r, g, b, a float64 }
	rect := dst.Bounds()
	dw := rect.Max.X - rect.Min.X
	dh := rect.Max.Y - rect.Min.Y
	c := make([][]medi, dw)
	for i := 0; i < dw; i++ {
		c[i] = make([]medi, dh)
	}
	rect = src.Bounds()
	sw := rect.Max.X - rect.Min.X
	sh := rect.Max.Y - rect.Min.Y
	for i := 0; i < sw; i++ {
		for j := 0; j < sh; j++ {
			r, g, b, a := src.At(i, j).RGBA()
			sx, sy := float64(i), float64(j)
			x0, y0 := prj(sx, sy)
			x1, y1 := prj(sx, sy+1)
			x2, y2 := prj(sx+1, sy+1)
			x3, y3 := prj(sx+1, sy)
			cv := Convex([]Point{{x0, y0}, {x1, y1}, {x2, y2}, {x3, y3}})
			lx, rx := math.Floor(Min(x0, x1, x2, x3)), math.Ceil(Max(x0, x1, x2, x3))
			by, ty := math.Floor(Min(y0, y1, y2, y3)), math.Ceil(Max(y0, y1, y2, y3))
			if lx >= float64(dw) || rx < 0 {
				continue
			}
			if by >= float64(dh) || ty < 0 {
				continue
			}
			if lx < 0 {
				lx = 0
			}
			if by < 0 {
				by = 0
			}
			if rx > float64(dw) {
				rx = float64(dw)
			}
			if ty > float64(dh) {
				ty = float64(dh)
			}
			for x := lx; x < rx; x += 1 {
				for y := by; y < ty; y += 1 {
					vc := Convex([]Point{
						{x, y},
						{x, y + 1},
						{x + 1, y + 1},
						{x + 1, y}})
					ad := And(cv, vc)
					k := ad.Area()
					if k != 0 {
						p := &c[int(x)][int(y)]
						p.r += float64(r) * k
						p.g += float64(g) * k
						p.b += float64(b) * k
						p.a += float64(a) * k
					}
				}
			}
		}
	}
	for i := 0; i < dw; i++ {
		for j := 0; j < dh; j++ {
			p := &c[i][j]
			dst.Set(i, j, color.RGBA64{uint16(p.r), uint16(p.g), uint16(p.b), uint16(p.a)})
		}
	}
}