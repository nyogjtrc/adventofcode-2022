package day06

import (
	"log"

	"github.com/nyogjtrc/adventofcode-2022/input"
)

func Puzzle01(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Scan() {
		answer = differentMarker(scanner.Text(), 4)
	}
	return
}

func differentMarker(s string, distinct int) (marker int) {
	buf := ""

	for i := 0; i < len(s); i++ {
		for k := range buf {
			if s[i] == buf[k] {
				buf = buf[k+1:]
				break
			}
		}
		buf = buf + string(s[i])

		if len(buf) == distinct {
			return i + 1
		}
	}

	return
}

func Puzzle02(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Scan() {
		answer = differentMarker(scanner.Text(), 14)
	}
	return
}
