package v4l2

// These constants were mimicked from:
// /include/uapi/asm-generic/ioctl.h
const (
	const_IOC_NRBITS   = 8
	const_IOC_TYPEBITS = 8

	const_IOC_SIZEBITS = 14
	const_IOC_DIRBITS  = 2

	const_IOC_NRMASK   = ((1 << const_IOC_NRBITS)-1)
	const_IOC_TYPEMASK = ((1 << const_IOC_TYPEBITS)-1)
	const_IOC_SIZEMASK = ((1 << const_IOC_SIZEBITS)-1)
	const_IOC_DIRMASK  = ((1 << const_IOC_DIRBITS)-1)

	const_IOC_NONE  = 0
	const_IOC_WRITE = 1
	const_IOC_READ  = 2

	const_IOC_NRSHIFT   = 0
	const_IOC_TYPESHIFT = (const_IOC_NRSHIFT   + const_IOC_NRBITS)
	const_IOC_SIZESHIFT = (const_IOC_TYPESHIFT + const_IOC_TYPEBITS)
	const_IOC_DIRSHIFT  = (const_IOC_SIZESHIFT + const_IOC_SIZEBITS)
)

