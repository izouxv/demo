package options

import (
	"image"
)

// Resizer is used to resize images. See the nfnt package for a default implementation using
// github.com/nfnt/resize.
type Resizer interface {
	Resize(img image.Image, width, height uint) image.Image
}
