package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	data := strings.TrimSpace(string(bytes))

	part1(data)
	part2(data)
}

func part1(data string) {
	// Regular expression to match valid mul(X,Y) instructions
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Find all matches
	matches := re.FindAllStringSubmatch(data, -1)

	sum := 0
	for _, match := range matches {
		// Extract the numbers
		x, err1 := strconv.Atoi(match[1])
		y, err2 := strconv.Atoi(match[2])
		if err1 != nil || err2 != nil {
			continue
		}

		// Perform the multiplication and add to the sum
		sum += x * y
	}

	fmt.Println("[Part 1] Sum of all multiplications:", sum)
}

func part2(data string) {
	// Regular expression to match valid mul(X,Y) instructions
	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)

	// Find all matches
	matches := re.FindAllStringSubmatch(data, -1)

	sum := 0
	multEnabled := true

	for _, match := range matches {
		// Extract the name and numbers
		cmd := match[1]

		switch {
		case cmd == "do()":
			multEnabled = true
			fmt.Println("Enabled multiplication")
		case cmd == "don't()":
			multEnabled = false
			fmt.Println("Disabled multiplication")
		case strings.HasPrefix(cmd, "mul("):
			if multEnabled {
				x, err1 := strconv.Atoi(match[2])
				y, err2 := strconv.Atoi(match[3])
				if err1 != nil || err2 != nil {
					continue
				}
				sum += x * y
			}
		default:
			panic("Invalid command " + cmd)
		}
	}

	fmt.Println("[Part 2] Sum of all multiplications:", sum)
}
