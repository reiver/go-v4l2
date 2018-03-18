package v4l2

import (
	"golang.org/x/sys/unix"

	"unsafe"
)

// FrameSizes returns an iterator that enables you to list out all the supported frame sizes by the format family.
func (receiver *FormatFamily) FrameSizes() (FrameSizes, error) {
	if nil == receiver {
		return FrameSizes{}, errNilReceiver
	}

	device := receiver.device
	if nil == device {
		return FrameSizes{}, errInternalError
	}

	if err := device.unfit(); nil != err {
		return FrameSizes{}, err
	}

	return FrameSizes{
		device:      receiver.device,
		pixelFormat: receiver.internal.pixelFormat,
	}, nil
}

// FrameSizes is an interator that enables you to list out all the supported formats by the format family.
type FrameSizes struct {
	device      *Device
	pixelFormat  uint32
	err          error
	datum        FrameSize
}

// Close closes the FrameSizes iterator.
func (receiver *FrameSizes) Close() error {
	if nil == receiver {
		return nil
	}

	receiver.device      = nil
	receiver.err         = nil
	receiver.datum.index = 0

	return nil
}

// Decode loads the next frame size (previously obtained by calling Next).
func (receiver FrameSizes) Decode(x interface{}) error {
	if nil != receiver.err {
		return receiver.err
	}

	p, ok := x.(*FrameSize)
	if !ok {
		return errUnsupportedType
	}

	*p = receiver.datum

	return nil
}

// Err returns any errors that occurred when Next was called.
func  (receiver *FrameSizes) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.err
}

// Next fetches the next frame size.
//
// If there is a next frame size, it returns true.
// And the next frame size get be obtained by calling Decode.
//
// If there is not next frame size, then it returns false.
func (receiver *FrameSizes) Next() bool {
	if nil == receiver {
		return false
	}

	device := receiver.device
	if nil == device {
		receiver.err = errInternalError
		return false
	}

	receiver.datum.pixelFormat = receiver.pixelFormat

	_, _, errorNumber := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(device.fileDescriptor),
		const_VIDIOC_ENUM_FRAMESIZES,
		uintptr(unsafe.Pointer(&receiver.datum)),
	)
	if unix.EINVAL == errorNumber {
		return false
	}
	if 0 != errorNumber {
		receiver.err = errorNumber
		return false
	}

	receiver.datum.index++

	return true
}
