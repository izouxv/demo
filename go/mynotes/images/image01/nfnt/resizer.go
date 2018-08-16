package nfnt

import (
	"image"
	"mynotes/images/image01/options"
	"github.com/nfnt/resize"
)

type nfntResizer struct {
	interpolationType resize.InterpolationFunction
}

func (r nfntResizer) Resize(img image.Image, width, height uint) image.Image {
	return resize.Resize(width, height, img, r.interpolationType)
}

// NewResizer creates a new Resizer with the given interpolation type.
func NewResizer(interpolationType resize.InterpolationFunction) options.Resizer {
	return nfntResizer{interpolationType: interpolationType}
}

// NewDefaultResizer creates a new Resizer with the default interpolation type.
func NewDefaultResizer() options.Resizer {
	return NewResizer(resize.Bicubic)
}
