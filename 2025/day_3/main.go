package main

import (
	"fmt"
	"os"

	"github.com/GarrettWells/AdventOfCode/util"
	"github.com/joho/godotenv"
)

func part1(input string) int {
	return 0
}

func main() {
	godotenv.Load(".env")
	input := util.GetInput("https://adventofcode.com/2025/day/1/input", os.Getenv("session_cookie"))
	fmt.Println(part1(input))
}
