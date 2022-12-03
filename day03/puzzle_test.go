package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle01(t *testing.T) {
	assert.Equal(t, 157, Puzzle01("./example.input"))

	t.Log(Puzzle01("./puzzle.input"))
}

func TestPriority(t *testing.T) {
	var tests = []struct {
		expected int
		given    byte
	}{
		{16, 'p'},
		{38, 'L'},
		{42, 'P'},
		{22, 'v'},
		{20, 't'},
		{19, 's'},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, priority(tt.given), "given %s", string(tt.given))
	}
}

func TestPuzzle02(t *testing.T) {
	assert.Equal(t, 70, Puzzle02("./example.input"))

	t.Log(Puzzle02("./puzzle.input"))
}

func TestLookForSameItemWithGroup(t *testing.T) {
	assert.Equal(t, byte('r'), lookForSameItemWithGroup([]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}))

	assert.Equal(t, byte('Z'), lookForSameItemWithGroup([]string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}))
}
