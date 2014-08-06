package paint

import (
	"errors"
	"github.com/hydra13142/geom"
	"image"
	"image/color"
	"image/draw"
	"math"
)

// 将图像按照prj定义的映射函数进行映射，一般用于随意角度选择、哈哈镜等效果，速度慢
// prj函数的映射必须是不改变平面关系的，即凸多边形的各点的映射（顺序不变）仍是凸多边形
func Project(dst draw.Image, src image.Image, prj func(x, y float64) (x1, y1 float64)) error {
	type medi struct{ r, g, b, a float64 }
	sr := dst.Bounds()
	dr := dst.Bounds()
	dw := sr.Max.X - sr.Min.X
	dh := sr.Max.Y - sr.Min.Y
	if dw <= 0 || dh <= 0 {
		return errors.New("source image is empty or noncanonical")
	}
	if dr.Min.X >= dr.Max.X || dr.Min.Y >= dr.Max.Y {
		return errors.New("target image is empty or noncanonical")
	}
	c := make([][]medi, dw)
	for i := 0; i < dw; i++ {
		c[i] = make([]medi, dh)
	}
	sr = src.Bounds()
	sw := sr.Max.X - sr.Min.X
	sh := sr.Max.Y - sr.Min.Y
	for i := 0; i < sw; i++ {
		for j := 0; j < sh; j++ {
			r, g, b, a := src.At(i+sr.Min.X, j+sr.Min.Y).RGBA()
			sx, sy := float64(i), float64(j)
			x0, y0 := prj(sx, sy)
			x1, y1 := prj(sx, sy+1)
			x2, y2 := prj(sx+1, sy+1)
			x3, y3 := prj(sx+1, sy)
			cv := geom.UnsafeConvex([]geom.Point{{x0, y0}, {x1, y1}, {x2, y2}, {x3, y3}})
			lx, rx := math.Floor(min(x0, x1, x2, x3)), math.Ceil(max(x0, x1, x2, x3))
			by, ty := math.Floor(min(y0, y1, y2, y3)), math.Ceil(max(y0, y1, y2, y3))
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
					vc := geom.UnsafeConvex([]geom.Point{{x, y}, {x, y + 1}, {x + 1, y + 1}, {x + 1, y}})
					k := cv.And(vc).Area()
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
			dst.Set(dr.Min.X+i, dr.Min.Y+j, color.RGBA64{uint16(p.r), uint16(p.g), uint16(p.b), uint16(p.a)})
		}
	}
	return nil
}
