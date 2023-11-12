package main

func findSingleNumberBitwise(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
		// fmt.Println("res", result)
	}
	return result
}

func findSingleNumber(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	for num, count := range counts {
		if count == 1 {
			return num
		}
	}

	// Return -1 or any other value to indicate no single number found
	return -1
}

func findNonDuplicates(arr []int) []int {
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}

	var result []int
	for num, count := range counts {
		if count == 1 {
			result = append(result, num)
		}
	}

	return result
}

// func main() {
// 	array := []int{2, 4, 6, 2, 4, 7, 10}
// 	nonDuplicates := findSingleNumber(array)
// 	fmt.Printf("Non-duplicated numbers: %v\n", nonDuplicates)
// }
