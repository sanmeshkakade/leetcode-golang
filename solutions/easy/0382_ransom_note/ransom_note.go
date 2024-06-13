/*
383. Ransom Note
https://leetcode.com/problems/ransom-note/

Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine
and false otherwise.

Each letter in magazine can only be used once in ransomNote.
*/
package ransomnote

// Time Complexity: O(N) where N is length of magazine
// Space Complexity: O(1)

func canConstruct(ransomNote string, magazine string) bool {

	if len(magazine) < len(ransomNote) {
		return false
	}

	var char_counter = [26]rune{}
	var ransomLen = len(ransomNote)
	for i := range magazine {
		char_counter[magazine[i]-'a']++
		if i < ransomLen {
			char_counter[ransomNote[i]-'a']--
		}
	}

	for _, v := range char_counter {
		if v < 0 {
			return false
		}
	}

	return true
}
