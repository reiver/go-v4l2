package v4l2

import (
	"unsafe"
)

// ioctl codes for V4L2 (Video4Linux version 2) video devices.
const (
	// A Golang conversion of the following C code:
	//
	// #define  VIDIOC_QUERYCAP  _IOR('V',  0, struct v4l2_capability)
	const_VIDIOC_QUERYCAP = (const_IOC_READ                      << const_IOC_DIRSHIFT)  |
	                        (uintptr('V')                        << const_IOC_TYPESHIFT) |
	                        (0                                   << const_IOC_NRSHIFT)   |
	                        (unsafe.Sizeof(internalCapability{}) << const_IOC_SIZESHIFT)

	// A Golang conversion of the following C code:
	//
	// #define  VIDIOC_RESERVED   _IO('V',  1)
	const_VIDIOC_RESERVED = (const_IOC_NONE                      << const_IOC_DIRSHIFT)  |
	                        (uintptr('V')                        << const_IOC_TYPESHIFT) |
	                        (1                                   << const_IOC_NRSHIFT)   |
	                        (0                                   << const_IOC_SIZESHIFT)

	// A Golang conversion of the following C code:
	//
	// #define  VIDIOC_ENUM_FMT _IOWR('V',  2, struct v4l2_fmtdesc)
	const_VIDIOC_ENUM_FMT = ((const_IOC_READ | const_IOC_WRITE)         << const_IOC_DIRSHIFT)  |
	                        (uintptr('V')                               << const_IOC_TYPESHIFT) |
	                        (2                                          << const_IOC_NRSHIFT)   |
	                        (unsafe.Sizeof(internalFormatDescription{}) << const_IOC_SIZESHIFT)

	// A Golang conversion of the following C code:
	//
	// #define VIDIOC_ENUM_FRAMESIZES  _IOWR('V', 74, struct v4l2_frmsizeenum)
	const_VIDIOC_ENUM_FRAMESIZES = ((const_IOC_READ | const_IOC_WRITE) << const_IOC_DIRSHIFT)  |
	                               (uintptr('V')                       << const_IOC_TYPESHIFT) |
	                               (74                                 << const_IOC_NRSHIFT)   |
	                               (unsafe.Sizeof(FrameSize{})         << const_IOC_SIZESHIFT)
)

