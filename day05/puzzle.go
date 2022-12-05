package day05

import (
	"log"
	"strconv"
	"strings"

	"github.com/nyogjtrc/adventofcode-2022/input"
)

func Puzzle01(stacks [][]string, filename string) (answer string) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		amount, _ := strconv.Atoi(s[1])
		from, _ := strconv.Atoi(s[3])
		to, _ := strconv.Atoi(s[5])

		stacks = move(stacks, amount, from-1, to-1)
	}

	return stacksTop(stacks)
}

func move(stacks [][]string, amount int, from int, to int) [][]string {
	for i := 0; i < amount; i++ {
		last := len(stacks[from]) - 1
		c := stacks[from][last]
		stacks[from] = stacks[from][:last]
		stacks[to] = append(stacks[to], c)
	}
	return stacks
}

func stacksTop(stacks [][]string) (top string) {
	for k := range stacks {
		top = top + stacks[k][len(stacks[k])-1]
	}
	return
}

func Puzzle02(stacks [][]string, filename string) (answer string) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		amount, _ := strconv.Atoi(s[1])
		from, _ := strconv.Atoi(s[3])
		to, _ := strconv.Atoi(s[5])

		stacks = pick(stacks, amount, from-1, to-1)
	}

	return stacksTop(stacks)
}

func pick(stacks [][]string, amount int, from int, to int) [][]string {
	//last := len(stacks[from]) - 1
	c := stacks[from][len(stacks[from])-amount : len(stacks[from])]
	stacks[from] = stacks[from][:len(stacks[from])-amount]
	stacks[to] = append(stacks[to], c...)

	return stacks
}
