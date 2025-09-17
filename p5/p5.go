package p5

import "strings"

func Run(s string) string {
	return longestPalindrome(s)
}

func longestPalindrome(s string) string {
	maxLen := 0
	retPalindrome := ""

	start := 0
	end := 0

	splitString := strings.Split(s, "")

	for key, _ := range splitString {
		left, right := key, key

		for left >= 0 && right < len(splitString) && splitString[left] == splitString[right] {
			if (right - left + 1) > maxLen {
				maxLen = right - left + 1
				start = left
				end = right
			}
			left -= 1
			right += 1
		}

		left, right = key, key+1
		for left >= 0 && right < len(splitString) && splitString[left] == splitString[right] {
			if (right - left + 1) > maxLen {
				maxLen = right - left + 1
				start = left
				end = right
			}
			left -= 1
			right += 1
		}
	}

	retPalindrome = strings.Join(splitString[start:end+1], "")

	return retPalindrome
}
