package v4l2

import (
	"fmt"
)

// PixelFormat returns the pixel format as a FOURCC (4CC) code.
//
// Some examples of what might be returned by PixelFormat are;
//
// • "YUYV",
//
// • "MJPG",
//
// • etc,
func (receiver FormatDescription) PixelFormat() string {
	a := byte( receiver.internal.pixelformat        & 0xff)
	b := byte((receiver.internal.pixelformat >>  8) & 0xff)
	c := byte((receiver.internal.pixelformat >> 16) & 0xff)
	d := byte((receiver.internal.pixelformat >> 24) & 0xff)

	return fmt.Sprintf("%c%c%c%c", a, b, c, d)
}
