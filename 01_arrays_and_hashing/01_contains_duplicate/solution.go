package main

func containsDuplicate(nums []int) bool {
	hashset := make(map[int]int)

	for _, num := range nums {
		hashset[num]++

		if hashset[num] > 1 {
			return true
		}
	}

	// if we get here, then there were no duplicates
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}

	if containsDuplicate(nums) {
		println("The array contains duplicates")
	} else {
		println("The array does not contain duplicates")
	}
}
