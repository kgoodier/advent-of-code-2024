package main

import (
	"aoc24/elves"
	"fmt"
)

func main() {
	data, err := elves.ReadRows("input.txt", elves.ToNumbers)
	if err != nil {
		panic(err)
	}

	safeReports := 0
	for _, report := range data {
		prev := 0
		safe := true
		for i := 0; i < len(report)-1; i++ {
			d := report[i+1] - report[i]
			if d < 0 && prev <= 0 && d >= -3 {
				prev = d
				continue
			} else if d > 0 && prev >= 0 && d <= 3 {
				prev = d
				continue
			} else {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
