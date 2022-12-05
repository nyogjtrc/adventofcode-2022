package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle01(t *testing.T) {
	exampleStacks := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	assert.Equal(t, "CMZ", Puzzle01(exampleStacks, "./example.input"))

	puzzleStacks := [][]string{
		{"R", "P", "C", "D", "B", "G"},
		{"H", "V", "G"},
		{"N", "S", "Q", "D", "J", "P", "M"},
		{"P", "S", "L", "G", "D", "C", "N", "M"},
		{"J", "B", "N", "C", "P", "F", "L", "S"},
		{"Q", "B", "D", "Z", "V", "G", "T", "S"},
		{"B", "Z", "M", "H", "F", "T", "Q"},
		{"C", "M", "D", "B", "F"},
		{"F", "C", "Q", "G"},
	}
	t.Log("puzzle 01:", Puzzle01(puzzleStacks, "./puzzle.input"))
}

func TestMoveCrate(t *testing.T) {
	assert.Equal(t, [][]string{
		{"Z", "N", "D"},
		{"M", "C"},
		{"P"},
	}, move([][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}, 1, 1, 0))

	assert.Equal(t, [][]string{
		{},
		{"M", "C"},
		{"P", "D", "N", "Z"},
	}, move([][]string{
		{"Z", "N", "D"},
		{"M", "C"},
		{"P"},
	}, 3, 0, 2))
}

func TestPuzzle02(t *testing.T) {
	exampleStacks := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	assert.Equal(t, "MCD", Puzzle02(exampleStacks, "./example.input"))

	puzzleStacks := [][]string{
		{"R", "P", "C", "D", "B", "G"},
		{"H", "V", "G"},
		{"N", "S", "Q", "D", "J", "P", "M"},
		{"P", "S", "L", "G", "D", "C", "N", "M"},
		{"J", "B", "N", "C", "P", "F", "L", "S"},
		{"Q", "B", "D", "Z", "V", "G", "T", "S"},
		{"B", "Z", "M", "H", "F", "T", "Q"},
		{"C", "M", "D", "B", "F"},
		{"F", "C", "Q", "G"},
	}
	t.Log("puzzle 02:", Puzzle02(puzzleStacks, "./puzzle.input"))
}

func TestPickCrate(t *testing.T) {
	assert.Equal(t, [][]string{
		{"Z", "N", "D"},
		{"M", "C"},
		{"P"},
	}, pick([][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}, 1, 1, 0))

	assert.Equal(t, [][]string{
		{},
		{"M", "C"},
		{"P", "Z", "N", "D"},
	}, pick([][]string{
		{"Z", "N", "D"},
		{"M", "C"},
		{"P"},
	}, 3, 0, 2))
}
