/*
217. Contains Duplicate
https://leetcode.com/problems/contains-duplicate/

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/

package containsduplicate

// Time Complexity: O(N)
// Space Complexity: O(N)

func containsDuplicate(nums []int) bool {
	var record = make(map[int]struct{})

	for _, val := range nums {
		if _, exists := record[val]; exists {
			return true
		}
		record[val] = struct{}{}
	}

	return false
}
