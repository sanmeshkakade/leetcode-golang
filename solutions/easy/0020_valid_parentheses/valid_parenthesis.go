/*
20. Valid Parentheses
https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

package validparentheses

// Time Complexity: O(N)
// Space Complexity: O(N)

func isValid(s string) bool {

	var (
		stack   = []rune{}
		mapping = map[rune]rune{'}': '{', ')': '(', ']': '['}
	)

	for _, ch := range s {
		if ch == '}' || ch == ')' || ch == ']' {
			var topElement rune
			if len(stack) > 0 {
				topElement = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				topElement = '#'
			}

			if mapping[ch] != topElement {
				return false
			}
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}
