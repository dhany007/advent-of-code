package part1

import "testing"

func Test_CalcRac(t *testing.T) {
	testCases := []struct {
		name     string
		times    int
		distance int
		want     int
	}{
		{
			name:     "race1",
			times:    7,
			distance: 9,
			want:     4,
		},
		{
			name:     "race2",
			times:    15,
			distance: 40,
			want:     8,
		},
		{
			name:     "race3",
			times:    30,
			distance: 200,
			want:     9,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcRac(tt.times, tt.distance)
			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRace(t *testing.T) {
	testCases := []struct {
		name string
		file string
		want int
	}{
		{
			name: "testing",
			file: "input_test.txt",
			want: 288,
		},
		{
			name: "testcases",
			file: "../input.txt",
			want: 227850,
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
