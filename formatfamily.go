package v4l2

// Format is a format.
//
// A format contains:
//
// • a human-readable description,
//
// • a pixel format (as a FOURCC code), and
//
// • flags.
//
// The possible flags are given by the constants:
//
// • v4l2.FormatFlagCompressed, and
//
// • v4l2.FormatFlagEmulated.
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
//	var format v4l2.Format // <---- NOTE THAT THIS IS THE v4l2.Format TYPE.
//	forformats.Next() {
//
//		err := formats.Decode(&format) // <---- NOTE THAT WE ARE PUTTING A NEW VALUE INTO THE v4l2.Format HERE.
//		if nil != err {
//			fmt.Fprintf(os.Stderr, "ERROR: Problem decoding format: (%T) %v \n", err, err)
//			return err
//		}
//		
//		fmt.Printf("[format] %q (%q) {compressed=%t} {emulated=%t} \n",
//			format.Description(),
//			format.PixelFormat(),
//			format.HasFlags(v4l2.FormatFlagCompressed),
//			format.HasFlags(v4l2.FormatFlagEmulated),
//		)
//	}
//	if err := formats.Err(); nil != err {
//		return err
//	}
//
// v4l2.Format is the same as the V4L2 (Video4Linux version 2) type v4l2_fmtdesc.
type Format struct {
	device  *Device
	internal internalFormat
}

type internalFormat struct {
	index           uint32 // Format number
	typ             uint32 // enum v4l2_buf_type
	flags           uint32
	description [32]byte   // Description string
	pixelFormat     uint32 // Format fourcc
	reserved     [4]uint32
}
