package image02

import (
	"testing"
	"os"
	"image/jpeg"
	"github.com/nfnt/resize"
	"fmt"
	"image"
)

//resize包对图片的处理
func Test_resize01(t *testing.T) {
	file, err := os.Open("./imagepkg/13.png")
	if err != nil {
		fmt.Println("Open:",err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Decode error:",err)
	}
	defer file.Close()
	x := img.Bounds().Size().X
	y := img.Bounds().Size().Y
	fmt.Println(x,y,uint(x),uint(y))
	m := resize.Thumbnail(uint(x), uint(y), img, resize.Bilinear)
	out, err := os.Create("test1.png")
	if err != nil {
		fmt.Println("Create:",err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
}

func Test_resize02(t *testing.T) {
	file,err := os.Open("./imagepkg/13.png")
	defer file.Close()
	src, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("failed to open image: ", err)
	}
	x := src.Bounds().Size().X
	y := src.Bounds().Size().Y
	fmt.Println(x,y)
	m := resize.Thumbnail(uint(x), uint(y), src, resize.Bilinear)
	out, err := os.Create("test2.png")
	if err != nil {
		fmt.Println("Create:",err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
}