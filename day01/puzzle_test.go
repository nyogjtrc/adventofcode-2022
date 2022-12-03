package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01Puzzle01(t *testing.T) {
	assert.Equal(t, 24000, Puzzle01("./example.input"))

	t.Log(Puzzle01("./puzzle.input"))
}

func TestDay01Puzzle02(t *testing.T) {
	assert.Equal(t, 45000, Puzzle02("./example.input"))

	t.Log(Puzzle02("./puzzle.input"))
}
