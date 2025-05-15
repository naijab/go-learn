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
	maxLen := 0
	startIndex := 0

	for i := 0; i < n; i++ {
		vis := make(map[byte]bool)
		for j := i; j < n; j++ {
			// check duplicate
			if vis[s[j]] {
				break
			}
			vis[s[j]] = true

			// Now:
			// Start index: i = 2
			// End index: j = 4
			// Length of "cde" = 4 - 2 + 1 = 3

			// check  max length
			if j-i+1 > maxLen {
				maxLen = j - i + 1
				startIndex = i
			}
		}
	}

	return s[startIndex : startIndex+maxLen]
}
