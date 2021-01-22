package bingo

import (
	"fmt"
	"strconv"
)

// PrintBoard takes map of integer containing an empty struct to control balls already revealed and an integer with the last ball revealed in order to print the main board.
func PrintBoard(revealed map[int]struct{}, lastRevealed int) {
	pace := 0
	headers := [10]int{1, 2, 16, 17, 31, 32, 46, 47, 61, 62}
	fmt.Printf(titleColor, "B                I                N                G                O")
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