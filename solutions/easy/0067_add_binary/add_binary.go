/*
67. Add Binary
https://leetcode.com/problems/add-binary/

Given two binary strings a and b, return their sum as a binary string.
*/

package addbinary

// Time Complexity: O(max(len(a), len(b)))
// Space Complexity: O(max(len(a), len(b)))

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		return addBinary(b, a)
	}

	var (
		lenA = len(a)
		lenB = len(b)
		revA = getReversedArray(a)
		revB = getReversedArray(b)

		ans   = make([]byte, max(lenA))
		carry = false
	)

	for i := 0; i < lenA; i++ {
		sum := revA[i] - '0'

		if i < lenB {
			sum += revB[i] - '0'
		}

		if carry {
			sum += 1
		}

		if sum > 1 {
			carry = true
			sum -= 2
		} else {
			carry = false
		}

		ans[lenA-1-i] = '0' + sum
	}

	if carry {
		return "1" + string(ans)
	}
	return string(ans)
}

func getReversedArray(s string) []byte {
	var ans = []byte{}
	for i := len(s) - 1; i > -1; i-- {
		ans = append(ans, s[i])
	}
	return ans
}
