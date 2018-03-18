package v4l2

// Version returns the version number of the driver.
//
// Example values that Version might return are:
//
// • "4.4.98"
//
// • "4.14.0"
func (receiver Device) Version() (string, error) {
	if !receiver.opened {
		return "", errNotOpen
	}
	if !receiver.capQueried {
		return "", errInternalError
	}

	return receiver.cap.Version(), nil
}

// MustVersion is like Version, except it panic()s if there is an error.
//
// Example:
//
//      device := v4l2.MustOpen(v4l2.Video0)
//      defer device.MustClose()
//      
//      fmt.Printf("Version: %q \n", device.MustVersion())
func (receiver Device) MustVersion() string {
	datum, err := receiver.Version()
	if nil != err {
		panic(err)
	}

	return datum
}
