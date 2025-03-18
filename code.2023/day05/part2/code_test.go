package part2

import "testing"

func TestFertilizer(t *testing.T) {
	testcases := []struct {
		filename string
		result   int64
	}{
		{
			filename: "../input_test.txt",
			result:   46,
		},
		{
			filename: "../input.txt",
			result:   340994526,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.filename, func(t *testing.T) {
			result := GetLocationGarden(tc.filename)
			if result != tc.result {
				t.Errorf("got: %d, want %d", result, tc.result)
			}
		})
	}
}
