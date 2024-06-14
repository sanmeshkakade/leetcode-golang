/*
169. Majority Element
https://leetcode.com/problems/majority-element/

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
*/

package majorityelement

// Time Complexity: O(N)
// Space Complexity: O(1)

func majorityElement(nums []int) int {
	num, freq := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != num {
			if freq == 1 {
				num = nums[i]
			} else {
				freq--
			}
		} else {
			freq++
		}
	}
	return num
}

// Time Complexity: O(N)
// Space Complexity: O(N)
func majorityElement1(nums []int) int {
	numMap := map[int]int{}

	for _, num := range nums {
		numMap[num]++
	}

	for num, freq := range numMap {
		if freq > len(nums)/2 {
			return num
		}
	}
	return -1
}
