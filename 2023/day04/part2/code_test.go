package part1

import "testing"

func Test_ScratchedCards(t *testing.T) {
	testCases := []struct {
		desc string
		text string
		want int
	}{
		{
			desc: "input testing",
			text: "input_test.txt",
			want: 30,
		},
		{
			desc: "input",
			text: "../input.txt",
			want: 5659035,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := ScratchedCards(tC.text)
			if got != tC.want {
				t.Fatalf("got: %d, want: %d", got, tC.want)
			}
		})
	}
}
