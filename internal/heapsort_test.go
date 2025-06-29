package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_heapSort(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected []int
		errType  error
	}{
		{
			name:     "success/ random numbers",
			numbers:  []int{2, 5, 7, 8, 3, 7, 9, 1, 4},
			expected: []int{1, 2, 3, 4, 5, 7, 7, 8, 9},
			errType:  nil,
		},
		{
			name:     "success/ already sorted numbers",
			numbers:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			errType:  nil,
		},
		{
			name:     "success/ reverse sorted numbers",
			numbers:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			errType:  nil,
		},
		{
			name:     "success/ two elements",
			numbers:  []int{5, 2},
			expected: []int{2, 5},
			errType:  nil,
		},
		{
			name:     "success/ duplicate elements",
			numbers:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
			errType:  nil,
		},
		{
			name:     "success/ negative numbers",
			numbers:  []int{-3, -1, -4, -1, -5, -9, -2, -6},
			expected: []int{-9, -6, -5, -4, -3, -2, -1, -1},
			errType:  nil,
		},
		{
			name:     "success/ mixed positive and negative",
			numbers:  []int{3, -1, 4, -1, 5, -9, 2, 6},
			expected: []int{-9, -1, -1, 2, 3, 4, 5, 6},
			errType:  nil,
		},
		{
			name:     "error/ empty array",
			numbers:  []int{},
			expected: []int{},
			errType:  fmt.Errorf("invalid heap parameters: n %d, i %d, len(numbers) %d", 0, 0, 0),
		},
		{
			name:     "error/ single element array",
			numbers:  []int{42},
			expected: []int{42},
			errType:  fmt.Errorf("invalid input: numbers array must contain at least two elements"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.numbers))
			copy(input, tt.numbers)

			err := HeapSort(input)

			if tt.errType == nil {
				require.NoError(t, err)
				require.Equal(t, tt.expected, input)
			} else {
				require.EqualError(t, err, tt.errType.Error())
				require.Equal(t, tt.numbers, input)
			}
		})
	}
}

func Test_heapify_behavior(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		i        int
		n        int
		expected []int
		errType  error
	}{
		{
			name:     "success/ heapify root of unsorted array",
			input:    []int{2, 5, 7, 8, 3, 7, 9, 1, 4},
			i:        0,
			n:        9,
			expected: []int{7, 5, 9, 8, 3, 7, 2, 1, 4},
		},
		{
			name:     "success/ heapify with already max-heap at root",
			input:    []int{9, 5, 7, 8, 3, 7, 2, 1, 4},
			i:        0,
			n:        9,
			expected: []int{9, 5, 7, 8, 3, 7, 2, 1, 4},
		},
		{
			name:     "success/ heapify subtree",
			input:    []int{9, 5, 7, 1, 3, 7, 2, 8, 4},
			i:        3,
			n:        9,
			expected: []int{9, 5, 7, 8, 3, 7, 2, 1, 4},
		},
		{
			name:     "success/ heapify leaf node (no change)",
			input:    []int{9, 5, 7, 8, 3, 7, 2, 1, 4},
			i:        7,
			n:        9,
			expected: []int{9, 5, 7, 8, 3, 7, 2, 1, 4},
		},
		{
			name:     "success/ heapify single element",
			input:    []int{42},
			i:        0,
			n:        1,
			expected: []int{42},
		},
		{
			name:    "error/ negative index",
			input:   []int{1, 2, 3},
			i:       -1,
			n:       3,
			errType: fmt.Errorf("invalid heap parameters: i %d, n %d", -1, 3),
		},
		{
			name:    "error/ index out of bounds",
			input:   []int{1, 2, 3},
			i:       3,
			n:       3,
			errType: fmt.Errorf("invalid heap parameters: i %d, n %d", 3, 3),
		},
		{
			name:    "error/ negative heap size",
			input:   []int{1, 2, 3},
			i:       0,
			n:       -1,
			errType: fmt.Errorf("invalid heap parameters: i %d, n %d", 0, -1),
		},
		{
			name:    "error/ zero heap size",
			input:   []int{1, 2, 3},
			i:       0,
			n:       0,
			errType: fmt.Errorf("invalid heap parameters: i %d, n %d", 0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy to avoid modifying the original test data
			output := make([]int, len(tt.input))
			copy(output, tt.input)

			err := Heapify(output, tt.i, tt.n)

			if tt.errType == nil {
				require.NoError(t, err)
				require.Equal(t, tt.expected, output)
			} else {
				require.EqualError(t, err, tt.errType.Error())
				// For error cases, the array should remain unchanged
				require.Equal(t, tt.input, output)
			}
		})
	}
}