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

That example opens up the V4L2 device at "/dev/video0" on the file system, and displays some basic information about the device.

(Of course, we could have opened one of the other V4L2 devices. Such as: v4l2.Video1, v4l2.Video2, ..., or v4l2.Video63.)

Continuing this same example:

	formats, err := device.Formats()
	if nil != err {
		return err
	}
	defer formats.Close()
	
	var format v4l2.Format
	for formats.Next() {
		
		if err := formats.Decode(&format); nil != err {
			return err
		}
		
		fmt.Printf("[format description] %q (%q) {compressed=%t} {emulated=%t} \n",
			format.Description(),
			format.PixelFormat(),
			format.HasFlags(v4l2.FormatFlagCompressed),
			format.HasFlags(v4l2.FormatFlagEmulated),
		)
		
		
		//@TODO
	}
	if err := formats.Err(); nil != err {
		return err
	}

Here we have iterating through the formats that are supported for this device.

Extending that last code block, to fill in that "//@TODO", we can iterate through the frame sizes, for each format, we the following:

		frameSizes, err := format.FrameSizes()
		if nil != err {
			return return
		}
		defer frameSizes.Close()
		
		var frameSize v4l2.FrameSize
		for frameSizes.Next() {
			
			if err := frameSizes.Decode(&frameSize); nil != err {
				return err
			}
			
			casted, err := frameSize.Cast()
			if nil != err {
				return err
			}
			
			switch t := casted.(type) {
			case v4l2.FrameSizeDiscrete:
				fmt.Printf("\t [frame size discrete] pixel_format=%q, width=%d, height=%d \n",
					t.PixelFormat(),
					t.Width,
					t.Height,
				)
			case v4l2.FrameSizeContinuous:
				fmt.Printf("\t [frame size continuous] pixel_format=%q, min_width=%d, max_width=%d, min_height=%d, max_height=% \n",
					t.PixelFormat(),
					t.MinWidth,
					t.MaxWidth,
					t.MinHeight,
					t.MaxHeight,
				)
			case v4l2.FrameSizeStepwise:
				fmt.Printf("\t [frame size stepwise] pixel_format=%q, min_width=%d, max_width=%d, min_height=%d, max_height=% \n",
					t.PixelFormat(),
					t.MinWidth,
					t.MaxWidth,
					t.MinHeight,
					t.MaxHeight,
				)
			default:
				return err
			}

		}
		if err := frameSizes.Err(); nil != err {
			return err
		}


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
