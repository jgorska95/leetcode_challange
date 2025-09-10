package p3

/*
Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func Run(s string) int {
	return lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	maxStringLenght := 0
	start := 0

	compareMap := make(map[rune]int)

	for end, val := range s {
		pos, exist := compareMap[val]
		if exist && pos >= start {
			//move start to next letter
			start = pos + 1
		}

		tmpMax := end - start + 1
		if tmpMax > maxStringLenght {
			maxStringLenght = tmpMax
		}

		compareMap[val] = end
	}

	return maxStringLenght
}
