package v4l2_framesize

import (
	"fmt"
	"unsafe"
)

const (
	const_TYPE_DISCRETE   = 1 // (V4L2_FRMSIZE_TYPE_DISCRETE)
	const_TYPE_CONTINUOUS = 2 // (V4L2_FRMSIZE_TYPE_CONTINUOUS)
	const_TYPE_STEPWISE   = 3 // (V4L2_FRMSIZE_TYPE_STEPWISE)
)

// Cast tries to turn a v4l2_framesize.Type into a:
// v4l2_framesize.Discrete,
// v4l2_framesize.Continuous, or
// v4l2_framesize.Stepwise.
//
// Example:
//
//	switch casted := frameSize.Cast().(type) {
//	case v4l2_framesize.Continuous
//		//@TODO
//	case v4l2_framesize.Discrete:
//		//@TODO
//	case v4l2_framesize.Stepwise
//		//@TODO
//	default:
//		//@TODO: Error, unknown framesize type.
//	}
func (receiver Type) Cast() (interface{}, error) {
	switch receiver.typ {
	case const_TYPE_DISCRETE:
		return *(*Discrete)(unsafe.Pointer(&receiver)), nil
	case const_TYPE_CONTINUOUS:
		return *(*Continuous)(unsafe.Pointer(&receiver)), nil
	case const_TYPE_STEPWISE:
		return *(*Stepwise)(unsafe.Pointer(&receiver)), nil
	default:
		return nil, fmt.Errorf("Unexpected frame size type: %d", receiver.typ)
	}
}
