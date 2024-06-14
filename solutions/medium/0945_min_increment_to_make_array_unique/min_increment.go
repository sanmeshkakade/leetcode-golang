/*
945. Minimum Increment to Make Array Unique
https://leetcode.com/problems/minimum-increment-to-make-array-unique/

You are given an integer array nums. In one move, you can pick an index i where 0 <= i < nums.length and increment nums[i] by 1.

Return the minimum number of moves to make every value in nums unique.

The test cases are generated so that the answer fits in a 32-bit integer.
*/

package mitmau

// count sort optimisation
// Time Complexity: O(N)
// Space Complexity: O(N)

func minIncrementForUnique(nums []int) int {

	var count = make([]int, 100_001)
	for _, v := range nums {
		count[v]++
	}

	minIncrement, waitingForPlace := 0, 0
	for _, repeated := range count {
		minIncrement += waitingForPlace
		if repeated > 1 {
			waitingForPlace += repeated - 1
		} else {
			if waitingForPlace > 0 && repeated == 0 {
				waitingForPlace--
			}
		}
	}

	for waitingForPlace > 0 {
		minIncrement += waitingForPlace
		waitingForPlace--
	}

	return minIncrement
}

// Time Complexity: O(N^2)
// Space Complexity: O(N)

func minIncrementForUnique1(nums []int) int {

	var fmap = map[int]int{}
	for _, num := range nums {
		fmap[num]++
	}

	ans, lowest := 0, 0
	for k, v := range fmap {
		lowest := max(lowest, k)
		if v > 1 {
			for v > 1 {
				cur := lowest
				for {
					if _, exist := fmap[cur]; !exist {
						ans += cur - k
						fmap[cur]++
						lowest = cur + 1
						break
					}
					cur++
				}
				v--
			}
		}
	}

	return ans
}
