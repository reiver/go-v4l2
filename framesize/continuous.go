package v4l2_framesize

import (
	"github.com/reiver/go-v4l2/pixelformat"
)

type Continuous struct {
	index              uint32                // Frame size number
	PixelFormat        v4l2_pixelformat.Type // Pixel format
	typ                uint32                // Frame size type the device supports.

	MinWidth           uint32                // Minimum frame width [pixel]
	MaxWidth           uint32                // Maximum frame width [pixel]
	StepWidth          uint32                // Frame width step size [pixel]
	MinHeight          uint32                // Minimum frame height [pixel]
	MaxHeight          uint32                // Maximum frame height [pixel]
	StepHeight         uint32                // Frame height step size [pixel]

	reserved        [2]uint32                // Reserved space for future use
}
