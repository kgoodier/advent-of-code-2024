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

	p := Puzzle{
		lines:  lines,
		width:  len(lines[0]),
		height: len(lines),
	}

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
