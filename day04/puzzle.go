package day04

import (
	"log"
	"strconv"
	"strings"

	"github.com/nyogjtrc/adventofcode-2022/input"
)

func Puzzle01(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), ",")
		s1 := convertToSections(splits[0])
		s2 := convertToSections(splits[1])

		if contains(s1, s2) || contains(s2, s1) {
			answer = answer + 1
		}
	}
	return
}

func convertToSections(s string) []int {
	splits := strings.Split(s, "-")
	begin, _ := strconv.Atoi(splits[0])
	end, _ := strconv.Atoi(splits[1])

	return []int{begin, end}
}

func contains(s1, s2 []int) bool {
	return s1[0] <= s2[0] && s1[1] >= s2[1]
}

func Puzzle02(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), ",")
		s1 := convertToSections(splits[0])
		s2 := convertToSections(splits[1])

		if overlaps(s1, s2) {
			answer++
		}
	}
	return
}

func overlaps(s1, s2 []int) bool {
	sMap := make(map[int]bool)

	for i := s1[0]; i <= s1[1]; i++ {
		sMap[i] = true
	}

	for j := s2[0]; j <= s2[1]; j++ {
		if sMap[j] {
			return true
		}
	}
	return false
}
