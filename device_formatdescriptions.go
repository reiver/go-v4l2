package v4l2

import (
	"golang.org/x/sys/unix"

	"unsafe"
)

// FormatDescriptions returns an iterator that enables you to list out all the supported formats by the device.
//
// Example:
//
//      var device v4l2.Device
//      
//      // ...
//      
//      formatDescriptions, err := device.FormatDescriptions() // <---- NOTE THAT THIS IS WHERE THE v4l2.Device.FormatDescriptions() METHOD IS CALLED.
//      if nil != err {
//              return err
//      }
//      defer formatDescriptions.Close()
//      
//      var formatDescription v4l2.Format
//      forformatDescriptions.Next() {
//
//              err := formatDescriptions.Decode(&formatDescription)
//              if nil != err {
//                      fmt.Fprintf(os.Stderr, "ERROR: Problem decoding format description: (%T) %v \n", err, err)
//                      return err
//              }
//
//              fmt.Printf("[format description] %q (%q) {compressed=%t} {emulated=%t} \n",
//                      formatDescription.Description(),
//                      formatDescription.PixelFormat(),
//                      formatDescription.HasFlags(v4l2.FormatFlagCompressed),
//                      formatDescription.HasFlags(v4l2.FormatFlagEmulated),
//              )
//      }
//      if err := formatDescriptions.Err(); nil != err {
//              return err
//      }
func (receiver *Device) FormatDescriptions() (FormatDescriptions, error) {
	if err := receiver.unfit(); nil != err {
		return FormatDescriptions{}, err
	}

	return FormatDescriptions{
		device: receiver,
	}, nil
}

// FormatDescriptions is an interator that enables you to list out all the supported formats by the device.
type FormatDescriptions struct {
	device *Device
	err     error
	datum   internalFormat
}

// Close closes the FormatDescriptions iterator.
func (receiver *FormatDescriptions) Close() error {
	if nil == receiver {
		return nil
	}

	receiver.device      = nil
	receiver.err         = nil
	receiver.datum.index = 0

	return nil
}

// Decode loads the next format description (previously obtained by calling Next).
func (receiver FormatDescriptions) Decode(x interface{}) error {
	if nil != receiver.err {
		return receiver.err
	}

	p, ok := x.(*Format)
	if !ok {
		return errUnsupportedType
	}

	p.device   = receiver.device
	p.internal = receiver.datum

	return nil
}

// Err returns any errors that occurred when Next was called.
func  (receiver *FormatDescriptions) Err() error {
	if nil == receiver {
		return errNilReceiver
	}

	return receiver.err
}

// Next fetches the next format description.
//
// If there is a next format description, it returns true.
// And the next format description get be obtained by calling Decode.
//
// If there is not next format description, then it returns false.
func (receiver *FormatDescriptions) Next() bool {
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
