package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01Puzzle01(t *testing.T) {
	assert.Equal(t, 24000, Puzzle01("./day_01_example.input"))

	t.Log(Puzzle01("./day_01.input"))
}

func TestDay01Puzzle02(t *testing.T) {
	assert.Equal(t, 45000, Puzzle02("./day_01_example.input"))

	t.Log(Puzzle02("./day_01.input"))
}
