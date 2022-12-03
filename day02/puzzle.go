package day02

import (
	"log"
	"strings"

	"github.com/nyogjtrc/adventofcode-2022/input"
)

func Puzzle01(filename string) int {
	shapeScoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	gameScoreMap := map[string]int{
		"A Y": 6,
		"B Z": 6,
		"C X": 6,
		"A X": 3,
		"B Y": 3,
		"C Z": 3,
		"A Z": 0,
		"B X": 0,
		"C Y": 0,
	}

	score := 0

	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		score = score + gameScoreMap[scanner.Text()]

		tmp := strings.Split(scanner.Text(), " ")
		score = score + shapeScoreMap[tmp[1]]
	}
	return score
}

// (1 for Rock, 2 for Paper, and 3 for Scissors)
// X means you need to lose, Y means you need to draw, and Z means you need to win
func Puzzle02(filename string) int {
	shapeScoreMap := map[string]int{
		"A X": 3,
		"A Y": 1,
		"A Z": 2,
		"B X": 1,
		"B Y": 2,
		"B Z": 3,
		"C X": 2,
		"C Y": 3,
		"C Z": 1,
	}

	gameScoreMap := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	score := 0

	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		score = score + shapeScoreMap[scanner.Text()]

		tmp := strings.Split(scanner.Text(), " ")
		score = score + gameScoreMap[tmp[1]]
	}
	return score
}
