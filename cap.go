package v4l2

// V4L2 (Video4Linux version 2) capability constants.
//
// Use these with the v4l2.Device.HasCapability() method.
//
// For example:
//
//	var device v4l2.Device
//	
//	// ...
//	
//	if ! device.HasCapability(v4l2.CapabilityVideoCapture) {
//		//@TODO
//	}
//
// These are mimicked from:
// /usr/include/linux/videodev2.h
const (

	CapabilityVideoCapture            = 0x00000001 // (V4L2_CAP_VIDEO_CAPTURE)        Is a video capture device
	CapabilityVideoOutput             = 0x00000002 // (V4L2_CAP_VIDEO_OUTPUT)         Is a video output device
	CapabilityVideoOverlay            = 0x00000004 // (V4L2_CAP_VIDEO_OVERLAY)        Can do video overlay
	CapabilityVbiCapture              = 0x00000010 // (V4L2_CAP_VBI_CAPTURE)          Is a raw VBI capture device
	CapabilityVbiOutput               = 0x00000020 // (V4L2_CAP_VBI_OUTPUT)           Is a raw VBI output device
	CapabilitySlicedVbiCapture        = 0x00000040 // (V4L2_CAP_SLICED_VBI_CAPTURE)   Is a sliced VBI capture device
	CapabilitySlicedVbiOutput         = 0x00000080 // (V4L2_CAP_SLICED_VBI_OUTPUT)    Is a sliced VBI output device
	CapabilityRDSCapture              = 0x00000100 // (V4L2_CAP_RDS_CAPTURE)          RDS data capture
	CapabilityVideoOutputOverlay      = 0x00000200 // (V4L2_CAP_VIDEO_OUTPUT_OVERLAY) Can do video output overlay
	CapabilityHardwareFrequencySeek   = 0x00000400 // (V4L2_CAP_HW_FREQ_SEEK)         Can do hardware frequency seek
	CapabilityRDSOutput               = 0x00000800 // (V4L2_CAP_RDS_OUTPUT)           Is an RDS encoder

	CapabilityVideoCaptureMultiplanar = 0x00001000 // (V4L2_CAP_VIDEO_CAPTURE_MPLANE) Is a video capture device that supports multiplanar formats
	CapabilityVideoOutputMultiplanar  = 0x00002000 // (V4L2_CAP_VIDEO_OUTPUT_MPLANE)  Is a video output device that supports multiplanar formats
	CapabilityVideoM2MMultiplanar     = 0x00004000 // (V4L2_CAP_VIDEO_M2M_MPLANE)     Is a video mem-to-mem device that supports multiplanar formats
	CapabilityVideoM2M                = 0x00008000 // (V4L2_CAP_VIDEO_M2M)            Is a video mem-to-mem device

	CapabilityTuner                   = 0x00010000 // (V4L2_CAP_TUNER)                has a tuner
	CapabilityAudio                   = 0x00020000 // (V4L2_CAP_AUDIO)                has audio support
	CapabilityRadio                   = 0x00040000 // (V4L2_CAP_RADIO)                is a radio device
	CapabilityModulator               = 0x00080000 // (V4L2_CAP_MODULATOR)            has a modulator

	CapabilityReadWrite               = 0x01000000 // (V4L2_CAP_READWRITE)            read/write systemcalls
	CapabilityAsyncIO                 = 0x02000000 // (V4L2_CAP_ASYNCIO)              async I/O
	CapabilityStreaming               = 0x04000000 // (V4L2_CAP_STREAMING)            streaming I/O ioctls

	CapabilityDeviceCaps              = 0x80000000 // (V4L2_CAP_DEVICE_CAPS)         sets device capabilities field
)
