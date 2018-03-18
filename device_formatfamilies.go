package v4l2

import (
	"golang.org/x/sys/unix"

	"unsafe"
)

// FormatFamilies returns an iterator that enables you to list out all the supported formats by the device.
//
// Example:
//
//      var device v4l2.Device
//      
//      // ...
//      
//      formats, err := device.FormatFamilies() // <---- NOTE THAT THIS IS WHERE THE v4l2.Device.FormatFamilies() METHOD IS CALLED.
//      if nil != err {
//              return err
//      }
//      defer formats.Close()
//      
//      var formatFamily v4l2.FormatFamily
//      forformats.Next() {
//
//              err := formats.Decode(&formatFamily)
//              if nil != err {
//                      fmt.Fprintf(os.Stderr, "ERROR: Problem decoding format: (%T) %v \n", err, err)
//                      return err
//              }
//
//              fmt.Printf("[format] %q (%q) {compressed=%t} {emulated=%t} \n",
//                      formatDescription.Description(),
//                      formatDescription.PixelFormat(),
//                      formatDescription.HasFlags(v4l2.FormatFamilyFlagCompressed),
//                      formatDescription.HasFlags(v4l2.FormatFamilyFlagEmulated),
//              )
//      }
//      if err := formats.Err(); nil != err {
//              return err
//      }
func (receiver *Device) FormatFamilies() (FormatFamilies, error) {
	if err := receiver.unfit(); nil != err {
		return FormatFamilies{}, err
	}

	return FormatFamilies{
		device: receiver,
	}, nil
}

// FormatFamilies is an interator that enables you to list out all the supported formats by the device.
type FormatFamilies struct {
	device *Device
	err     error
	datum   internalFormatFamily
}

// Close closes the FormatFamilies iterator.
func (receiver *FormatFamilies) Close() error {
	if nil == receiver {
		return nil
	}

	receiver.device      = nil
	receiver.err         = nil
	receiver.datum.index = 0

	return nil
}

// Decode loads the next format (previously obtained by calling Next).
func (receiver FormatFamilies) Decode(x interface{}) error {
	if nil != receiver.err {
		return receiver.err
	}

	p, ok := x.(*FormatFamily)
	if !ok {
		return errUnsupportedType
	}

	p.device   = receiver.device
	p.internal = receiver.datum

	return nil
}

// Err returns any errors that occurred when Next was called.
func  (receiver *FormatFamilies) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.err
}

// Next fetches the next format.
//
// If there is a next format, it returns true.
// And the next format get be obtained by calling Decode.
//
// If there is not next format, then it returns false.
func (receiver *FormatFamilies) Next() bool {
	if nil == receiver {
		return false
	}

	device := receiver.device
	if nil == device {
		receiver.err = errInternalError
		return false
	}

	receiver.datum.typ = const_V4L2_BUF_TYPE_VIDEO_CAPTURE

	_, _, errorNumber := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(device.fileDescriptor),
		const_VIDIOC_ENUM_FMT,
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
