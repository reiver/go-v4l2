/*
Package v4l2_pixelformat provides a type used to represents a V4L2 (Video4Linux version 2) pixel format.

Example

	var pixelFormat v4l2_pixelformat.Type
	
	pixelFormat = v4l2_pixelformat.FourCC("MJPG")

Examle

	var pixelFormat
	
	// ...
	
	var humanReadable string = pixelFormat.String()
*/
package v4l2_pixelformat
