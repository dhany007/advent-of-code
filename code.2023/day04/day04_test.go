package day04

import "testing"

func Test_ScratchedCards_P2(t *testing.T) {
	testCases := []struct {
		desc string
		text string
		want int
	}{
		//{
		//	desc: "input testing",
		//	text: "input_test.txt",
		//	want: 30,
		//},
		{
			desc: "input",
			text: "input.txt",
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

//
//func Test_ProcessScratchCards_P1(t *testing.T) {
//	testCases := []struct {
//		desc string
//		text string
//		want int64
//	}{
//		{
//			desc: "testing",
//			text: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
//			want: 8,
//		},
//	}
//
//	for _, tC := range testCases {
//		t.Run(tC.desc, func(t *testing.T) {
//			got := ProcessScartchcards(tC.text)
//			if got != tC.want {
//				t.Fatalf("got: %d, want: %d", got, tC.want)
//			}
//		})
//	}
//}
//
//func Test_ScratchedCards_P1(t *testing.T) {
//	testCases := []struct {
//		desc string
//		text string
//		want int64
//	}{
//		{
//			desc: "input testing",
//			text: "input_test.txt",
//			want: 13,
//		},
//		{
//			desc: "input",
//			text: "input.txt",
//			want: 24160,
//		},
//	}
//
//	for _, tC := range testCases {
//		t.Run(tC.desc, func(t *testing.T) {
//			got := ScratchedCards(tC.text)
//			if got != tC.want {
//				t.Fatalf("got: %d, want: %d", got, tC.want)
//			}
//		})
//	}
//}
