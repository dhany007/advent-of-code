package day01

import "testing"

func Test_ProcessCalibration(t *testing.T) {
	testCases := []struct {
		arg  string
		want int
	}{
		{
			arg:  "two1nine",
			want: 29,
		},
		{
			arg:  "eightwothree",
			want: 83,
		},
		{
			arg:  "abcone2threexyz",
			want: 13,
		},
		{
			arg:  "xtwone3four",
			want: 24,
		},
		{
			arg:  "4nineeightseven2",
			want: 42,
		},
		{
			arg:  "zoneight234",
			want: 14,
		},
		{
			arg:  "7pqrstsixteen",
			want: 76,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.arg, func(t *testing.T) {
			got := ProcessCalibration(tC.arg)
			if got != tC.want {
				t.Fatalf("got: %d, want: %d", got, tC.want)
			}
		})
	}
}

func Test_Calibration(t *testing.T) {
	testCases := []struct {
		file string
		want int64
	}{
		{
			file: "input_test.txt",
			want: 281,
		},
		{
			file: "input.txt",
			want: 54530,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.file, func(t *testing.T) {
			got := Calibration(tC.file)
			if got != tC.want {
				t.Fatalf("got: %d, want: %d", got, tC.want)
			}
		})
	}
}
