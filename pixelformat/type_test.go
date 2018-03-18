package v4l2_pixelformat

import (
	"testing"
)

func TestFourCC(t *testing.T) {

	tests := []struct{
		Value    string
		Expected uint32
	}{
		{
			Value:    "",
			Expected: 0x20202020,
		},



		{
			Value:    " ",
			Expected: 0x20202020,
		},
		{
			Value:    "  ",
			Expected: 0x20202020,
		},
		{
			Value:    "   ",
			Expected: 0x20202020,
		},
		{
			Value:    "    ",
			Expected: 0x20202020,
		},



		{
			Value:    "a",
			Expected: 0x20202061,
		},
		{
			Value:    "ab",
			Expected: 0x20206261,
		},
		{
			Value:    "abc",
			Expected: 0x20636261,
		},
		{
			Value:    "abcd",
			Expected: 0x64636261,
		},
		{
			Value:    "abcde",
			Expected: 0x64636261,
		},
		{
			Value:    "abcdef",
			Expected: 0x64636261,
		},



		{
			Value:    "A",
			Expected: 0x20202041,
		},
		{
			Value:    "AB",
			Expected: 0x20204241,
		},
		{
			Value:    "ABC",
			Expected: 0x20434241,
		},
		{
			Value:    "ABCD",
			Expected: 0x44434241,
		},
		{
			Value:    "ABCDE",
			Expected: 0x44434241,
		},
		{
			Value:    "ABCDEF",
			Expected: 0x44434241,
		},
		{
			Value:    "ABCDEFG",
			Expected: 0x44434241,
		},



		{
			Value:    "A",
			Expected: 0x20202041,
		},



		{
			Value:    "GREY",
			Expected: 0x59455247,
		},
		{
			Value:    "RGB1",
			Expected: 0x31424752,
		},
		{
			Value:    "H264",
			Expected: 0x34363248,
		},
		{
			Value:    "MJPG",
			Expected: 0x47504a4d,
		},
		{
			Value:    "MPEG",
			Expected: 0x4745504d,
		},
		{
			Value:    "PAL8",
			Expected: 0x384c4150,
		},
		{
			Value:    "VP80",
			Expected: 0x30385056,
		},
		{
			Value:    "XVID",
			Expected: 0x44495658,
		},



		{
			Value:    "UV8",
			Expected: 0x20385655,
		},
		{
			Value:    "UV8 ",
			Expected: 0x20385655,
		},



		{
			Value:    "Y04",
			Expected: 0x20343059,
		},
		{
			Value:    "Y04 ",
			Expected: 0x20343059,
		},
	}


	for testNumber, test := range tests {

		pixelFormat := FourCC(test.Value)

		if actual, expected := pixelFormat.value, test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %#x (%q), but actually got %#x (%q).", testNumber, expected, Type{expected}.String(), actual, Type{actual}.String())
			continue
		}
	}
}
