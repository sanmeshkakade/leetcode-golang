/*
242. Valid Anagram
https://leetcode.com/problems/valid-anagram/

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
*/

package validanagram

// Time Complexity: O(len(s))
// Space Complexity: O(1)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq_counter := [26]int{}

	for i := range s {
		freq_counter[s[i]-'a']++
		freq_counter[t[i]-'a']--
	}

	for _, freq := range freq_counter {
		if freq != 0 {
			return false
		}
	}

	return true
}
