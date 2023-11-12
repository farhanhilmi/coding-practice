package main

import (
	"math/rand"
	"time"
)

// bubbleSort performs the Bubble Sort algorithm on an integer slice.
func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements and swap them if they are in the wrong order
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// func main() {
// 	arr := []int{99, 20, 50, 1, 21, 100, 40, 77, 53, 89, 7}
// 	// res := bubleSort(arr)
// 	res := mergeSort(arr)

// 	log.Println(res)
// }

// merge is a helper function that merges two sorted slices into a single sorted slice.
func merge(left []int, right []int) []int {
	// result := make([]int, 0, len(left)+len(right))
	result := []int{}
	i, j := 0, 0

	// Compare elements from both slices and merge them into the result slice
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append the remaining elements from both slices (if any)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// mergeSort is the recursive function that implements the Merge Sort algorithm.
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the slice into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// Randomized pivot
		rand.Seed(time.Now().UnixNano())
		pivotIndex := low + rand.Intn(high-low+1)
		arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

		// Partition the array, arr[p] is now at the correct position
		p := partition(arr, low, high)

		// Recursively sort the elements before and after partition
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose the rightmost element as the pivot
	pivot := arr[high]

	// Index of smaller element
	i := low - 1

	// Traverse through the array and rearrange elements
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Place the pivot element at its correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// Return the index where the pivot element is placed
	return i + 1
}
