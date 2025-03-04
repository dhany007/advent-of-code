package day02

import "testing"

func Test_SumValidBag(t *testing.T) {
	testCases := []struct {
		desc string
		file string
		want int64
	}{
		{
			desc: "input_test",
			file: "input_test.txt",
			want: 2286,
		},
		{
			desc: "input",
			file: "input.txt",
			want: 72513,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := SumValidBag(tC.file)
			if got != tC.want {
				t.Fatalf("got: %d, want: %d", got, tC.want)
			}
		})
	}
}
