/*
75. Sort Colors
https://leetcode.com/problems/sort-colors/

Given an array with n objects colored red, white or blue,
sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.
*/

// date: 12-06-24

package sortcolors


// time complexity: O(n)
// space complexity: O(1)

func sortColors(nums []int) {
	low, high := -1, len(nums)

	for i := range nums {
		if nums[i] == 0 {
			low++
		} else if nums[i] == 2 {
			high--
		}
	}

	for i := range nums {
		if i <= low {
			nums[i] = 0
		} else if i >= high {
			nums[i] = 2
		} else {
			nums[i] = 1
		}
	}
}
