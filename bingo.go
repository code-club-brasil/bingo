package bingo

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	sortedColor = "\033[1;36m%s\033[0m\t"
	titleColor = "\033[1;31m%s\033[0m\n"
	totalBalls  = 75
)

// Random is global variable used to control Seed value
var Random = rand.New(rand.NewSource(time.Now().UnixNano()))

func Bingo(players []Player, revealed map[int]struct{}) bool {
	// for i, v := range players {
	// 	fmt.Println(i)
	// 	fmt.Println(v)
	// }
	return false
}

// Run takes no parameter and run until the game is finished
func Run() {
	randomCard := Random.Perm(totalBalls)
	revealed := make(map[int]struct{}, 1)
	SetupPlayers()
	var enter string
	var totalRevealed int
	for i := 0; i < totalBalls; i++ {
		n := randomCard[0] + 1
		randomCard = randomCard[1:]
		revealed[n] = struct{}{}
		PrintBoard(revealed, n)
		totalRevealed++
		fmt.Printf("Rodada #%d\n", totalRevealed)
		fmt.Printf("Alguem ganhou? [s|n] ")
		fmt.Scanf("%s", &enter)
		if totalRevealed < 24 {
			continue
		}
		Bingo(players, revealed)
	}
}
