package v4l2_format

import (
	"github.com/reiver/go-v4l2/buftype"
	"github.com/reiver/go-v4l2/pixelformat"

	"unsafe"
)

// This is more or less equivalent to the Linux V4L2 (Video4Linux version 2)
// v4l2_format.type with the v4l2_format.pix (v4l2_pix_format).
//
// And thus, where v4l2_format.type is V4L2_BUF_TYPE_VIDEO_CAPTURE.
type Pix struct {
	buftype      v4l2_buftype.Type

	Width        uint32
	Height       uint32
	PixelFormat  v4l2_pixelformat.Type
	field        uint32 // enum v4l2_field
	bytesperline uint32 // for padding, zero if unused
	sizeimage    uint32
	colorspace   uint32 // enum v4l2_colorspace
	priv         uint32 // private data, depends on pixelformat
	flags        uint32 // format flags (V4L2_PIX_FMT_FLAG_*)
	ycbcrEnc     uint32 // enum v4l2_ycbcr_encoding
	quantization uint32 // enum v4l2_quantization

	padding [200-(11*4)]byte // This struct MUST be 200 bytes long, so we add padding to make it that.
}

// CastFormat returns the v4l2_format.Pix as a v4l2_format.Type.
//
// Basically, doing something similar to a type cast.
// Although there is copying involved.
func (receiver Pix) CastFormat() Type {

	receiver.buftype = v4l2_buftype.Datum(v4l2_buftype.VideoCapture)

	return *(*Type)(unsafe.Pointer(&receiver))
}
