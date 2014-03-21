package main

import (
	"fmt"
	"github.com/hydra13142/paint"
	"image"
	_ "image/jpeg"
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
	dst := image.NewRGBA(image.Rect(0, 0, 262, 500))
	err = paint.Resize(dst, img)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err = os.Create("C:/Users/LiuZhiXi/Downloads/42297164_3.png")
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
