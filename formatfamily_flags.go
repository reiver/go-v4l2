package v4l2

const (
	FormatFamilyFlagCompressed = 0x0001 // V4L2_FMT_FLAG_COMPRESSED
	FormatFamilyFlagEmulated   = 0x0002 // V4L2_FMT_FLAG_EMULATED
)

// HasFlags returns true if the format family has all the flags inquired about,
// and returns false otherwise.
//
// Example:
//
//	var formatFamily v4l2.FormatFamily
//	
//	// ...
//	
//	if formatFamily.HasFlags(v4l2.FormatFamilyFlagCompressed) {
//		//@TODO
//	}
//
// Example:
//
//	var formatFamily v4l2.FormatFamily
//	
//	// ...
//	
//	if formatFamily.HasFlags(v4l2.FormatFamilyFlagEmulated) {
//		//@TODO
//	}
//
// Example:
//
//	var formatFamily v4l2.FormatFamily
//	
//	// ...
//	
//	if formatFamily.HasFlags(v4l2.FormatFamilyFlagCompressed | v4l2.FormatFamilyFlagEmulated) {
//		//@TODO
//	}
func (receiver FormatFamily) HasFlags(flags uint32) bool {
	return 0 != (receiver.internal.flags & flags)
}
