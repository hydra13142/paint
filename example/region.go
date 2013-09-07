package main

import "github.com/hydra13142/paint"

func main() {
	canvas := paint.NewCanvas(400, 280)
	canvas.Rect(100, 10, 299, 29)
	canvas.Region(100, 40, 299, 59, 1, 5)
	canvas.Region(100, 70, 299, 89, 2, 5)
	canvas.Region(100, 100, 299, 119, 4, 5)
	canvas.Region(100, 130, 299, 149, 8, 5)
	canvas.Region(100, 160, 299, 179, 3, 5)
	canvas.Region(100, 190, 299, 209, 12, 5)
	canvas.Region(100, 220, 299, 239, 15, 5)
	canvas.Block(100, 250, 299, 269)
	canvas.Save("region.png")
}
