package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
	"github.com/joho/godotenv"
)

func part1(input string) int {
	rows, cols := grid_size(input)

	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	for i, row := range strings.Split(input, "\n") {
		for j, val := range strings.Split(row, "") {
			if val == "@" {
				increment_grid_squares(i, j, &grid, 1)
			} else {
				grid[i][j] = 999
			}
		}
	}

	count := 0
	for _, row := range grid {
		for _, val := range row {
			if val < 4 {
				count++
			}
		}
	}
	return count
}

func part2(input string) int {
	rows, cols := grid_size(input)

	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	for i, row := range strings.Split(input, "\n") {
		for j, val := range strings.Split(row, "") {
			if val == "@" {
				increment_grid_squares(i, j, &grid, 1)
			} else {
				grid[i][j] = 999
			}
		}
	}

	count := 0
	for {
		finished := true
		for i, row := range grid {
			for j, val := range row {
				if val < 4 && val >= 0 {
					count++
					increment_grid_squares(i, j, &grid, -1)
					grid[i][j] = -1
					finished = false
				}
			}
		}
		if finished {
			break
		}
	}

	return count
}

func increment_grid_squares(row int, col int, grid *[][]int, increment_value int) {
	left := col > 0
	right := col < len((*grid)[0])-1
	top := row > 0
	bottom := row < len(*grid)-1

	_grid := (*grid)
	if left {
		_grid[row][col-1] += increment_value
		if top {
			_grid[row-1][col-1] += increment_value
		}
		if bottom {
			_grid[row+1][col-1] += increment_value
		}
	}
	if right {
		_grid[row][col+1] += increment_value
		if top {
			_grid[row-1][col+1] += increment_value
		}
		if bottom {
			_grid[row+1][col+1] += increment_value
		}
	}
	if top {
		_grid[row-1][col] += increment_value
	}
	if bottom {
		_grid[row+1][col] += increment_value
	}

}

func grid_size(input string) (int, int) {
	x := strings.Count(input, "\n") + 1
	y := strings.Index(input, "\n")
	return x, y
}

func main() {
	godotenv.Load(".env")
	input := util.GetInput("https://adventofcode.com/2025/day/4/input", os.Getenv("session_cookie"))
	fmt.Println(part2(input))
	// fmt.Println(part2("..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@."))
}
