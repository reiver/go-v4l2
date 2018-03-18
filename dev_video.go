package v4l2

// Someone claimed that the V4L2 documentation says that there can only be (a maximum of) 64 allowed devices with the prefix "/dev/video",
// and starting with "/dev/video0"
//
// Thus, these constants represent the paths to these devices. Such that:
//
// • v4l2.Video0 = "/dev/video0"
//
// • v4l2.Video1 = "/dev/video1"
//
// • v4l2.Video2 = "/dev/video2"
//
// …
//
// • v4l2.Video63 = "/dev/video63"
//
// Use these with the v4l2.Device.Open() method or the v4l2.Open() func.
//
// For example:
//
//	device, err := v4l2.Open(v4l2.Video0)
//
// Or, alternatively, for example:
//
//	var device v4l2.Device
//	
//	err := device.Open(v4l2.Video0)
const (
	Video0 = "/dev/video0"
	Video1 = "/dev/video1"
	Video2 = "/dev/video2"
	Video3 = "/dev/video3"
	Video4 = "/dev/video4"
	Video5 = "/dev/video5"
	Video6 = "/dev/video6"
	Video7 = "/dev/video7"
	Video8 = "/dev/video8"
	Video9 = "/dev/video9"

	Video10 = "/dev/video10"
	Video11 = "/dev/video11"
	Video12 = "/dev/video12"
	Video13 = "/dev/video13"
	Video14 = "/dev/video14"
	Video15 = "/dev/video15"
	Video16 = "/dev/video16"
	Video17 = "/dev/video17"
	Video18 = "/dev/video18"
	Video19 = "/dev/video19"

	Video20 = "/dev/video20"
	Video21 = "/dev/video21"
	Video22 = "/dev/video22"
	Video23 = "/dev/video23"
	Video24 = "/dev/video24"
	Video25 = "/dev/video25"
	Video26 = "/dev/video26"
	Video27 = "/dev/video27"
	Video28 = "/dev/video28"
	Video29 = "/dev/video29"

	Video30 = "/dev/video30"
	Video31 = "/dev/video31"
	Video32 = "/dev/video32"
	Video33 = "/dev/video33"
	Video34 = "/dev/video34"
	Video35 = "/dev/video35"
	Video36 = "/dev/video36"
	Video37 = "/dev/video37"
	Video38 = "/dev/video38"
	Video39 = "/dev/video39"

	Video40 = "/dev/video40"
	Video41 = "/dev/video41"
	Video42 = "/dev/video42"
	Video43 = "/dev/video43"
	Video44 = "/dev/video44"
	Video45 = "/dev/video45"
	Video46 = "/dev/video46"
	Video47 = "/dev/video47"
	Video48 = "/dev/video48"
	Video49 = "/dev/video49"

	Video50 = "/dev/video50"
	Video51 = "/dev/video51"
	Video52 = "/dev/video52"
	Video53 = "/dev/video53"
	Video54 = "/dev/video54"
	Video55 = "/dev/video55"
	Video56 = "/dev/video56"
	Video57 = "/dev/video57"
	Video58 = "/dev/video58"
	Video59 = "/dev/video59"

	Video60 = "/dev/video60"
	Video61 = "/dev/video61"
	Video62 = "/dev/video62"
	Video63 = "/dev/video63"
)
