package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
	"github.com/joho/godotenv"
)

func split_input(input string) (string, string) {
	split := strings.Split(input, "\n\n")
	return split[0], split[1]
}

func part1(input string) int {
	ranges, ids := split_input(input)

	fresh_ingredients_list := FreshIngredientsList{}

	for _, _range := range strings.Split(ranges, "\n") {
		_range_split := strings.Split(_range, "-")
		lower, _ := strconv.Atoi(_range_split[0])
		upper, _ := strconv.Atoi(_range_split[1])
		fresh_ingredients_list.arr = append(fresh_ingredients_list.arr, FreshIngredientsRange{lower: lower, upper: upper})
	}

	count := 0
	for _, id := range strings.Split(ids, "\n") {
		id_int, _ := strconv.Atoi(id)
		if fresh_ingredients_list.isFresh(id_int) {
			count++
		}
	}

	return count
}

func part2(input string) uint64 {
	ranges, _ := split_input(input)

	fresh_ingredients_list := FreshIngredientsList{}

	for _, _range := range strings.Split(ranges, "\n") {
		_range_split := strings.Split(_range, "-")
		lower, _ := strconv.Atoi(_range_split[0])
		upper, _ := strconv.Atoi(_range_split[1])
		_range := FreshIngredientsRange{lower: lower, upper: upper}
		fresh_ingredients_list.compacted_add(_range)
	}

	var count uint64 = 0

	for _, _range := range fresh_ingredients_list.arr {
		count += uint64(_range.upper - _range.lower + 1)
	}

	return count
}

func main() {
	godotenv.Load(".env")
	input := util.GetInput("https://adventofcode.com/2025/day/5/input", os.Getenv("session_cookie"))
	fmt.Println(part2(input))
	// fmt.Println(part2("100-200\n50-250\n1-2\n50-150\n75-175\n125-175\n175-225\n300-400\n\n1\n5\n8\n11\n17\n32"))
}
