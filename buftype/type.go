package v4l2_buftype

import (
	"fmt"
)

// Type represents a V4L2 (Video4Linux version 2) buffer type.
type Type struct {

	// NOTE THAT THIS STRUCT MUST BE EXACLTLY 4 BYTES IN SIZE.
	//
	// THIS IS WHAT THE V4L2 (Video4Linux version 2) DRIVERS EXPECT.
	//
	// THUS WE CANNOT PUT EXTRA (USEFUL) HIDDEN FIELDS IN HERE.
	//
	// THERE MUST ONLY BE A SINGLE uint32 FIELD IN THIS STRUCT!

	value uint32
}

// Datum returns a v4l2_buftype.Type from a uint32.
//
// Example
//
//	var buftype v4l2_buftype.Type
//	
//	buftype = v4l2_buftype.Datum(v4l2_buftype.VideoCapture)
func Datum(value uint32) Type {
	return Type{value}
}

// String returns a human-readable version of the V4L2 (Video4Linux version 2) buffer type.
//
// So, rather than returning the uint32 1, it returns the string "V4L2_BUF_TYPE_VIDEO_CAPTURE".
func (receiver Type) String() string {

	ui32 := receiver.value

	switch ui32 {
	case VideoCapture:
		return "V4L2_BUF_TYPE_VIDEO_CAPTURE"
	case VideoOutput:
		return "V4L2_BUF_TYPE_VIDEO_OUTPUT"
	case VideoOverlay:
		return "V4L2_BUF_TYPE_VIDEO_OVERLAY"
	case VBICapture:
		return "V4L2_BUF_TYPE_VBI_CAPTURE"
	case VBIOutput:
		return "V4L2_BUF_TYPE_VBI_OUTPUT"
	case SlicedVBICapture:
		return "V4L2_BUF_TYPE_SLICED_VBI_CAPTURE"
	case SlicedVBIOutput:
		return "V4L2_BUF_TYPE_SLICED_VBI_OUTPUT"
	case VideoOutputOverlay:
		return "V4L2_BUF_TYPE_VIDEO_OUTPUT_OVERLAY"
	case VideoCaptureMultiplanar:
		return "V4L2_BUF_TYPE_VIDEO_CAPTURE_MPLANE"
	case VideoOutputMultiplanar:
		return "V4L2_BUF_TYPE_VIDEO_OUTPUT_MPLANE"
	case SDRCapture:
		return "V4L2_BUF_TYPE_SDR_CAPTURE"

	case Private:
		return "V4L2_BUF_TYPE_PRIVATE"

	default:
		return fmt.Sprintf("[UNKNOWN: %d]", ui32)
	}
}
