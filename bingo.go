package bingo

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const cols, rows, totalCardNumbers, totalBalls = 5, 5, 24, 75

var Random = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewCard(w io.Writer) {
	randomCard := Random.Perm(75)
	for i, card := range randomCard[:25] {
		if i%5 == 0 {
			fmt.Fprintln(w)
		}
		if i == 12 {
			fmt.Fprint(w, "  XX")
			continue
		}
		fmt.Fprintf(w, "%4d", card+1)
	}
	fmt.Fprintln(w)
}

func Raffle() {
	// B    I      N      G      O
	// 1 2  16 17  31 32  46 47  61 62
	// 3 4  18 19
	// 5 6  20 21
	// 7 8
	// ...
	// 15
	var bingo bool
	randomCard := Random.Perm(totalBalls)
	for !bingo {
		n := randomCard[0]
		randomCard = randomCard[1:]
		fmt.Println(len(randomCard))
		fmt.Println(n)
		os.Exit(0)
	}
	fmt.Println("      B       I     N      G     O")
	for x := 1; x <= 14; x += 2 {
		b1 := x
		b2 := x + 1
		i1 := b1 + 15
		i2 := b2 + 15
		n1 := i1 + 15
		n2 := i2 + 15
		g1 := n1 + 15
		g2 := n2 + 15
		o1 := g1 + 15
		o2 := g2 + 15
		fmt.Printf("%4d %4d %4d %4d %4d %4d %4d %4d %4d %4d\n", b1, b2, i1, i2, n1, n2, g1, g2, o1, o2)
	}
	fmt.Printf("%6d %10d %9d %10d %8d\n", 15, 30, 45, 60, 75)
}
