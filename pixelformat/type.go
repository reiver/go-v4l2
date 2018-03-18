package v4l2_pixelformat

import (
	"fmt"
)

// Type represents a V4L2 (Video4Linux version 2) pixel format.
type Type struct {
	value uint32
}

// FourCC returns a pixelformat.Type from a FOURCC code.
//
// Example:
//
//	var pixelFormat v4l2_pixelformat.Type
//	
//	pixelFormat = v4l2_pixelformat.FourCC("YUYV")
func FourCC(value string) Type {

	var a byte = ' '
	var b byte = ' '
	var c byte = ' '
	var d byte = ' '

	{
		length := len(value)

		if 1 <= length  {
			a = byte(value[0])
		}
		if 2 <= length  {
			b = byte(value[1])
		}
		if 3 <= length  {
			c = byte(value[2])
		}
		if 4 <= length  {
			d = byte(value[3])
		}
	}

	var code uint32

	code =  uint32(a)        |
	       (uint32(b) <<  8) |
	       (uint32(c) << 16) |
	       (uint32(d) << 24)

	return Type{code}
}

// String returns a human-readable version of v4l2_pixelformat.Type.
//
// So, rather than returning a machine-readable binary uint32,
// it returns a human-readable FOURCC code string.
//
// So, for example, even if the value stored as the machine-readable uint32 0x47504a4d,
// String will return this as "MJPG".
//
// Or also, for example, even if the value stored as the machine-readable uint32 0x30385056,
// String will return this as "VP80".
func (receiver Type) String() string {
	a := byte( receiver.value        & 0xff)
	b := byte((receiver.value >>  8) & 0xff)
	c := byte((receiver.value >> 16) & 0xff)
	d := byte((receiver.value >> 24) & 0xff)

	return fmt.Sprintf("%c%c%c%c", a, b, c, d)
}
