/*
53. Maximum Subarray
https://leetcode.com/problems/maximum-subarray/

Given an integer array nums, find the subarray with the largest sum, and return its sum.
*/

package maximumsubarray

import "math"

// Time Complexity: O(N)
// Space Complexity: O(1)
// Kadane's Algorithm

func maxSubArray(nums []int) int {
	var (
		maxSum = math.MinInt
		sum    = 0
	)

	for index := range nums {
		sum += nums[index]
		if nums[index] > sum {
			sum = nums[index]
		}
		if maxSum < sum {
			maxSum = sum
		}
	}
	return maxSum
}

// Time Complexity: O(N^2)
// Space Complexity: O(N)

func maxSubArray1(nums []int) int {
	var (
		maxSum = math.MinInt
		sum    = 0
		preSum = make([]int, len(nums)+1)
	)

	for index := range nums {
		preSum[index] = sum
		sum += nums[index]
	}

	preSum[len(nums)] = sum

	for i := 1; i < len(preSum); i++ {
		for j := 0; j < i; j++ {
			maxSum = max(maxSum, preSum[i]-preSum[j])
		}
	}

	return maxSum
}

// Time Complexity: O(N^3)
// Space Complexity: O(1)

func maxSubArray2(nums []int) int {
	var maxSum = math.MinInt

	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			sum := 0
			for k := j; k <= i; k++ {
				sum += nums[k]
			}
			maxSum = max(maxSum, sum)
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
