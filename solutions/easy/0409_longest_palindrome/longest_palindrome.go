/*
409. Longest Palindrome
https://leetcode.com/problems/longest-palindrome/

Given a string s which consists of lowercase or uppercase letters, return the length of the longest
palindrome
that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome.
*/
package longestpalindrome

// Time Complexity: O(N)
// Space Complexity: O(1)

func longestPalindrome(s string) int {
	var char_counter = [52]int{}

	for _, ch := range s {
		if ch >= 'a' {
			char_counter[ch-'a'+26]++
		} else {
			char_counter[ch-'A']++
		}
	}

	var (
		ans    = 0
		hasOdd = false
	)

	for _, v := range char_counter {
		ans += v
		if v%2 == 1 {
			ans--
			hasOdd = true
		}
	}

	if hasOdd {
		ans++
	}

	return ans
}
