package bingo

import (
	"fmt"
	"io"
)

func CardPrinter(balls []int, w io.Writer) {
	for i, card := range balls {
		if i%5 == 0 {
			fmt.Fprintln(w)
		}
		if i == 12 {
			fmt.Fprint(w, "\033[1;31m")
			fmt.Fprint(w, "   *")
			fmt.Fprint(w, "\033[0m")
			continue
		}
		fmt.Fprintf(w, "%4d", card)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w)	
}

func NewCard() []int {
	randomBalls := Random.Perm(75)
	card := randomBalls[:25]
	for i, v := range card {
		card[i] = v+1
	}
	return card
}