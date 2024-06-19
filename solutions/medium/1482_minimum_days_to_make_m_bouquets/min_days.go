/*
1482. Minimum Number of Days to Make m Bouquets
https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/

You are given an integer array bloomDay, an integer m and an integer k.

You want to make m bouquets. To make a bouquet, you need to use k adjacent flowers from the garden.

The garden consists of n flowers, the ith flower will bloom in the bloomDay[i] and then can be used in exactly one bouquet.

Return the minimum number of days you need to wait to be able to make m bouquets from the garden. If it is impossible to make m bouquets return -1.
*/

package mindays

// Time Complexity: O(NlogN)
// Space Complexity: O(1)

func minDays(bloomDay []int, m int, k int) int {

	if len(bloomDay) < m*k {
		return -1
	}

	var low, high int = 1, 1e9

	for low < high {
		var mid = (low + high) / 2
		if canBouquetsBeMade(bloomDay, mid, m, k) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func canBouquetsBeMade(bloomDay []int, day, m, k int) bool {

	var (
		totalBouquet = 0
		flowerCount  = 0
	)

	for _, bDay := range bloomDay {
		if bDay <= day {
			flowerCount++
		} else {
			flowerCount = 0
		}

		if flowerCount == k {
			flowerCount = 0
			totalBouquet++
		}
	}

	return totalBouquet >= m
}
