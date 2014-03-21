package main

import (
	"github.com/hydra13142/paint"
	"image"
	"image/png"
	"os"
)

func main() {
	canvas := image.NewRGBA(image.Rect(0, 0, 400, 400))
	paint.Rect(canvas, 50, 50, 349, 349, image.Black)
	paint.Line(canvas, 50, 50, 349, 349, image.Black)
	paint.Line(canvas, 50, 349, 349, 50, image.Black)
	paint.Ellipse(canvas, 100, 100, 299, 299, image.Black)
	file, _ := os.Create("circle.png")
	defer file.Close()
	png.Encode(file, canvas)
}
