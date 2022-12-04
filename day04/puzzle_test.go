package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle01(t *testing.T) {
	assert.Equal(t, 2, Puzzle01("./example.input"))

	t.Log("puzzle 01:", Puzzle01("./puzzle.input"))
}

func TestContains(t *testing.T) {
	assert.False(t, contains([]int{2, 4}, []int{6, 8}))
	assert.False(t, contains([]int{6, 8}, []int{2, 4}))

	assert.True(t, contains([]int{2, 8}, []int{3, 7}))

	assert.True(t, contains([]int{4, 6}, []int{6, 6}))
}

func TestPuzzle02(t *testing.T) {
	assert.Equal(t, 4, Puzzle02("./example.input"))

	t.Log("puzzle 02:", Puzzle02("./puzzle.input"))
}

func TestOverlap(t *testing.T) {
	assert.False(t, overlaps([]int{2, 4}, []int{6, 8}))
	assert.False(t, overlaps([]int{2, 3}, []int{4, 5}))
	assert.True(t, overlaps([]int{5, 7}, []int{7, 9}))
}
