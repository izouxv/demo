package image02

import (
	"fmt"
	"os"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"testing"
)

//原生image包
func Test_imagePkg(t *testing.T) {
	f1, err := os.Open("./images/image02/imagepkg/01.jpg")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Open("./images/image02/imagepkg/02.jpg")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	f3, err := os.Create("./images/image02/imagepkg/03.jpg")
	if err != nil {
		panic(err)
	}
	defer f3.Close()
	m1, err := jpeg.Decode(f1)
	if err != nil {
		panic(err)
	}
	bounds := m1.Bounds()
	m2, err := jpeg.Decode(f2)
	if err != nil {
		panic(err)
	}
	m := image.NewRGBA(bounds)
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
	draw.Draw(m, bounds, m1, image.ZP, draw.Src)
	draw.Draw(m, image.Rect(200, 200, 400, 400), m2, image.Pt(200, 60), draw.Src)
	err = jpeg.Encode(f3, m, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}
	fmt.Printf("okn")
}