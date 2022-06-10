package main

import (
	"fmt"
	"math"
)

func longestConsecutive(nums []int) int {
	// create a map to store the nums
	numMap := make(map[int]bool)

	// iterate through the nums and add them to the map
	numMap = addNumsToMap(numMap, nums)

	// create a variable to store the longest consecutive sequence
	longest := 0

	// iterate through the map
	// if num - 1 is not in the map,
	// then we know that num is the start of a new sequence
	for num, _ := range numMap {

		// if num - 1 is in the set,
		// then we know that num is *not* the start of a new sequence
		if numMap[num-1] {
			continue
		}

		count := 0
		for numMap[num+count] {
			count++
		}

		longest = int(math.Max(float64(longest), float64(count)))
	}
	return longest
}

func addNumToMap(numMap map[int]bool, num int) map[int]bool {
	numMap[num] = true
	return numMap
}

func addNumsToMap(numMap map[int]bool, nums []int) map[int]bool {
	for _, num := range nums {
		addNumToMap(numMap, num)
	}
	return numMap
}

func printTest(testCase []int, longest int) {
	fmt.Println("Test Case:", testCase)
	fmt.Println("Longest:", longest)
}

func main() {
	test1 := []int{100, 4, 200, 1, 3, 2}
	test2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	printTest(test1, longestConsecutive(test1))
	printTest(test2, longestConsecutive(test2))
}
