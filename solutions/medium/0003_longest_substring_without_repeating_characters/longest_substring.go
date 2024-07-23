/*
3. Longest Substring Without Repeating Characters
https://leetcode.com/problems/longest-substring-without-repeating-characters/

Given a string s, find the length of the longest substring without repeating characters.
*/

package longestsubstring

func lengthOfLongestSubstring(s string) int {
	var (
		record  = [128]int{}
		longest = 0
		low     = 0
	)

	for index := range record {
		record[index] = -1
	}

	for high := range s {
		if record[s[high]] != -1 {
			low = max(low, record[s[high]]+1)
		}
		record[s[high]] = high
		if high-low+1 > longest {
			longest = high - low + 1
		}
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
