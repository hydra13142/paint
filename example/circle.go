package main

import (
	"github.com/hydra13142/paint"
	"image"
	"image/png"
	"os"
)

func main() {
	canvas := paint.Image{image.NewRGBA(image.Rect(0, 0, 400, 400)), paint.Black, paint.White}
	canvas.Rect(50, 50, 349, 349)
	canvas.Line(50, 50, 349, 349)
	canvas.Line(50, 349, 349, 50)
	canvas.Ellipse(100, 100, 299, 299)
	file, _ := os.Create("circle.png")
	defer file.Close()
	png.Encode(file, canvas)
}
