package paint

import (
	"image/color"
	"image/draw"
	"math"
)

// 常用的色彩
var (
	Black  = color.RGBA{0, 0, 0, 255}
	Red    = color.RGBA{255, 0, 0, 255}
	Green  = color.RGBA{0, 255, 0, 255}
	Blue   = color.RGBA{0, 0, 255, 255}
	Yellow = color.RGBA{255, 255, 0, 255}
	Purple = color.RGBA{255, 0, 255, 255}
	Ching  = color.RGBA{0, 255, 255, 255}
	White  = color.RGBA{255, 255, 255, 255}
)

// 表示一个可以被绘制的图像
// 因为此图像内部维护的状态，以及图像本身的限制，不支持并发
type Image struct {
	draw.Image             // 画布
	FR         color.Color // 前景色
	BG         color.Color // 背景色
}

func max(x float64, r ...float64) float64 {
	for _, o := range r {
		if x < o {
			x = o
		}
	}
	return x
}

func min(x float64, r ...float64) float64 {
	for _, o := range r {
		if x > o {
			x = o
		}
	}
	return x
}

func round(x float64) int {
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
