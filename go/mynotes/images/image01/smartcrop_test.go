package image01

import (
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"mynotes/images/image01/nfnt"
)

var (
	testFile = "./examples/gopher.jpg"
)

// Moved here and unexported to decouple the resizer implementation.
func smartCrop(img image.Image, width, height int) (image.Rectangle, error) {
	analyzer := NewAnalyzer(nfnt.NewDefaultResizer())
	return analyzer.FindBestCrop(img, width, height)
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func TestCrop(t *testing.T) {
	fi, _ := os.Open(testFile)
	defer fi.Close()

	img, _, err := image.Decode(fi)
	if err != nil {
		t.Fatal(err)
	}

	topCrop, err := smartCrop(img, 250, 250)
	if err != nil {
		t.Fatal(err)
	}
	expected := image.Rect(464, 24, 719, 279)
	if topCrop != expected {
		t.Fatalf("expected %v, got %v", expected, topCrop)
	}

	sub, ok := img.(SubImager)
	if ok {
		cropImage := sub.SubImage(topCrop)
		// cropImage := sub.SubImage(image.Rect(topCrop.X, topCrop.Y, topCrop.Width+topCrop.X, topCrop.Height+topCrop.Y))
		writeImage("jpeg", cropImage, "./smartcrop.jpg")
	} else {
		t.Error(errors.New("no SubImage support"))
	}
}

func BenchmarkCrop(b *testing.B) {
	fi, err := os.Open(testFile)
	if err != nil {
		b.Fatal(err)
	}
	defer fi.Close()

	img, _, err := image.Decode(fi)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := smartCrop(img, 250, 250); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEdge(b *testing.B) {
	fi, err := os.Open(testFile)
	if err != nil {
		b.Fatal(err)
	}
	defer fi.Close()

	img, _, err := image.Decode(fi)
	if err != nil {
		b.Fatal(err)
	}

	rgbaImg := toRGBA(img)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		o := image.NewRGBA(img.Bounds())
		edgeDetect(rgbaImg, o)
	}
}

func BenchmarkImageDir(b *testing.B) {
	files, err := ioutil.ReadDir("./examples")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for _, file := range files {
		if strings.Contains(file.Name(), ".jpg") {
			fi, _ := os.Open("./examples/" + file.Name())
			defer fi.Close()

			img, _, err := image.Decode(fi)
			if err != nil {
				b.Error(err)
				continue
			}

			topCrop, err := smartCrop(img, 220, 220)
			if err != nil {
				b.Error(err)
				continue
			}
			fmt.Printf("Top crop: %+v\n", topCrop)

			sub, ok := img.(SubImager)
			if ok {
				cropImage := sub.SubImage(topCrop)
				// cropImage := sub.SubImage(image.Rect(topCrop.X, topCrop.Y, topCrop.Width+topCrop.X, topCrop.Height+topCrop.Y))
				writeImage("jpeg", cropImage, "/tmp/smartcrop/smartcrop-"+file.Name())
			} else {
				b.Error(errors.New("no SubImage support"))
			}
		}
	}
	// fmt.Println("average time/image:", b.t)
}
