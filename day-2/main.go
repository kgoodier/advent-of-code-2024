package main

import (
	"aoc24/elves"
	"fmt"
	"slices"
)

func main() {
	data, err := elves.ReadRows("input.txt", elves.ToNumbers)
	if err != nil {
		panic(err)
	}

	p1(data)
	p2(data)
}

func p1(data [][]int) {
	safeReports := 0
	for _, report := range data {
		ok, _ := checkReport(report)
		if ok {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func p2(data [][]int) {
	safeReports := 0
	for _, report := range data {
		ok, _ := checkReport(report)
		if ok {
			safeReports++
			continue
		}

		if checkWithDeletion(report) {
			safeReports++
			continue
		}
	}

	fmt.Println(safeReports)
}

func checkWithDeletion(report []int) bool {
	for i := 0; i < len(report); i++ {
		reportLeft := deleteElement(report, i)
		ok, _ := checkReport(reportLeft)
		if ok {
			return true
		}
	}
	return false
}

func deleteElement(s []int, i int) []int {
	c := slices.Clone(s)
	return slices.Delete(c, i, i+1)
}

func checkReport(report []int) (bool, int) {
	prev := 0
	for i := 0; i < len(report)-1; i++ {
		d := report[i+1] - report[i]
		if d < 0 && prev <= 0 && d >= -3 {
			prev = d
			continue
		} else if d > 0 && prev >= 0 && d <= 3 {
			prev = d
			continue
		} else {
			return false, i
		}
	}

	return true, 0
}
