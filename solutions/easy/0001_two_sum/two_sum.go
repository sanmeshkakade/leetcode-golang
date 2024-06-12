/*
1. Two Sum
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

*/

package twosum

// Time Complexity: O(N)
// Space Complexity: O(N)

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for index, num := range nums {
		complement := target - num
		if compIndex, exists := record[complement]; exists {
			return []int{compIndex, index}
		}
		record[num] = index
	}

	return []int{}
}

// Time Complexity: O(N^2)
// Space Complexity: O(1)

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
