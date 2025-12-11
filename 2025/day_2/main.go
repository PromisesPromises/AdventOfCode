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
	sum := 0
	for _, line := range strings.Split(input, ",") {
		ranges := strings.Split(line, "-")
		start, _ := strconv.Atoi(ranges[0])
		end, _ := strconv.Atoi(ranges[1])

		for i := start; i <= end; i++ {
			j := strconv.Itoa(i)
			if len(j)%2 == 1 {
				continue
			}

			if j[:len(j)/2] == j[len(j)/2:] {
				sum += i
			}
		}
	}
	return sum
}

func part2(input string) int64 {
	var sum int64 = 0
	for _, line := range strings.Split(input, ",") {
		ranges := strings.Split(line, "-")
		start, _ := strconv.Atoi(ranges[0])
		end, _ := strconv.Atoi(ranges[1])

		for i := start; i <= end; i++ {
			i_str := strconv.Itoa(i)

			if len(i_str) == 1 {
				continue
			}

			for k := 0; k <= len(i_str); k++ {
				count := strings.Count(i_str, i_str[0:k])
				if k*count == len(i_str) && count > 1 {
					sum += int64(i)
					fmt.Println(i_str)
					break
				}
			}

		}
	}
	return sum
}

func main() {
	godotenv.Load(".env")
	input := util.GetInput("https://adventofcode.com/2025/day/2/input", os.Getenv("session_cookie"))
	fmt.Println(part2(input))
	// fmt.Println(part2("824824821-824824827"))
	// fmt.Println(part2("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"))
}
