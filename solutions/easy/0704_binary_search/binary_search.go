/*
704. Binary Search
https://leetcode.com/problems/binary-search/

Given an array of integers nums which is sorted in ascending order, and an integer target,
write a function to search target in nums. If target exists, then return its index. Otherwise,
return -1.

You must write an algorithm with O(log n) runtime complexity.
*/

package binarysearch

// Time Complexity: O(log n)
// Space Complexity: O(1)

func search(nums []int, target int) int {

	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		}
		if nums[mid] > target {
			high = mid - 1
		}
	}

	return -1
}
