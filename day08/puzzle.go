package day08

import (
	"log"
	"math"
	"strconv"

	"github.com/nyogjtrc/adventofcode-2022/input"
)

func Puzzle01(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	grid := []int{}
	size := 0

	for scanner.Scan() {
		t := scanner.Text()
		size = len(scanner.Text())
		for k := range t {
			h, _ := strconv.Atoi(string(t[k]))

			grid = append(grid, h)
		}
	}
	//fmt.Println(size, grid)
	for k := range grid {
		if isVisible(size, grid, k) {
			answer++
		}
	}
	return
}

func isVisible(size int, grid []int, position int) bool {
	row := position / size
	col := int(math.Mod(float64(position), float64(size)))

	//fmt.Println("position", row, col)

	if row == 0 {
		return true
	}

	if col == 0 {
		return true
	}

	if row == size-1 {
		return true
	}

	if col == size-1 {
		return true
	}

	// row
	leftVisible := true
	rightVisible := true
	for i := 0; i < size; i++ {
		index := i + size*row
		//fmt.Printf("%d ", index)
		// skip self
		if position == index {
			continue
		}

		if grid[position] <= grid[index] {
			if index < position {
				leftVisible = false
			} else {
				rightVisible = false
			}
		}
	}

	// col
	upVisible := true
	downVisible := true
	for i := 0; i < size; i++ {
		index := col + size*i
		//fmt.Printf("%d ", index)
		// skip self
		if position == index {
			continue
		}

		if grid[position] <= grid[index] {
			if index < position {
				upVisible = false
			} else {
				downVisible = false
			}
		}
	}
	//fmt.Println(leftVisible, rightVisible, upVisible, downVisible)

	return leftVisible || rightVisible || upVisible || downVisible
}

func Puzzle02(filename string) (answer int) {
	scanner, err := input.Load(filename)
	if err != nil {
		log.Fatal(err)
	}

	grid := []int{}
	size := 0

	for scanner.Scan() {
		t := scanner.Text()
		size = len(scanner.Text())
		for k := range t {
			h, _ := strconv.Atoi(string(t[k]))

			grid = append(grid, h)
		}
	}
	//fmt.Println(size, grid)
	for k := range grid {

		score := scenicScore(grid, size, k)

		if score > answer {
			answer = score
		}
	}
	return
}

func scenicScore(grid []int, size, position int) int {
	row := position / size
	col := int(math.Mod(float64(position), float64(size)))

	//fmt.Println(row, col,
	//"s",
	//upScore(grid, size, position, row, col),
	//downScore(grid, size, position, row, col),
	//leftScore(grid, size, position, row, col),
	//rightScore(grid, size, position, row, col),
	//)

	return upScore(grid, size, position, row, col) *
		downScore(grid, size, position, row, col) *
		leftScore(grid, size, position, row, col) *
		rightScore(grid, size, position, row, col)
}

func toPosition(row, col, size int) int {
	return row*size + col
}

func upScore(grid []int, size, position, row, col int) int {
	if row == 0 {
		return 0
	}

	for i := 1; ; i++ {
		// touch edge
		if row-i == 0 {
			return i
		}

		if grid[toPosition(row-i, col, size)] >= grid[position] {
			return i
		}
	}
}

func downScore(grid []int, size, position, row, col int) int {
	if row == size-1 {
		return 0
	}

	for i := 1; ; i++ {
		// touch edge
		if row+i == size-1 {
			return i
		}

		if grid[toPosition(row+i, col, size)] >= grid[position] {
			return i
		}
	}
}

func leftScore(grid []int, size, position, row, col int) int {
	if col == 0 {
		return 0
	}

	for i := 1; ; i++ {
		// touch edge
		if col-i == 0 {
			return i
		}

		if grid[toPosition(row, col-i, size)] >= grid[position] {
			return i
		}
	}
}
func rightScore(grid []int, size, position, row, col int) int {
	if col == size-1 {
		return 0
	}

	for i := 1; ; i++ {
		// touch edge
		if col+i == size-1 {
			return i
		}

		if grid[toPosition(row, col+i, size)] >= grid[position] {
			return i
		}
	}
}
