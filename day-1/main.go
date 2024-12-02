package main

import (
	"errors"
	"fmt"
	"slices"

	"aoc24/elves"
)

func main() {
	data, err := elves.ReadColumns("input1.txt", elves.ToNumbers)
	if err != nil {
		panic(err)
	}
	if len(data) != 2 || len(data[0]) != len(data[1]) {
		panic(errors.New("invalid input"))
	}

	slices.Sort(data[0])
	slices.Sort(data[1])

	d := 0
	for i := 0; i < len(data[0]); i++ {
		if data[0][i] < data[1][i] {
			d += data[1][i] - data[0][i]
		} else {
			d += data[0][i] - data[1][i]
		}
	}

	fmt.Println(d)
}
