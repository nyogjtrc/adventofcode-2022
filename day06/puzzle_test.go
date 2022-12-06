package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle01(t *testing.T) {
	assert.Equal(t, 7, differentMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
	assert.Equal(t, 5, differentMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(t, 6, differentMarker("nppdvjthqldpwncqszvftbrmjlhg", 4))
	assert.Equal(t, 10, differentMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
	assert.Equal(t, 11, differentMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))

	t.Log("puzzle 01:", Puzzle01("./puzzle.input"))
}

func TestPuzzle02(t *testing.T) {
	assert.Equal(t, 19, differentMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	assert.Equal(t, 23, differentMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
	assert.Equal(t, 23, differentMarker("nppdvjthqldpwncqszvftbrmjlhg", 14))
	assert.Equal(t, 29, differentMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14))
	assert.Equal(t, 26, differentMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14))

	t.Log("puzzle 02:", Puzzle02("./puzzle.input"))
}
