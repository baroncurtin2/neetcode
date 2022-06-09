package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	hashset := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		hashset[key] = append(hashset[key], str)
	}

	result := make([][]string, 0)

	for _, v := range hashset {
		result = append(result, v)
	}

	return result
}

func sortString(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	for _, v := range result {
		println(strings.Join(v, ","))
	}
}
