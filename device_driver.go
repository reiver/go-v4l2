package v4l2

// Driver returns the name of the driver.
//
// Example values that Driver might return are:
//
// • "uvcvideo"
//
// • "bttv"
//
// If an application can only work with specific drivers, then the information returned
// from Driver can be used to detect which driver is behind this device.
//
// Also, if certain drivers have bugs, then this can (in part) be used to create driver
// specific work-arounds.
func (receiver Device) Driver() (string, error) {
	if !receiver.opened {
		return "", errNotOpen
	}
	if !receiver.capQueried {
		return "", errInternalError
	}

	return receiver.cap.Driver(), nil
}

// MustDriver is like Driver, except it panic()s if there is an error.
//
// Example:
//
//	device := v4l2.MustOpen(v4l2.Video0)
//	defer device.MustClose()
//	
//	fmt.Printf("Driver: %q \n", device.MustDriver())
func (receiver Device) MustDriver() string {
	datum, err := receiver.Driver()
	if nil != err {
		panic(err)
	}

	return datum
}
