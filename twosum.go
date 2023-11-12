package main

func twoSum(arr []int, total int) []int {
	n := len(arr)
	for i := range arr {
		for j := 0; j < n-i-1; j++ {
			if arr[j]+arr[j+1] == total {
				return []int{arr[j], arr[j+1]}
			}
		}
	}

	return []int{}
}

func findPairWithSum(arr []int, targetSum int) (int, int) {
	seen := make(map[int]bool)

	for _, num := range arr {
		complement := targetSum - num

		if seen[complement] {
			return complement, num
		}

		seen[num] = true
	}

	return -1, -1
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	// res := twoSum(arr, 4)
// 	res1, res2 := findPairWithSum(arr, 4)

// 	fmt.Println(res1, res2)
// }
