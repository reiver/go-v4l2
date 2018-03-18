package v4l2

// BusInfo returns the location of the V4L2 device in the system.
//
// Example values that BusInfo might return are:
//
// • "usb-0000:00:1a.0-1.5"
//
// • "PCI:0000:05:06.0"
//
// The information returned by BusInfo is intended for users to be able to distinguish between multiple identical devices.
//
// For example, if you have 2 of the exact same (hardware) devices that plug into the USB ports of your computer,
// and capture HDMI video, then the driver, the card, all the capabilities, etc will be the same for these 2 devices.
// However, what BusInfo returns will be different.
//
// In the cases wwere no such information is available for the BusInfo (at the Linux or driver level) then BusInfo will
// be the count of the devices controlled by the driver ("platform:vivi-000").
//
// What is returned from BusInfo will start with:
//
// • "PCI:" for PCI boards,
// • "PCIe:" for PCI Express boards,
// • "usb-" for USB devices,
// • "I2C:" for i2c devices,
// • "ISA:" for ISA devices,
// • "parport" for parallel port devices, and
// • "platform:" for platform devices.
//
// What BusInfo returns is called "bus_info" in the Linux V4L2 documention, which is found in the Linux "v4l2_capability" data structure.
func (receiver Device) BusInfo() (string, error) {
	if !receiver.opened {
		return "", errNotOpen
	}
	if !receiver.capQueried {
		return "", errInternalError
	}

	return receiver.cap.BusInfo(), nil
}

// MustBusInfo is like BusInfo, except it panic()s if there is an error.
func (receiver Device) MustBusInfo() string {
	datum, err := receiver.BusInfo()
	if nil != err {
		panic(err)
	}

	return datum
}
