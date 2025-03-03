package code2023

import "testing"

func Test_ProcessCalibration(t *testing.T) {
	testCases := []struct {
		arg  string
		want int
	}{
		{
			arg:  "1abc2",
			want: 12,
		},
		{
			arg:  "pqr3stu8vwx",
			want: 38,
		},
		{
			arg:  "a1b2c3d4e5f",
			want: 15,
		},
		{
			arg:  "treb7uchet",
			want: 77,
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
			want: 142,
		},
		{
			file: "input.txt",
			want: 56049,
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
