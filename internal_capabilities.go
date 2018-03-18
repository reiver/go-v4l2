package v4l2

import (
	"golang.org/x/sys/unix"

	"fmt"
	"unsafe"
)

// internalCapability is used to receive the capabilities of a V4L2 (Video4Linux video 2) device.
//
// Example:
//
//	device, err := v4l2.Open(v4l2.Video0)
//	
//	//...
//	
//	var cap internalCapability
//	
//	err := cap.Query(videoDevice)
//
// This is used internally of v4l2.Device
type internalCapability struct {
	driver       [16]uint8
	card         [32]uint8
	busInfo      [32]uint8
	version          uint32
	capabilities     uint32
	deviceCaps       uint32
	reserved      [3]uint32
}

func (receiver *internalCapability) QueryFd(fileDesciptor int) error {
	if nil == receiver {
		return errNilReceiver
	}

	_, _, errorNumber := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fileDesciptor),
		const_VIDIOC_QUERYCAP,
		uintptr(unsafe.Pointer(receiver)),
	)
	if 0 != errorNumber {
		return errorNumber
	}

	return nil
}

func (receiver internalCapability) BusInfo() string {

	max := len(receiver.busInfo)

	index := 0

	for ; index < max; index++ {
		if 0 == receiver.busInfo[index] {
			break
		}
	}

	return string(receiver.busInfo[:index])
}

func (receiver internalCapability) Card() string {

	max := len(receiver.card)

	index := 0

	for ; index < max; index++ {
		if 0 == receiver.card[index] {
			break
		}
	}

	return string(receiver.card[:index])
}

func (receiver internalCapability) Driver() string {

	max := len(receiver.driver)

	index := 0

	for ; index < max; index++ {
		if 0 == receiver.driver[index] {
			break
		}
	}

	return string(receiver.driver[:index])
}

func (receiver internalCapability) Version() string {
	version := receiver.version

	major := (version >> 16) & 0xFF
	minor := (version >>  8) & 0xFF
	patch := (version      ) & 0xFF

	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
