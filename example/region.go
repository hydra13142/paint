package main

import (
	"github.com/hydra13142/paint"
	"image"
)

func main() {
	canvas := paint.NewCanvas(400, 280)
	paint.Rect(canvas, 100, 10, 299, 29, image.Black)
	paint.Region(canvas, 100, 40, 299, 59, paint.Level, 5, image.Black)
	paint.Region(canvas, 100, 70, 299, 89, paint.Plumb, 5, image.Black)
	paint.Region(canvas, 100, 100, 299, 119, paint.Slant, 5, image.Black)
	paint.Region(canvas, 100, 130, 299, 149, paint.Twill, 5, image.Black)
	paint.Region(canvas, 100, 160, 299, 179, paint.Level|paint.Plumb, 5, image.Black)
	paint.Region(canvas, 100, 190, 299, 209, paint.Slant|paint.Twill, 5, image.Black)
	paint.Region(canvas, 100, 220, 299, 239, paint.Level|paint.Plumb|paint.Slant|paint.Twill, 5, image.Black)
	paint.Block(canvas, 100, 250, 299, 269, image.Black)
	paint.SaveCanvas("region.png", canvas)
}
