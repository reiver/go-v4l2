package v4l2

// FormatDescription is a format description.
//
// A format description contains:
//
// • a human-readable description,
//
// • a pixel format (as a FOURCC code), and
//
// • flags.
//
// The possible flags are given by the constants:
//
// • v4l2.FormatDescriptionFlagCompressed, and
//
// • v4l2.FormatDescriptionFlagEmulated.
//
// An example format description migh have have values such as:
//
// • description:  "YUYV 4:2:2"
//
// • pixel format: "YUYV"
//
// • (flag) compressed: false
//
// • (flag) emulated:   false
//
// Or an example format description migh have have values such as:
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
// Usually, one would get a series of format descriptions by  iterating through all the support formats
// that are supported by a device.
//
// Example:
//
//	var device v4l2.Device
//	
//	// ...
//	
//	formatDescriptions, err := device.FormatDescriptions()
//	if nil != err {
//		return err
//	}
//	defer formatDescriptions.Close()
//	
//	var formatDescription v4l2.FormatDescription // <---- NOTE THAT THIS IS THE v4l2.FormatDescription TYPE.
//	forformatDescriptions.Next() {
//
//		err := formatDescriptions.Decode(&formatDescription) // <---- NOTE THAT WE ARE PUTTING A NEW VALUE INTO THE v4l2.FormatDescription HERE.
//		if nil != err {
//			fmt.Fprintf(os.Stderr, "ERROR: Problem decoding format description: (%T) %v \n", err, err)
//			return err
//		}
//		
//		fmt.Printf("[format description] %q (%q) {compressed=%t} {emulated=%t} \n",
//			formatDescription.Description(),
//			formatDescription.PixelFormat(),
//			formatDescription.HasFlags(v4l2.FormatDescriptionFlagCompressed),
//			formatDescription.HasFlags(v4l2.FormatDescriptionFlagEmulated),
//		)
//	}
//	if err := formatDescriptions.Err(); nil != err {
//		return err
//	}
//
// v4l2.FormatDescription is the same as the V4L2 (Video4Linux version 2) type v4l2_fmtdesc.
type FormatDescription struct {
	index           uint32 // Format number
	typ             uint32 // enum v4l2_buf_type
	flags           uint32
	description [32]byte   // Description string
	pixelformat     uint32 // Format fourcc
	reserved     [4]uint32
}
