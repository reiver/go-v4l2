package v4l2

import (
	"github.com/reiver/go-v4l2/pixelformat"
)

// PixelFormat returns the formal family's pixel format.
//
// Some examples of what might be returned by PixelFormat are;
//
// • v4l2_pixelformat.FourCC("YUYV"),
//
// • v4l2_pixelformat.FourCC("MJPG"),
//
// • etc,
func (receiver FormatFamily) PixelFormat() v4l2_pixelformat.Type {
	return receiver.internal.pixelFormat
}
