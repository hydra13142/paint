package paint

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"math"
)

// 用于对图像的大小进行调整
func Resize(dst draw.Image, src image.Image) error {
	sr := src.Bounds()
	dr := dst.Bounds()
	if sr.Min.X >= sr.Max.X || sr.Min.Y >= sr.Max.Y {
		return errors.New("source image is empty or noncanonical")
	}
	if dr.Min.X >= dr.Max.X || dr.Min.Y >= dr.Max.Y {
		return errors.New("target image is empty or noncanonical")
	}
	kx := float64(sr.Max.X-sr.Min.X) / float64(dr.Max.X-dr.Min.X)
	ky := float64(sr.Max.Y-sr.Min.Y) / float64(dr.Max.Y-dr.Min.Y)
	for i, w := 0, dr.Max.X-dr.Min.X; i < w; i++ {
		for j, h := 0, dr.Max.Y-dr.Min.Y; j < h; j++ {
			fx := kx*float64(i) + float64(sr.Min.X)
			tx := fx + kx
			fy := ky*float64(j) + float64(sr.Min.Y)
			ty := fy + ky
			if tx <= math.Floor(fx)+1 {
				if ty <= math.Floor(fy)+1 {
					dst.Set(dr.Min.X+i, dr.Min.Y+j, src.At(int(fx), int(fy)))
				} else {
					R, G, B, A := 0.0, 0.0, 0.0, 0.0
					var r, g, b, a uint32
					var s, d float64
					d = math.Ceil(fy) - fy
					if d != 0 {
						r, g, b, a = src.At(int(fx), int(fy)).RGBA()
						s := kx * d
						R += s * float64(r)
						G += s * float64(g)
						B += s * float64(b)
						A += s * float64(a)
					}
					for fy += d; fy < ty; fy += d {
						d = 1.0
						if fy+d > ty {
							d = ty - fy
						}
						s = kx * d
						r, g, b, a = src.At(int(fx), int(fy)).RGBA()
						R += s * float64(r)
						G += s * float64(g)
						B += s * float64(b)
						A += s * float64(a)
					}
					s = kx * ky
					dst.Set(dr.Min.X+i, dr.Min.Y+j, color.RGBA64{uint16(R / s), uint16(G / s), uint16(B / s), uint16(A / s)})
				}
			} else {
				if ty <= math.Floor(fy)+1 {
					R, G, B, A := 0.0, 0.0, 0.0, 0.0
					var r, g, b, a uint32
					var s, d float64
					d = math.Ceil(fx) - fx
					if d != 0 {
						r, g, b, a = src.At(int(fx), int(fy)).RGBA()
						s := d * ky
						R += s * float64(r)
						G += s * float64(g)
						B += s * float64(b)
						A += s * float64(a)
					}
					for fx += d; fx < tx; fx += d {
						d = 1.0
						if fx+d > tx {
							d = tx - fx
						}
						s = d * ky
						r, g, b, a = src.At(int(fx), int(fy)).RGBA()
						R += s * float64(r)
						G += s * float64(g)
						B += s * float64(b)
						A += s * float64(a)
					}
					s = kx * ky
					dst.Set(dr.Min.X+i, dr.Min.Y+j, color.RGBA64{uint16(R / s), uint16(G / s), uint16(B / s), uint16(A / s)})
				} else {
					R, G, B, A := 0.0, 0.0, 0.0, 0.0
					var r, g, b, a uint32
					var s, dx, dy float64
					for px := fx; px < tx; px += dx {
						dx = math.Ceil(px) - px
						if dx == 0 {
							dx = 1
						}
						if px+dx > tx {
							dx = tx - px
						}
						for py := fy; py < ty; py += dy {
							dy = math.Ceil(py) - py
							if dy == 0 {
								dy = 1
							}
							if py+dy > ty {
								dy = ty - py
							}
							r, g, b, a = src.At(int(px), int(py)).RGBA()
							s = dx * dy
							R += s * float64(r)
							G += s * float64(g)
							B += s * float64(b)
							A += s * float64(a)
						}
					}
					s = kx * ky
					dst.Set(dr.Min.X+i, dr.Min.Y+j, color.RGBA64{uint16(R / s), uint16(G / s), uint16(B / s), uint16(A / s)})
				}
			}
		}
	}
	return nil
}
