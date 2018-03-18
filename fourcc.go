package v4l2

import (
	"fmt"
)

type fourcc uint32

func (receiver fourcc) String() string {
	a := byte( receiver        & 0xff)
	b := byte((receiver >>  8) & 0xff)
	c := byte((receiver >> 16) & 0xff)
	d := byte((receiver >> 24) & 0xff)

	return fmt.Sprintf("%c%c%c%c", a, b, c, d)
}
