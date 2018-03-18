# go-v4l2

Package v4l2 exposes the V4L2 (Video4Linux version 2) API to Golang.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-v4l2

[![GoDoc](https://godoc.org/github.com/reiver/go-v4l2?status.svg)](https://godoc.org/github.com/reiver/go-v4l2)


## Example
```go
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
```
