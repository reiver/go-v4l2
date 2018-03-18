package v4l2

// PixelFormat returns the pixel format as a FOURCC (4CC) code.
//
// Some examples of what might be returned by PixelFormat are;
//
// • "YUYV",
//
// • "MJPG",
//
// • etc,
func (receiver FormatFamily) PixelFormat() string {
	return fourcc(receiver.internal.pixelFormat).String()
}
