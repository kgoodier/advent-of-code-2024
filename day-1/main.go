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

	list1 := data[0]
	list2 := data[1]

	slices.Sort(list1)
	slices.Sort(list2)

	// Part 1
	d := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] < list2[i] {
			d += list2[i] - list1[i]
		} else {
			d += list1[i] - list2[i]
		}
	}
	fmt.Println(d)

	// Part 2
	counts := map[int]int{}
	for _, v := range list2 {
		counts[v]++
	}
	s := 0
	for _, v := range list1 {
		s += v * counts[v]
	}
	fmt.Println(s)
}
