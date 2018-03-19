package v4l2_format

import (
	"github.com/reiver/go-v4l2/buftype"
)

// This is more or less equivalent to the Linux V4L2 (Video4Linux version 2)
// v4l2_format.type with the v4l2_format.raw_data.
type Type struct {
	buftype      v4l2_buftype.Type
	rawData [200]byte
}

// BufferType returns the buffer type of the format.
func (receiver Type) BufferType() v4l2_buftype.Type {
	return receiver.buftype
}
