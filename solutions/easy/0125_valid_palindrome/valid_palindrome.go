/*
125. Valid Palindrome
https://leetcode.com/problems/valid-palindrome/

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

package validpalindrome

// Time Complexity: O(len(s))
// Space Complexity: O(len(s))

func isPalindrome(s string) bool {
	str := sanatise(s)

	low, high := 0, len(str)-1

	for low < high {
		if str[low] != str[high] {
			return false
		}
		low++
		high--
	}
	return true
}

func sanatise(s string) string {
	var strRune = []rune{}

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
			strRune = append(strRune, ch)
		} else if ch >= 'A' && ch <= 'Z' {
			strRune = append(strRune, 'a'+(ch-'A'))
		}
	}
	return string(strRune)
}
