package v4l2

import (
	"github.com/reiver/go-v4l2/pixelformat"
)

// FormatFamily is a format family.
//
// A format family contains:
//
// • a human-readable description,
//
// • a pixel format (as a FOURCC code), and
//
// • flags.
//
// As well as logicall "containing" a list of frame sizes.
//
// The possible flags are given by the constants:
//
// • v4l2.FormatFamilyFlagCompressed, and
//
// • v4l2.FormatFamilyFlagEmulated.
//
// Samples:
//
// An example format migh have have values such as:
//
// • description:  "YUYV 4:2:2"
//
// • pixel format: "YUYV"
//
// • (flag) compressed: false
//
// • (flag) emulated:   false
//
// Or an example format migh have have values such as:
//
// • description:  "Motion-JPEG"
//
// • pixel format: "MJPG"
//
// • (flag) compressed: true
//
// • (flag) emulated:   false
//
// Etc.
//
// Usually, one would get a series of formats by iterating through all the supported formats
// that are supported by a device.
//
// Example:
//
//	var device v4l2.Device
//	
//	// ...
//	
//	formats, err := device.Formats()
//	if nil != err {
//		return err
//	}
//	defer formats.Close()
//	
//	var formatFamily v4l2.FormatFamily // <---- NOTE THAT THIS IS THE v4l2.FormatFamily TYPE.
//	forformats.Next() {
//
//		err := formats.Decode(&formatFamily) // <---- NOTE THAT WE ARE PUTTING A NEW VALUE INTO THE v4l2.FormatFamily HERE.
//		if nil != err {
//			fmt.Fprintf(os.Stderr, "ERROR: Problem decoding format family: (%T) %v \n", err, err)
//			return err
//		}
//		
//		fmt.Printf("[format family] %q (%q) {compressed=%t} {emulated=%t} \n",
//			formatFamily.Description(),
//			formatFamily.PixelFormat(),
//			formatFamily.HasFlags(v4l2.FormatFamilyFlagCompressed),
//			formatFamily.HasFlags(v4l2.FormatFamilyFlagEmulated),
//		)
//	}
//	if err := formats.Err(); nil != err {
//		return err
//	}
//
// v4l2.FormatFamily is the same as the V4L2 (Video4Linux version 2) type v4l2_fmtdesc.
type FormatFamily struct {
	device  *Device
	internal internalFormatFamily
}

type internalFormatFamily struct {
	index           uint32                // Format number
	typ             uint32                // enum v4l2_buf_type
	flags           uint32
	description [32]byte                  // Description string
	pixelFormat     v4l2_pixelformat.Type // Format fourcc
	reserved     [4]uint32
}
