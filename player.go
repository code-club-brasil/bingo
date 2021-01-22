package bingo

import (
	"os"
	"fmt"
)

var totalPlayers int
var players []Player

type Player struct {
	id int
	name string	
	card []int
}

func SetupPlayers() {
	i := 0
	fmt.Print("Total de jogadores? ")
	fmt.Scanf("%d", &i)
	players = make([]Player, i)
	for x := 0; x < i; x++ {
		var name string
		fmt.Printf("Nome jogagor #%d? ", x+1)
		fmt.Scanf("%s", &name)
		card := NewCard()
		player := NewPlayer(name, card)
		PlayerPrinter(player)
		players[x] = player
		fmt.Println()
	}
}

func NewPlayer(name string, card []int) Player {
	totalPlayers++
	p := Player{totalPlayers, name, card}
	return p
}

func PlayerPrinter(p Player) {
	fmt.Printf("Jogador #%d > %s\n", p.id, p.name)
	CardPrinter(p.card, os.Stdout)
}