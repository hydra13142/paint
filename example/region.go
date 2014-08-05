package main

import (
	"github.com/hydra13142/paint"
	"image"
	"image/png"
	"os"
)

func main() {
	canvas := paint.Image{image.NewRGBA(image.Rect(0, 0, 400, 400)), paint.Black, paint.White}
	canvas.Rect(100, 10, 299, 29)
	canvas.Block(100, 40, 299, 59, paint.Level, 5)
	canvas.Block(100, 70, 299, 89, paint.Plumb, 5)
	canvas.Block(100, 100, 299, 119, paint.Slant, 5)
	canvas.Block(100, 130, 299, 149, paint.Twill, 5)
	canvas.Block(100, 160, 299, 179, paint.Level|paint.Plumb, 5)
	canvas.Block(100, 190, 299, 209, paint.Slant|paint.Twill, 5)
	canvas.Block(100, 220, 299, 239, paint.Level|paint.Plumb|paint.Slant|paint.Twill, 5)
	canvas.Bar(100, 250, 299, 269)
	file, _ := os.Create("region.png")
	defer file.Close()
	png.Encode(file, canvas)
}
