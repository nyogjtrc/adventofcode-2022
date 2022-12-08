package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle01(t *testing.T) {
	assert.Equal(t, 21, Puzzle01("example.input"))

	assert.True(t, isVisible(5, []int{}, 0))
	assert.True(t, isVisible(5, []int{}, 4))
	assert.True(t, isVisible(5, []int{}, 5))
	assert.True(t, isVisible(5, []int{}, 21))
	assert.True(t, isVisible(5, []int{}, 24))

	t.Log("puzzle 01:", Puzzle01("./puzzle.input"))
}

func TestPuzzle02(t *testing.T) {
	assert.Equal(t, 8, Puzzle02("example.input"))

	assert.Equal(t, 0, upScore([]int{}, 5, 0, 0, 0))

	t.Log("puzzle 02:", Puzzle02("./puzzle.input"))
}
