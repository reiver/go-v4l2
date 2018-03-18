package v4l2

import (
	"golang.org/x/sys/unix"
)

// Device represents a V4L2 (Video4Linux version 2) device.
type Device struct {
	opened         bool
	capQueried     bool
	fileDescriptor int
	cap internalCapability
}

// Open opens a V4L2 device.
//
// Example:
//
//	device, err := v4l2.Open(v4l2.Video0)
//	if nil != err {
//		return err
//	})
//	defer device.Close()
func Open(name string) (*Device, error) {
	var device Device

	err := device.Open(name)
	if nil != err {
		return nil, err
	}

	return &device, nil
}

// MustOpen is like Open, except it panic()s if there is an error.
func MustOpen(name string) *Device {
	datum, err := Open(name)
	if nil != err {
		panic(err)
	}

	return datum
}

// Open opens a V4L2 device.
//
// (This method exist so as not to create GC pressure.)
//
// Example:
//
//	var device Device
//	
//	err := device.Open(v4l2.Video0)
//	if nil != err {
//		return err
//	})
//	defer device.Close()
func (receiver *Device) Open(name string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if receiver.opened {
		return errOpen
	}

	{
		var err error

		receiver.fileDescriptor, err = unix.Open(name, unix.O_RDWR|unix.O_NONBLOCK, 0666)
		if  nil != err {
			return err
		}
		receiver.opened = true
	}

	if err := receiver.cap.QueryFd(receiver.fileDescriptor) ; nil != err {
		return err
	}
	receiver.capQueried = true

	return nil
}

// Close closes the device.
//
// You should always close a device after you finished with it, so that you do not create a resource leak.
func (receiver *Device) Close() error {
	if nil == receiver {
		return nil
	}

	if !receiver.opened {
		return nil
	}

	err := unix.Close(receiver.fileDescriptor)
	if nil != err {
		return err
	}

	receiver.opened = false

	return nil
}
