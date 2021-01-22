package bingo_test

import (
	"bingo"
	"testing"
)

func TestNewCard(t *testing.T) {
	t.Parallel()

	card := bingo.NewCard()
	if len(card) != 25 {
		t.Errorf("Wrong card length %d", len(card))
	}
}
