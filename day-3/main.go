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

	fmt.Println("Sum of all multiplications:", sum)
}
