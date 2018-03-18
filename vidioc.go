package v4l2

const (
	// A Golang conversion of the following C code:
	//
	// #define      VIDIOC_QUERYCAP   _IOR ('V', 0, struct v4l2_capability)
	const_VIDIOC_QUERYCAP = (const_IOC_READ                      << const_IOC_DIRSHIFT)  |
	                        (uintptr('V')                        << const_IOC_TYPESHIFT) |
	                        (0                                   << const_IOC_NRSHIFT)   |
	                        (unsafe.Sizeof(internalCapability{}) << const_IOC_SIZESHIFT)
)

