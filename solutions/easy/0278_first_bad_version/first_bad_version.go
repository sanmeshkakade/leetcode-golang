/*
278. First Bad Version
https://leetcode.com/problems/first-bad-version/

You are a product manager and currently leading a team to develop a new product.
Unfortunately, the latest version of your product fails the quality check.
Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad.
Implement a function to find the first bad version. You should minimize the number of calls to the API.
*/

package firstbadversion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var isBadVersion func(version int) bool

// Time Complexity: O(log N)
// space Complexity: O(1)

func firstBadVersion(n int) int {
	var low, high = 1, n
	var mid int

	for low <= high {
		mid = (low + high) / 2

		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
