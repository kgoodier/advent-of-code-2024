package main

import (
	"aoc24/elves"
	"fmt"
)

type Puzzle struct {
	lines  []string
	width  int
	height int
}

func main() {
	lines, err := elves.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	p := &Puzzle{
		lines:  lines,
		width:  len(lines[0]),
		height: len(lines),
	}

	part1(p)
	part2(p)
}

func part1(p *Puzzle) {
	total := 0
	word := "XMAS"
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			total += p.search(word, x, y, 1, 0)   // right
			total += p.search(word, x, y, -1, 0)  // left
			total += p.search(word, x, y, 0, 1)   // down
			total += p.search(word, x, y, 0, -1)  // up
			total += p.search(word, x, y, 1, -1)  // right up
			total += p.search(word, x, y, 1, 1)   // right down
			total += p.search(word, x, y, -1, -1) // left up
			total += p.search(word, x, y, -1, 1)  // left down
		}
	}

	fmt.Println("Part 1: ", total)
}

func part2(p *Puzzle) {
	// Looking for 2 diagonal "MAS" with the A overlapping
	total := 0
	word := "MAS"
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			c1 := 0
			c2 := 0

			// We're working on a 3x3 grid, so we have to nudge the starting point
			// of the searches out to the corners

			// / diagonal
			c1 += p.search(word, x-1, y+1, 1, -1) // right up
			c1 += p.search(word, x+1, y-1, -1, 1) // left down

			if c1 == 1 {
				// \ diagonal
				c2 += p.search(word, x-1, y-1, 1, 1)   // right down
				c2 += p.search(word, x+1, y+1, -1, -1) // left up
			}

			if c1 == 1 && c2 == 1 {
				total++
			}
		}
	}

	fmt.Println("Part 2: ", total)

}

func (p *Puzzle) get(x, y int) byte {
	if x < 0 || x >= p.width || y < 0 || y >= p.height {
		return 0
	}
	return p.lines[y][x]
}

func (p *Puzzle) search(word string, x, y, incX, incY int) int {
	for i := 0; i < len(word); i++ {
		if p.get(x+i*incX, y+i*incY) != word[i] {
			return 0
		}
	}
	return 1
}
