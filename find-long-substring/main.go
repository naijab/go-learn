package main

import "fmt"

func main() {
	input := []string{"abab", "xxxxaxabcdexyz1234xabcxx"}

	for _, value := range input {
		fmt.Println("input: ", value)
		output := findLongestSubString(value)
		fmt.Println("output: ", output)
		fmt.Println("-------")
	}

}

func findLongestSubString(s string) string {
	n := len(s)
	charIndex := make(map[byte]int)
	start := 0
	maxLen := 0
	startIndex := 0

	for end := 0; end < n; end++ {
		ch := s[end]

		// If the character is repeated, move start
		if idx, found := charIndex[ch]; found && idx >= start {
			start = idx + 1
		}

		charIndex[ch] = end

		if end-start+1 > maxLen {
			maxLen = end - start + 1
			startIndex = start
		}
	}

	return s[startIndex : startIndex+maxLen]
}
