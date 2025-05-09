package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	result := ""

	for i := 0; i < len(s); i++ {
		// Check for odd-length palindrome centered at i
		odd := expandAroundCenter(s, i, i)
		fmt.Println("odd: ", odd)

		// Check for even-length palindrome centered between i and i+1
		even := expandAroundCenter(s, i, i+1)
		fmt.Println("even: ", even)

		// Update longest if a longer palindrome is found
		if len(odd) > len(result) {
			result = odd
		}

		if len(even) > len(result) {
			result = even
		}
	}

	return result
}

func expandAroundCenter(s string, left, right int) string {
	// start check palindrome from middle and expand outwards left and right
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]
}

func main() {
	input := "xabaz"
	output := longestPalindrome(input)
	fmt.Println("input : ", input)
	fmt.Println("output : ", output)
}
