package containsDuplicate

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
