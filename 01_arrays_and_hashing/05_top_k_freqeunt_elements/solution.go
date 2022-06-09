package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	hashset := make(map[int]int)

	for _, num := range nums {
		hashset[num]++
	}

	reverseDic := map[int][]int{}
	for k, v := range hashset {
		reverseDic[v] = append(reverseDic[v], k)
	}

	result := []int{}
	for i := len(nums); len(result) < k; i-- {
		result = append(result, reverseDic[i]...)
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Println(result)
}
