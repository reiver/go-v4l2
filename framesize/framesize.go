package v4l2_framesize

import (
	"github.com/reiver/go-v4l2/pixelformat"

	"fmt"
	"unsafe"
)

const (
	const_FRMSIZE_TYPE_DISCRETE   = 1
	const_FRMSIZE_TYPE_CONTINUOUS = 2
	const_FRMSIZE_TYPE_STEPWISE   = 3
)

type Type struct {
	Index        uint32                // Frame size number
	PixelFormat  v4l2_pixelformat.Type // Pixel format
	typ          uint32                // Frame size type the device supports.

	frameSize [6]uint32                // Frame size

	reserved  [2]uint32                // Reserved space for future use
}

func (receiver Type) Cast() (interface{}, error) {
	switch receiver.typ {
	case const_FRMSIZE_TYPE_DISCRETE:
		return *(*Discrete)(unsafe.Pointer(&receiver)), nil
	case const_FRMSIZE_TYPE_CONTINUOUS:
		return *(*Continuous)(unsafe.Pointer(&receiver)), nil
	case const_FRMSIZE_TYPE_STEPWISE:
		return *(*Stepwise)(unsafe.Pointer(&receiver)), nil
	default:
		return nil, fmt.Errorf("Unexpected frame size type: %d", receiver.typ)
	}
}

type Discrete struct {
	index              uint32                // Frame size number
	PixelFormat        v4l2_pixelformat.Type // Pixel format
	typ                uint32                // Frame size type the device supports.

	Width              uint32
	Height             uint32
	restOfFrameSize [4]uint32

	reserved        [2]uint32                // Reserved space for future use
}
