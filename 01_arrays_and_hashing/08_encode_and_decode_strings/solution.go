package main

import (
	"fmt"
	"strconv"
)

func encoder(s string) string {
	return fmt.Sprintf("%d#%s", len(s), s)
}

func encode(strings []string) string {
	var encoded string
	for _, s := range strings {
		encoded += encoder(s)
	}
	return encoded
}

func decoder(s string, startIdx int, endIdx int) (int, string) {
	// get word length
	wordLen, _ := strconv.Atoi(s[startIdx:endIdx])

	// get word
	word := s[endIdx+1 : endIdx+1+wordLen]

	return wordLen, word
}

func decode(s string) []string {
	var decoded []string

	for startIdx := 0; startIdx < len(s); {
		endIdx := startIdx

		// iterate until we find '#' to get word length
		for s[endIdx] != '#' {
			endIdx++
		}

		wordLen, word := decoder(s, startIdx, endIdx)

		// append word to results
		decoded = append(decoded, word)

		// update start_idx
		startIdx = endIdx + wordLen + 1
	}

	return decoded
}

func printTest(testCase []string, encoded string, decoded []string) {
	fmt.Println("Test Case:", testCase)
	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decoded)

}

func main() {
	test1 := []string{"lint", "code", "love", "you"}
	test2 := []string{"we", "say", ":", "yes"}

	encoded1 := encode(test1)
	decoded1 := decode(encoded1)
	printTest(test1, encoded1, decoded1)

	encoded2 := encode(test2)
	decoded2 := decode(encoded2)
	printTest(test2, encoded2, decoded2)
}
