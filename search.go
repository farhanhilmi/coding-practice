package main

// binarySearch performs a binary search on a sorted array.
func binarySearch(arr []int, target int) bool {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return true // Element found, return its index
		} else if arr[mid] < target {
			low = mid + 1 // Search the right half
		} else {
			high = mid - 1 // Search the left half
		}
	}

	return false // Element not found
}

func linearSearch(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}

	return false
}

// func main() {
// 	arr := []int{1, 2, 5, 7, 8, 9, 10, 17, 20}
// 	// res := binarySearch(arr, 3)
// 	res := linearSearch(arr, 20)

// 	fmt.Println(res)
// }
