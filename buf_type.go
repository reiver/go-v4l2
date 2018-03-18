package v4l2

// These were mimicked from:
// api/linux/videodev2.h
const (
	const_V4L2_BUF_TYPE_VIDEO_CAPTURE        = 1
	const_V4L2_BUF_TYPE_VIDEO_OUTPUT         = 2
	const_V4L2_BUF_TYPE_VIDEO_OVERLAY        = 3
	const_V4L2_BUF_TYPE_VBI_CAPTURE          = 4
	const_V4L2_BUF_TYPE_VBI_OUTPUT           = 5
	const_V4L2_BUF_TYPE_SLICED_VBI_CAPTURE   = 6
	const_V4L2_BUF_TYPE_SLICED_VBI_OUTPUT    = 7
	const_V4L2_BUF_TYPE_VIDEO_OUTPUT_OVERLAY = 8 // Experimental
	const_V4L2_BUF_TYPE_VIDEO_CAPTURE_MPLANE = 9
	const_V4L2_BUF_TYPE_VIDEO_OUTPUT_MPLANE  = 10
	const_V4L2_BUF_TYPE_SDR_CAPTURE          = 11

	const_V4L2_BUF_TYPE_PRIVATE              = 0x80 // Deprecated, do not use
)

