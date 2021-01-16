package bingo_test

import (
	"bingo"
	"bytes"
	"testing"
)

func TestNewCard(t *testing.T) {
	t.Parallel()

	wantLen := 106
	var buffer bytes.Buffer
	bingo.NewCard(&buffer)
	if buffer.Len() != wantLen {
		t.Errorf("want %d, got %d", wantLen, buffer.Len())
	}
}
