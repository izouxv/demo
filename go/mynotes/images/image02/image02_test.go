package image02

import (
	"testing"
	"github.com/disintegration/imaging"
	"log"
	"os"
	"fmt"
)

//imaging包对图片的处理
func Test_imaging01(t *testing.T) {
	//打开一张图片
	src, err := imaging.Open("./imagepkg/branches.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	//以中心为锚点裁剪成300*300
	//dst := imaging.CropAnchor(src, 300, 300, imaging.Center)

	//将裁剪后的图像调整为宽度=200，保留纵横比
	//dst := imaging.Resize(src, 200, 0, imaging.Lanczos)

	//创建图像的模糊
	dst := imaging.Blur(src, 3.5)

	//创建具有更高对比度和锐度的图像的灰度
	//dst2 := imaging.Grayscale(src)//灰度
	//dst2 = imaging.AdjustContrast(src, 20)//对比度
	//dst2 = imaging.Sharpen(src, 2)//锐度

	//创建映像的反转
	//dst3 := imaging.Invert(src)

	//使用卷积过滤器创建图像的浮雕
	//dst4 := imaging.Convolve3x3(
	//	src, [9]float64{
	//		-1, -1, 0,
	//		-1, 1, 1,
	//		0, 1, 1,},nil, )

	//创建一个新的图像并粘贴四个生成的图像到它
	//dst := imaging.New(1200, 800, color.NRGBA{0, 0, 0, 0})
	//dst = imaging.Paste(dst, dst1, image.Pt(0, 0))
	//dst = imaging.Paste(dst, dst2, image.Pt(0, 400))
	//dst = imaging.Paste(dst, dst3, image.Pt(600, 0))
	//dst = imaging.Paste(dst, dst4, image.Pt(600, 400))

	//将结果图像保存为JPEG
	err = imaging.Save(dst, "example11.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func Test_imaging02(t *testing.T) {
	//打开一张图片
	src, err := imaging.Open("./imagepkg/branches.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	//调整尺寸并保持纵横比
	dst := imaging.Resize(src,200,200,imaging.Lanczos)
	//调整成为适合的尺寸
	//dst := imaging.Fit(src,200,200,imaging.Lanczos)
	//调整尺寸并充填
	//dst := imaging.Fill(src,200,200,imaging.Center,imaging.Lanczos)
	//NearestNeighbor、Linear、CatmullRom、Lanczos、Box、MitchellNetravali、CatmullRom、Gaussian、
	//将结果图像保存为JPEG
	err = imaging.Save(dst, "example11.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func Test_imaging03(t *testing.T) {
	file,err := os.Open("./imagepkg/12.png")
	defer file.Close()
	src, err := imaging.Decode(file)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	x := src.Bounds().Size().X
	y := src.Bounds().Size().Y
	fmt.Println(x,y)
	//调整尺寸并保持纵横比
	dst := imaging.Resize(src,x,y,imaging.Lanczos)
	//调整成为适合的尺寸
	//dst := imaging.Fit(src,200,200,imaging.Lanczos)
	//调整尺寸并充填
	//dst := imaging.Fill(src,200,200,imaging.Center,imaging.Lanczos)
	//NearestNeighbor、Linear、CatmullRom、Lanczos、Box、MitchellNetravali、CatmullRom、Gaussian、
	//将结果图像保存为JPEG
	err = imaging.Save(dst, "test3.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}
