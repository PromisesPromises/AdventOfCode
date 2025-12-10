package main

import (
	"fmt"
	"os"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) int {
	return 0
}

func main() {
	input := util.GetInput("https://adventofcode.com/2025/day/1/input", os.Getenv("session_cookie"))
	fmt.Println(part1(input))
}
