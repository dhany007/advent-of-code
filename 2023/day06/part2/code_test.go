package part2

import "testing"

func TestRace(t *testing.T) {
	testCases := []struct {
		name string
		file string
		want int64
	}{
		{
			name: "testing",
			file: "input_test.txt",
			want: 71503,
		},
		{
			name: "testcases",
			file: "../input.txt",
			want: 42948149,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := Race(tt.file)
			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
