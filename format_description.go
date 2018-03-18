package v4l2

// Description returns a human-readable description for the format description.
//
// Some examples of what might be returned by Description are;
//
// • "YUYV 4:2:2",
//
// • "Motion-JPEG",
//
// • etc.
func (receiver FormatDescription) Description() string {

	max := len(receiver.internal.description)

	index := 0

	for ; index < max; index++ {
		if 0 == receiver.internal.description[index] {
			break
		}
	}

	return string(receiver.internal.description[:index])
}

