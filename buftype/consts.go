package v4l2_buftype

// These were mimicked from:
// api/linux/videodev2.h
const (
	VideoCapture            =    1 // (V4L2_BUF_TYPE_VIDEO_CAPTURE)
	VideoOutput             =    2 // (V4L2_BUF_TYPE_VIDEO_OUTPUT)
	VideoOverlay            =    3 // (V4L2_BUF_TYPE_VIDEO_OVERLAY)
	VBICapture              =    4 // (V4L2_BUF_TYPE_VBI_CAPTURE)
	VBIOutput               =    5 // (V4L2_BUF_TYPE_VBI_OUTPUT)
	SlicedVBICapture        =    6 // (V4L2_BUF_TYPE_SLICED_VBI_CAPTURE)
	SlicedVBIOutput         =    7 // (V4L2_BUF_TYPE_SLICED_VBI_OUTPUT)
	VideoOutputOverlay      =    8 // (V4L2_BUF_TYPE_VIDEO_OUTPUT_OVERLAY)  Experimental
	VideoCaptureMultiplanar =    9 // (V4L2_BUF_TYPE_VIDEO_CAPTURE_MPLANE)
	VideoOutputMultiplanar  =   10 // (V4L2_BUF_TYPE_VIDEO_OUTPUT_MPLANE)
	SDRCapture              =   11 // (V4L2_BUF_TYPE_SDR_CAPTURE)

	Private                 = 0x80 // (V4L2_BUF_TYPE_PRIVATE)  Deprecated, do not use
)
