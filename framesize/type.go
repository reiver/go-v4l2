package v4l2_framesize

import (
	"github.com/reiver/go-v4l2/pixelformat"
)

type Type struct {
	Index        uint32                // Frame size number
	PixelFormat  v4l2_pixelformat.Type // Pixel format
	typ          uint32                // Frame size type the device supports.

	frameSize [6]uint32                // Frame size

	reserved  [2]uint32                // Reserved space for future use
}
