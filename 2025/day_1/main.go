package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/GarrettWells/AdventOfCode/util"
)

func part1(input string) int {
	password := 0
	value := 50
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		turn, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(line, "L") {
			turn *= -1
		}
		value = (value + turn) % 100
		if value == 0 {
			password++
		}
	}
	return password
}

func part2(input string) int {
	password := 0
	value := 50
	for _, line := range strings.Split(input, "\n") {
		log.Print("Current value:", value)
		log.Println("Processing line:", line)
		if line == "" {
			continue
		}
		turn, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(line, "L") {
			turn *= -1
		}
		old_value := value
		value += turn

		if value <= 0 {
			password += int(math.Abs(float64(value)))/100 + 1
			if old_value == 0 {
				password--
			}
		} else if value >= 100 {
			password += value / 100
		}

		value %= 100
		value += 100
		value %= 100
	}
	return password
}

func main() {
	input := util.GetInput("https://adventofcode.com/2025/day/1/input", os.Getenv("session_cookie"))
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
