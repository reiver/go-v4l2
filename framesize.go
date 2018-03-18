package v4l2

import (
	"fmt"
	"unsafe"
)

const (
	const_FRMSIZE_TYPE_DISCRETE   = 1
	const_FRMSIZE_TYPE_CONTINUOUS = 2
	const_FRMSIZE_TYPE_STEPWISE   = 3
)

type FrameSize struct {
	index        uint32 // Frame size number
	pixelFormat  uint32 // Pixel format
	typ          uint32 // Frame size type the device supports.

	frameSize [6]uint32 // Frame size

	reserved  [2]uint32 // Reserved space for future use
}

func (receiver FrameSize) PixelFormat() string {
	return fourcc(receiver.pixelFormat).String()
}

func (receiver FrameSize) Cast() (interface{}, error) {
	switch receiver.typ {
	case const_FRMSIZE_TYPE_DISCRETE:
		return *(*FrameSizeDiscrete)(unsafe.Pointer(&receiver)), nil
	case const_FRMSIZE_TYPE_CONTINUOUS:
		return *(*FrameSizeContinuous)(unsafe.Pointer(&receiver)), nil
	case const_FRMSIZE_TYPE_STEPWISE:
		return *(*FrameSizeStepwise)(unsafe.Pointer(&receiver)), nil
	default:
		return nil, fmt.Errorf("Unexpected frame size type: %d", receiver.typ)
	}
}

type FrameSizeDiscrete struct {
	index              uint32 // Frame size number
	pixelFormat        uint32 // Pixel format
	typ                uint32 // Frame size type the device supports.

	Width              uint32
	Height             uint32
	restOfFrameSize [4]uint32

	reserved        [2]uint32 // Reserved space for future use
}

func (receiver FrameSizeDiscrete) PixelFormat() string {
	return fourcc(receiver.pixelFormat).String()
}

type FrameSizeContinuous struct {
	index              uint32 // Frame size number
	pixelFormat        uint32 // Pixel format
	typ                uint32 // Frame size type the device supports.

	MinWidth           uint32 // Minimum frame width [pixel]
	MaxWidth           uint32 // Maximum frame width [pixel]
	StepWidth          uint32 // Frame width step size [pixel]
	MinHeight          uint32 // Minimum frame height [pixel]
	MaxHeight          uint32 // Maximum frame height [pixel]
	StepHeight         uint32 // Frame height step size [pixel]

	reserved        [2]uint32 // Reserved space for future use
}

func (receiver FrameSizeContinuous) PixelFormat() string {
	return fourcc(receiver.pixelFormat).String()
}

type FrameSizeStepwise struct {
	index              uint32 // Frame size number
	pixelFormat        uint32 // Pixel format
	typ                uint32 // Frame size type the device supports.

	MinWidth           uint32 // Minimum frame width [pixel]
	MaxWidth           uint32 // Maximum frame width [pixel]
	StepWidth          uint32 // Frame width step size [pixel]
	MinHeight          uint32 // Minimum frame height [pixel]
	MaxHeight          uint32 // Maximum frame height [pixel]
	StepHeight         uint32 // Frame height step size [pixel]

	reserved        [2]uint32 // Reserved space for future use
}

func (receiver FrameSizeStepwise) PixelFormat() string {
	return fourcc(receiver.pixelFormat).String()
}
