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
*/
package v4l2
