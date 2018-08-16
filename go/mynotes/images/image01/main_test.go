package image01

import (
	"image"
	_ "image/png"
	"os"
	"mynotes/images/image01/nfnt"
	"fmt"
	"testing"
)

func Test_Main(t *testing.T) {
	f, _ := os.Open("./examples/01.jpg")
	img, _, _ := image.Decode(f)
	analyzer := NewAnalyzer(nfnt.NewDefaultResizer())
	topCrop, _ := analyzer.FindBestCrop(img, 100, 200)
	// The crop will have the requested aspect ratio, but you need to copy/scale it yourself
	fmt.Println("Top crop:", topCrop)
	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	croppedimg := img.(SubImager).SubImage(topCrop)
	writeImage("jpeg", croppedimg, "./04.jpg")
	//fmt.Println(croppedimg)
	// ...
}
