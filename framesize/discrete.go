package v4l2_framesize

import (
	"github.com/reiver/go-v4l2/pixelformat"
)

type Discrete struct {
	index              uint32                // Frame size number
	PixelFormat        v4l2_pixelformat.Type // Pixel format
	typ                uint32                // Frame size type the device supports.

	Width              uint32
	Height             uint32
	restOfFrameSize [4]uint32

	reserved        [2]uint32                // Reserved space for future use
}
