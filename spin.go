package paint

import (
	"errors"
	"image"
	"image/draw"
)

// 用于翻转的常量
const (
	Horizontal = 0
	Vertical   = 1
)

// 用于旋转的常量
const (
	L0   = 0
	R0   = 0
	L90  = +1
	R90  = -1
	L180 = +2
	R180 = -2
	L270 = R90
	R270 = L90
)

// 翻转函数，要求两个图像大小契合，act&1 == 0则左右翻转，否则垂直翻转。
func Overturn(dst draw.Image, src image.Image, act int) error {
	var to func(int, int) (int, int)
	sr := src.Bounds()
	dr := dst.Bounds()
	W := dr.Max.X - dr.Min.X
	H := dr.Max.Y - dr.Min.Y
	if H <= 0 || W <= 0 {
		return errors.New("target image is empty or noncanonical")
	}
	if sr.Min.X >= sr.Max.X || sr.Min.Y >= sr.Max.Y {
		return errors.New("source image is empty or noncanonical")
	}
	if sr.Max.X-sr.Min.X != W || sr.Max.Y-sr.Min.Y != H {
		return errors.New("target and source must be same size!")
	}
	if act&1 == 0 {
		to = func(x, y int) (int, int) {
			return W - 1 - x, y
		}
	} else {
		to = func(x, y int) (int, int) {
			return x, H - 1 - y
		}
	}
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			x, y := to(i, j)
			dst.Set(dr.Min.X+x, dr.Min.Y+y, src.At(sr.Min.X+i, sr.Min.Y+j))
		}
	}
	return nil
}

// 旋转，每单位1代表左旋90度。注意采用的是向右向上的直角坐标系，和图像的向右向下不同。
// 本包的函数都是采用向右向上的直角坐标系。因此在图像显示上，每单位1右旋90度。
func Rotate(dst draw.Image, src image.Image, act int) error {
	var to func(int, int) (int, int)
	sr := src.Bounds()
	dr := dst.Bounds()
	W := dr.Max.X - dr.Min.X
	H := dr.Max.Y - dr.Min.Y
	if H <= 0 || W <= 0 {
		return errors.New("target image is empty or noncanonical")
	}
	if sr.Min.X >= sr.Max.X || sr.Min.Y >= sr.Max.Y {
		return errors.New("source image is empty or noncanonical")
	}
	switch act %= 4; act {
	case 2, -2:
		if sr.Max.X-sr.Min.X != W || sr.Max.Y-sr.Min.Y != H {
			return errors.New("target and source must be same size!")
		}
		to = func(x, y int) (int, int) {
			return W - 1 - x, H - 1 - y
		}
	case 1, -3:
		if sr.Max.X-sr.Min.X != H || sr.Max.Y-sr.Min.Y != W {
			return errors.New("target and source must be same size!")
		}
		to = func(x, y int) (int, int) {
			return H - 1 - y, x
		}
	case -1, 3:
		if sr.Max.X-sr.Min.X != H || sr.Max.Y-sr.Min.Y != W {
			return errors.New("target and source must be same size!")
		}
		to = func(x, y int) (int, int) {
			return y, W - 1 - x
		}
	case 0:
		if sr.Max.X-sr.Min.X != W || sr.Max.Y-sr.Min.Y != H {
			return errors.New("target and source must be same size!")
		}
		to = func(x, y int) (int, int) {
			return x, y
		}
	}
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			x, y := to(i, j)
			dst.Set(dr.Min.X+x, dr.Min.Y+y, src.At(sr.Min.X+i, sr.Min.Y+j))
		}
	}
	return nil
}
