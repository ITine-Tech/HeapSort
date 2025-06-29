package internal

import (
	"fmt"
)

// HeapSort sorts a slice of integers in ascending order using the Heap Sort algorithm.
// It builds a max heap from the input slice, then repeatedly extracts the maximum
// element and places it at the end of the slice, reducing the heap size each time.
// The sorting is performed in-place.
//
// Parameters:
//   - numbers: A slice of integers to be sorted.
//
// Returns:
//   - error: An error if the input slice has fewer than two elements.
//
// Note:
//   - This function modifies the input slice directly.
func HeapSort(numbers []int) error {
	if len(numbers) == 0 {
		return fmt.Errorf("invalid heap parameters: n %d, i %d, len(numbers) %d", 0, 0, len(numbers))
	}
	if len(numbers) <= 1 {
		return fmt.Errorf("invalid input: numbers array must contain at least two elements")
	}

	n := len(numbers)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		if err := Heapify(numbers, i, n); err != nil {
			return fmt.Errorf("heapify failed during heap build at i %d: %w", i, err)
		}
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		numbers[0], numbers[i] = numbers[i], numbers[0]
		if err := Heapify(numbers, 0, i); err != nil {
			return fmt.Errorf("heapify failed during sort at i %d: %w", i, err)
		}
	}
	return nil
}

// Heapify maintains the max-heap property for a subtree rooted at index i in a slice of integers.
func Heapify(numbers []int, i, n int) error {
	if i < 0 || i >= n || n <= 0 {
		return fmt.Errorf("invalid heap parameters: i %d, n %d", i, n)
	}

	for {
		max := i
		left := 2*i + 1
		right := 2*i + 2

		if left < n && numbers[left] > numbers[max] {
			max = left
		}
		if right < n && numbers[right] > numbers[max] {
			max = right
		}
		if max == i {
			break
		}
		numbers[i], numbers[max] = numbers[max], numbers[i]
		i = max
	}
	return nil
}
