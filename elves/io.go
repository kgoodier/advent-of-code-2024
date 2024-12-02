package elves

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

type Converter[T any] func(row int, col int, val string) (T, error)

func ToNumbers(row int, col int, val string) (int, error) {
	return strconv.Atoi(val)
}

func ReadColumns[T any](filename string, fn Converter[T]) ([][]T, error) {
	rows, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	columns := [][]T{}
	for r := range rows {
		cols := strings.Fields(rows[r])

		if len(columns) == 0 {
			columns = make([][]T, len(cols))
		}

		for c := range cols {
			val, err := fn(r, c, cols[c])
			if err != nil {
				return nil, err
			}
			columns[c] = append(columns[c], val)
		}
	}

	return columns, nil
}

func ReadRows[T any](filename string, fn Converter[T]) ([][]T, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	rows := [][]T{}
	for r := range lines {
		cols := strings.Fields(lines[r])
		row := make([]T, len(cols))

		for c := range cols {
			val, err := fn(r, c, cols[c])
			if err != nil {
				return nil, err
			}
			row[c] = val
		}
		rows = append(rows, row)
	}

	return rows, nil
}
