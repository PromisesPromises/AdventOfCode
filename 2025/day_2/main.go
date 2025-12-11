package main

import (
	"fmt"
	"strconv"
	"strings"

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

func main() {
	godotenv.Load(".env")
	// input := util.GetInput("https://adventofcode.com/2025/day/2/input", os.Getenv("session_cookie"))
	fmt.Println(part1("4440-4500,5555-5555"))
}
