package v4l2

const (
	FormatDescriptionFlagCompressed = 0x0001
	FormatDescriptionFlagEmulated   = 0x0002
)

// HasFlags returns true if the FormatDescription has all the flags inquired about,
// and returns false otherwise.
//
// Example:
//
//	var formatDescription v4l2.FormatDescription
//	
//	// ...
//	
//	if formatDescription.HasFlags(v4l2.FormatDescriptionFlagCompressed) {
//		//@TODO
//	}
//
// Example:
//
//	var formatDescription v4l2.FormatDescription
//	
//	// ...
//	
//	if formatDescription.HasFlags(v4l2.FormatDescriptionFlagEmulated) {
//		//@TODO
//	}
//
// Example:
//
//	var formatDescription v4l2.FormatDescription
//	
//	// ...
//	
//	if formatDescription.HasFlags(v4l2.FormatDescriptionFlagCompressed | v4l2.FormatDescriptionFlagEmulated) {
//		//@TODO
//	}
func (receiver FormatDescription) HasFlags(flags uint32) bool {
	return 0 != (receiver.internal.flags & flags)
}
