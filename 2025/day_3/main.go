package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
	"github.com/joho/godotenv"
)

func part1(input string) int {
	joltes := 0
	for _, line := range strings.Split(input, "\n") {
		index_1 := get_largest_digit(line[0 : len(line)-1])
		offset := index_1 + 1
		index_2 := get_largest_digit(line[offset:]) + offset
		line_joltes, _ := strconv.Atoi(fmt.Sprintf("%c%c", line[index_1], line[index_2]))
		joltes += line_joltes
	}
	return joltes
}

func part2(input string) int {
	total_joltes := 0
	for _, line := range strings.Split(input, "\n") {
		joltes_str := ""
		curr_index := -1
		for i := 11; i >= 0; i-- {
			index := get_largest_digit(line[curr_index+1 : len(line)-i])
			curr_index = curr_index + index + 1
			joltes_str = joltes_str + string(line[curr_index])
		}

		joltes, _ := strconv.Atoi(joltes_str)
		total_joltes += joltes
	}
	return total_joltes
}

// returns the index of the largest number
func get_largest_digit(input string) int {
	for i := 9; i >= 0; i-- {
		index := strings.Index(input, strconv.Itoa(i))
		if index != -1 {
			return index
		}
	}
	return -1
}

func main() {
	godotenv.Load(".env")
	input := util.GetInput("https://adventofcode.com/2025/day/3/input", os.Getenv("session_cookie"))
	fmt.Println(part2(input))
	// fmt.Println(part1("111116\n1181114"))
}
