package main

import (
	"aoc24/elves"
	"fmt"
)

func main() {
	lines, err := elves.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	x, y := findStartingPoint(lines)
	if x == -1 {
		panic("starting point not found")
	}

	// Direction of guard (as deltas)
	dir := "up"
	dx := 0
	dy := -1

	guardPositions := 1
	lines[y] = lines[y][:x] + "X" + lines[y][x+1:]
	for x >= 0 && x < len(lines[0]) && y >= 0 && y < len(lines) {
		// really the "next" tile
		tile := lines[y][x]
		if tile == '#' {
			// turn 90 degrees to the right and backtrack
			switch dir {
			case "up":
				dx = 1
				dy = 0
				dir = "right"
				y += 1
			case "right":
				dx = 0
				dy = 1
				dir = "down"
				x -= 1
			case "down":
				dx = -1
				dy = 0
				dir = "left"
				y -= 1
			case "left":
				dx = 0
				dy = -1
				dir = "up"
				x += 1
			}
			continue
		} else if tile != 'X' {
			guardPositions++
			lines[y] = lines[y][:x] + "X" + lines[y][x+1:]
		}
		x += dx
		y += dy
	}
	fmt.Printf("Day 1 answer: %d\n", guardPositions)
}

func findStartingPoint(lines []string) (int, int) {
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == '^' {
				return x, y
			}
		}
	}
	return -1, -1
}
