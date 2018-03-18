package v4l2_buftype

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
