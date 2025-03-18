package part2

import "testing"

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
			file: "../input.txt",
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
