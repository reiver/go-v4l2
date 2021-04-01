package v4l2

import "errors"

// HasCapability return whether the device has a capability a particular capability.
//
// Example:
//
//	device, err := v4l2.Open(v4l2.Video17)
//
//	// ...
//
//	hasVideoCaptureCapability, err := device.HasCapability(CapabilityVideoCapture)
func (receiver Device) HasCapability(cap uint32) (bool, error) {
	if err := receiver.unfit(); nil != err {
		return false, err
	}

	has := 0 != (receiver.cap.capabilities & cap)

	return has, nil
}

// MustHasCapability is like HasCapability, except it panic()s if there is an error.
//
//      device := v4l2.MustOpen(v4l2.Video0)
//      defer device.MustClose()
//      
//      fmt.Printf("Has Streaming Capability: %t \n", device.MustHasCapability(v4l2.CapabilityStreaming))
func (receiver Device) MustHasCapability(cap uint32) bool {
	datum, err := receiver.HasCapability(cap)
	if nil != err {
		panic(err)
	}

	return datum
}


// HasDeviceCapability return whether the device has a particular device capability.
//
// Example:
//
//	device, err := v4l2.Open(v4l2.Video17)
//
//	// ...
//
//	hasVideoCaptureCapability, err := device.HasDeviceCapability(CapabilityVideoCapture)
func (receiver Device) HasDeviceCapability(cap uint32) (bool, error) {
	if hasDeviceCaps, err := receiver.HasCapability(CapabilityDeviceCaps); err != nil {
		return false, err
	} else if !hasDeviceCaps {
		return false, errors.New("No device caps available")
	}
	
	has := 0 != (receiver.cap.deviceCaps & cap)
	return has, nil
}

// MustHasDeviceCapability return whether the device has a particular device capability.
//
// Example:
//
//	device, err := v4l2.Open(v4l2.Video17)
//
//	// ...
//
//	hasVideoCaptureCapability, err := device.MustHasDeviceCapability(CapabilityVideoCapture)
func (receiver Device) MustHasDeviceCapability(cap uint32) bool {
	if hasDeviceCaps := receiver.MustHasCapability(CapabilityDeviceCaps); !hasDeviceCaps {
		panic(errors.New("No device caps available"))
	}
	
	has := 0 != (receiver.cap.deviceCaps & cap)
	return has
}
