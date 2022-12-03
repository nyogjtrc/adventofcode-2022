package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle01(t *testing.T) {
	assert.Equal(t, 15, Puzzle01("./example.input"))

	t.Log(Puzzle01("./puzzle.input"))
}

func TestPuzzle02(t *testing.T) {
	assert.Equal(t, 12, Puzzle02("./example.input"))

	t.Log(Puzzle02("./puzzle.input"))
}
