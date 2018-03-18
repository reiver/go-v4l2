package v4l2

// Card returns the name of the device.
//
// Example values that Card might return are:
//
// • "Laptop_Integrated_Webcam_HD"
//
// • "Yoyodyne TV/FM"
//
// The information that Card returns is intended for users.
// I.e., intended to be human readable)
// And can be used for things such as being use in a menu of available devices.
//
// Note that it is possible for a system to have multiple cards from the same brand.
//
// And thus it is possible for multiple devices to return the exact same value from Card.
//
// To make the name, shown to the user, unique, it can be combined with either the file name path,
// or what BusInfo returns.
func (receiver Device) Card() (string, error) {
	if !receiver.opened {
		return "", errNotOpen
	}
	if !receiver.capQueried {
		return "", errInternalError
	}

	return receiver.cap.Card(), nil
}

// MustCard is like BusCard, except it panic()s if there is an error.
func (receiver Device) MustCard() string {
	datum, err := receiver.Card()
	if nil != err {
		panic(err)
	}

	return datum
}
