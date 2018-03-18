package v4l2

// safe returns an error if
func (receiver *Device) unfit() error {
	if nil == receiver {
		return errNilReceiver
	}

	if !receiver.opened {
		return errNotOpen
	}
	if !receiver.capQueried {
		return errInternalError
	}

	return nil
}
