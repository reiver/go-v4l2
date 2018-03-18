package v4l2_pixelformat

import (
	"fmt"
)

type Type struct {
	value uint32
}

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

func (receiver Type) String() string {
	a := byte( receiver.value        & 0xff)
	b := byte((receiver.value >>  8) & 0xff)
	c := byte((receiver.value >> 16) & 0xff)
	d := byte((receiver.value >> 24) & 0xff)

	return fmt.Sprintf("%c%c%c%c", a, b, c, d)
}
