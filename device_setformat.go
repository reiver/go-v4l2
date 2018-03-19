package v4l2

import (
	"github.com/reiver/go-v4l2/format"

	"golang.org/x/sys/unix"

	"unsafe"
)

func (receiver *Device) SetFormat(formatcaster v4l2_format.Caster) error {
	if err := receiver.unfit(); nil != err {
		return err
	}

	format := formatcaster.CastFormat()

	_, _, errorNumber := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(receiver.fileDescriptor),
		const_VIDIOC_S_FMT,
		uintptr(unsafe.Pointer(&format)),
	)
	if 0 != errorNumber {
		return errorNumber
	}

	return nil
}
