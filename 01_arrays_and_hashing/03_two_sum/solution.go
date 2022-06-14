package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for index, num := range nums {
		numTarget := target - num

		if seenIdx, isPresent := hash[numTarget]; isPresent {
			return []int{seenIdx, index}
		}

		hash[num] = index
	}

	return []int{0, 0}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}
