package day10

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle01(t *testing.T) {
	assert.Equal(t, 13140, Puzzle01("example.input"))

	t.Log("puzzle 01:", Puzzle01("./puzzle.input"))
}

func TestCpu(t *testing.T) {
	c := NewCpu()

	c.exec(NewNoop())
	c.Tick()
	assert.Equal(t, 1, c.cycle)
	assert.Equal(t, 1, c.register)

	c.exec(NewAddx(3))
	c.Tick()
	assert.Equal(t, 2, c.cycle)
	assert.Equal(t, 1, c.register)

	c.Tick()
	assert.Equal(t, 3, c.cycle)
	assert.Equal(t, 4, c.register)

	c.exec(NewAddx(-5))
	c.Tick()
	assert.Equal(t, 4, c.cycle)
	assert.Equal(t, 4, c.register)

	c.Tick()
	assert.Equal(t, 5, c.cycle)
	assert.Equal(t, -1, c.register)
}

func TestCRT(t *testing.T) {
	c := NewCRT()

	c.Exec(NewAddx(15))
	c.Tick()
	assert.Equal(t, []int{1}, c.Row)

	c.Tick()
	assert.Equal(t, []int{1, 1}, c.Row)

	c.Exec(NewAddx(-11))
	c.Tick()
	assert.Equal(t, []int{1, 1, 0}, c.Row)

	c.Tick()
	assert.Equal(t, []int{1, 1, 0, 0}, c.Row)
}

func TestPuzzle02(t *testing.T) {
	assert.Equal(t, `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`, Puzzle02("example.input"))

	fmt.Printf("puzzle 02:\n%s", Puzzle02("./puzzle.input"))
}
