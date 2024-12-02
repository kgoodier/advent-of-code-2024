package elves

import (
	"errors"
	"os"
	"strconv"
	"testing"
)

func TestReadColumns(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	content := "1 2 3\n4 5 6\n7  8   9"
	if _, err := file.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	// Define a converter function
	converter := func(row int, col int, val string) (int, error) {
		switch val {
		case "error":
			return 0, errors.New("conversion error")
		default:
			return strconv.Atoi(val)
		}
	}

	tests := []struct {
		filename string
		fn       Converter[int]
		expected [][]int
		hasError bool
	}{
		{
			filename: file.Name(),
			fn:       converter,
			expected: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
			hasError: false,
		},
		{
			filename: "nonexistentfile",
			fn:       converter,
			expected: nil,
			hasError: true,
		},
		{
			filename: file.Name(),
			fn: func(row int, col int, val string) (int, error) {
				return 0, errors.New("conversion error")
			},
			expected: nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		result, err := ReadColumns(test.filename, test.fn)
		if (err != nil) != test.hasError {
			t.Errorf("ReadColumns() error = %v, wantErr %v", err, test.hasError)
			continue
		}
		if !test.hasError && !equal(result, test.expected) {
			t.Errorf("ReadColumns() = %v, want %v", result, test.expected)
		}
	}
}
func TestReadRows(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	content := "1 2 3\n4 5 6\n7  8   9"
	if _, err := file.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	// Define a converter function
	converter := func(row int, col int, val string) (int, error) {
		switch val {
		case "error":
			return 0, errors.New("conversion error")
		default:
			return strconv.Atoi(val)
		}
	}

	tests := []struct {
		filename string
		fn       Converter[int]
		expected [][]int
		hasError bool
	}{
		{
			filename: file.Name(),
			fn:       converter,
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			hasError: false,
		},
		{
			filename: "nonexistentfile",
			fn:       converter,
			expected: nil,
			hasError: true,
		},
		{
			filename: file.Name(),
			fn: func(row int, col int, val string) (int, error) {
				return 0, errors.New("conversion error")
			},
			expected: nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		result, err := ReadRows(test.filename, test.fn)
		if (err != nil) != test.hasError {
			t.Errorf("ReadRows() error = %v, wantErr %v", err, test.hasError)
			continue
		}
		if !test.hasError && !equal(result, test.expected) {
			t.Errorf("ReadRows() = %v, want %v", result, test.expected)
		}
	}
}
func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
