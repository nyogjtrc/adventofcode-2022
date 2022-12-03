package day03

import (
	"log"

	"github.com/nyogjtrc/adventofcode-2022/input"
)

func Puzzle01(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		part1 := scanner.Text()[:len(scanner.Text())/2]
		part2 := scanner.Text()[len(scanner.Text())/2:]

		sameCompartment := lookForSameCompartment(part1, part2)

		answer = answer + priority(sameCompartment)
	}

	return
}

func lookForSameCompartment(part1, part2 string) byte {
	for k := range part1 {
		for j := range part2 {
			if part1[k] == part2[j] {
				return part1[k]
			}
		}
	}
	return 0
}

func priority(i byte) int {
	// upper
	if i > 64 && i < 91 {
		return int(i) - 64 + 26
	}

	// lower
	if i > 96 && i < 123 {
		return int(i) - 96
	}

	return 0
}

func Puzzle02(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	group := []string{}
	for scanner.Scan() {
		group = append(group, scanner.Text())

		if len(group) == 3 {
			item := lookForSameItemWithGroup(group)

			answer = answer + priority(item)

			group = []string{}
		}
	}

	return
}

func lookForSameItemWithGroup(group []string) byte {
	for k := range group[0] {
		for j := range group[1] {
			if group[0][k] == group[1][j] {
				for l := range group[2] {
					if group[0][k] == group[2][l] {
						return group[0][k]
					}
				}
			}
		}
	}
	return 0
}
