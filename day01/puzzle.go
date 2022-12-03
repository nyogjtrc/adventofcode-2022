package day01

import (
	"log"
	"strconv"

	"github.com/nyogjtrc/adventofcode-2022/input"
)

func Puzzle01(filename string) int {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	caloriesSlice := []int{}
	tmp := 0

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			caloriesSlice = append(caloriesSlice, tmp)
			tmp = 0
		}
		cal, _ := strconv.Atoi(scanner.Text())
		tmp = tmp + cal
	}

	if tmp != 0 {
		caloriesSlice = append(caloriesSlice, tmp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	theMost := 0

	for k := range caloriesSlice {
		if caloriesSlice[k] > theMost {
			theMost = caloriesSlice[k]
		}
	}
	return theMost
}

func Puzzle02(filename string) int {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	caloriesSlice := []int{}
	tmp := 0

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			caloriesSlice = append(caloriesSlice, tmp)
			tmp = 0
		}
		cal, _ := strconv.Atoi(scanner.Text())
		tmp = tmp + cal
	}

	if tmp != 0 {
		caloriesSlice = append(caloriesSlice, tmp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		for j := i; j < len(caloriesSlice); j++ {
			if caloriesSlice[i] < caloriesSlice[j] {
				caloriesSlice[i], caloriesSlice[j] = caloriesSlice[j], caloriesSlice[i]
			}

		}
	}

	return caloriesSlice[0] + caloriesSlice[1] + caloriesSlice[2]
}
