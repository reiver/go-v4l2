package v4l2

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
	if !receiver.opened {
		return false, errNotOpen
	}
	if !receiver.capQueried {
		return false, errInternalError
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
