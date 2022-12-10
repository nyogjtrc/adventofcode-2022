package day10

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/nyogjtrc/adventofcode-2022/input"
)

type instruction struct {
	cycle int
	value int
}

func NewNoop() instruction {
	return instruction{cycle: 1}
}

func NewAddx(value int) instruction {
	return instruction{cycle: 2, value: value}
}

type cpu struct {
	register    int
	cycle       int
	instruction *instruction
}

func NewCpu() *cpu {
	return &cpu{
		register: 1,
		cycle:    0,
	}
}

func (c *cpu) Tick() int {
	c.cycle++
	signal := c.signal()
	c.instruction.cycle--
	if c.instruction.cycle == 0 {
		c.register += c.instruction.value
		c.instruction = nil
	}

	return signal
}

func (c *cpu) exec(ins instruction) {
	c.instruction = &ins
}

func (c *cpu) signal() int {
	return c.cycle * c.register
}

func Puzzle01(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	inss := []instruction{}
	for scanner.Scan() {
		ins := parser(scanner.Text())
		inss = append(inss, ins)
	}

	c := NewCpu()
	insCounter := 0

	for {
		if c.instruction == nil {
			if insCounter == len(inss) {
				break
			}
			c.exec(inss[insCounter])
			insCounter++
		}
		signal := c.Tick()

		if math.Mod(float64(c.cycle), 40) == 20 {
			answer += signal
			fmt.Println(c.cycle, signal)
		}
		//fmt.Println(c.cycle, c.register, c.instruction)
	}
	return
}

func parser(cmd string) instruction {
	ss := strings.Split(cmd, " ")
	if len(ss) == 1 {
		return NewNoop()
	}

	val, _ := strconv.Atoi(ss[1])
	return NewAddx(val)
}

type CRT struct {
	Cycle       int
	Register    int
	instruction *instruction
	Row         []int
}

func NewCRT() *CRT {
	return &CRT{
		Cycle:    0,
		Register: 0,
	}
}

func (c *CRT) Exec(ins instruction) {
	c.instruction = &ins
}

func (c *CRT) Tick() {
	// drawn
	positionF := math.Mod(float64(c.Cycle), 40)
	position := int(positionF)

	if position >= c.Register && position < c.Register+3 {
		c.Row = append(c.Row, 1)
	} else {
		c.Row = append(c.Row, 0)
	}

	// exec
	c.Cycle++
	c.instruction.cycle--
	if c.instruction.cycle == 0 {
		c.Register += c.instruction.value
		c.instruction = nil
	}
}

func Puzzle02(filename string) (answer string) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	inss := []instruction{}
	for scanner.Scan() {
		ins := parser(scanner.Text())
		inss = append(inss, ins)
	}

	c := NewCRT()
	insCounter := 0

	for {
		if c.instruction == nil {
			if insCounter == len(inss) {
				break
			}
			c.Exec(inss[insCounter])
			insCounter++
		}
		c.Tick()
	}

	return produce(*c)
}

func produce(c CRT) (s string) {
	for k := range c.Row {
		if c.Row[k] == 0 {
			s += "."
		} else {
			s += "#"
		}

		if math.Mod(float64(k+1), 40) == 0 {
			s += "\n"
		}
	}
	return
}
