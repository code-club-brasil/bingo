package bingo

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

const (
	sortedColor = "\033[1;36m%s\033[0m\t"
	totalBalls  = 75
)

// Random is global variable used to control Seed value
var Random = rand.New(rand.NewSource(time.Now().UnixNano()))

// NewCard takes an io.Writer and prints a new bingo Card
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

// PrintBoard takes map of integer containing an empty struct to control balls already revealed and an integer with the last ball revealed in order to print the main board.
func PrintBoard(revealed map[int]struct{}, lastRevealed int) {
	pace := 0
	headers := [10]int{1, 2, 16, 17, 31, 32, 46, 47, 61, 62}
	fmt.Println("B                I                N                G                O")
	for x := 0; x < 8; x++ {
		for y := 0; y < 10; y++ {
			// skip 16, 31, 46, 61 and 76
			if x == 7 && (headers[y]+pace)%15 == 1 {
				fmt.Print("        ")
				continue
			}
			_, ok := revealed[headers[y]+pace]
			if !ok {
				fmt.Print(headers[y]+pace, "\t")
				continue
			}
			fmt.Printf(sortedColor, strconv.Itoa(headers[y]+pace))
		}
		fmt.Println()
		pace += 2
	}
	fmt.Printf("Ultimo numero sorteado: %d\n", lastRevealed)
}

// Run takes no parameter and run until the game is finished
func Run() {
	var bingo bool
	randomCard := Random.Perm(totalBalls)
	revealed := make(map[int]struct{}, 1)
	var enter string
	for !bingo {
		n := randomCard[0]
		randomCard = randomCard[1:]
		revealed[n] = struct{}{}
		PrintBoard(revealed, n)
		fmt.Printf("Alguem ganhou? [s|n]")
		fmt.Scanln(&enter)
		if enter == "s" {
			bingo = true
		}
	}
}
