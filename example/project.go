package main

import (
	"fmt"
	"github.com/hydra13142/paint"
	"image"
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/LiuZhiXi/Downloads/42297164.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	dst := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	paint.Project(dst, img, func(x, y float64) (x1, y1 float64) {
		return x + y/2, y
	})

	file, err = os.Create("C:/Users/LiuZhiXi/Downloads/42297164_2.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	err = png.Encode(file, dst)
	if err != nil {
		fmt.Println(err)
		return
	}
}
