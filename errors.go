package v4l2

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errNilReceiver   = errors.New("Nil Receiver")
	errNotOpen       = errors.New("Not Open")
	errOpen          = errors.New("Open")
)
