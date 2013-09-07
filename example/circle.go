package main

import "github.com/hydra13142/paint"

func main() {
	canvas := paint.NewCanvas(400, 400)
	canvas.Rect(50, 50, 349, 349)
	canvas.Line(50, 50, 349, 349)
	canvas.Line(50, 349, 349, 50)
	canvas.Ellipse(100, 100, 299, 299)
	canvas.Save("circle.png")
}
