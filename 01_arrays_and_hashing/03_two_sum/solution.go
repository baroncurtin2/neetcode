package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, num := range nums {
		numTarget := target - num

		if idx, isPresent := hash[numTarget]; isPresent {
			return []int{idx, i}
		}

		hash[numTarget] = i
	}

	fmt.Println(hash)
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}
