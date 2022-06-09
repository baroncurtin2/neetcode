package main

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	hashset := make(map[byte]int)

	// if we get here, we know that the
	// lengths of both strings are equal
	for i, _ := range s1 {
		hashset[s1[i]]++
		hashset[s2[i]]--
	}

	for _, v := range hashset {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "anagram"
	s2 := "nagaram"

	if isAnagram(s1, s2) {
		println("The two strings are anagrams")
	} else {
		println("The two strings are not anagrams")
	}
}
