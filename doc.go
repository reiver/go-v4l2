/*
Package v4l2 exposes the V4L2 (Video4Linux version 2) API to Golang.

Example:

	device, err := v4l2.Open(v4l2.Video0)
	if nil != err {
		//@TODO
		return err
	}
	defer device.Close()
	
	fmt.Printf("Driver:  %q\n", device.MustDriver())
	fmt.Printf("Card:    %q\n", device.MustCard())
	fmt.Printf("BusInfo: %q\n", device.MustBusInfo())
	fmt.Printf("Version: %q\n", device.MustVersion())
	fmt.Println()
	fmt.Printf("Has Video Capture: %v\n", device.MustHasCapability(v4l2.CapabilityVideoCapture))
	fmt.Printf("Has Streaming I/O: %v\n", device.MustHasCapability(v4l2.CapabilityStreaming))

(That example opens up the V4L2 device at "/dev/video0" on the file system, and displays some basic information about the device.)

Device Paths

One thing to be cognisant of, is that on Linux systems, V4L2 (Video4Linux version 2) will create a (special) file in the
/dev/ directory for every V4L2 device.

These devices will name path names such as:

• /dev/video0

• /dev/video1

• /dev/video2

…

• /dev/video63

You would call Open using these names directly. For example:

	device, err := v4l2.Open("/dev/video3")

However, this package also provides constants that can be used. For example:

	device, err := v4l2.Open(v4l2.Video3)

*/
package v4l2
