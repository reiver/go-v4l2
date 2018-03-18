package v4l2

const (
	FormatFlagCompressed = 0x0001
	FormatFlagEmulated   = 0x0002
)

// HasFlags returns true if the Format has all the flags inquired about,
// and returns false otherwise.
//
// Example:
//
//	var format v4l2.Format
//	
//	// ...
//	
//	if format.HasFlags(v4l2.FormatFlagCompressed) {
//		//@TODO
//	}
//
// Example:
//
//	var format v4l2.Format
//	
//	// ...
//	
//	if format.HasFlags(v4l2.FormatFlagEmulated) {
//		//@TODO
//	}
//
// Example:
//
//	var format v4l2.Format
//	
//	// ...
//	
//	if format.HasFlags(v4l2.FormatFlagCompressed | v4l2.FormatFlagEmulated) {
//		//@TODO
//	}
func (receiver Format) HasFlags(flags uint32) bool {
	return 0 != (receiver.internal.flags & flags)
}
