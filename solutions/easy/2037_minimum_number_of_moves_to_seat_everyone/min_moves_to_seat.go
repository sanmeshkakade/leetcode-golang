/*
2037. Minimum Number of Moves to Seat Everyone
https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/

There are n seats and n students in a room. You are given an array seats of length n, where seats[i] is the position of the ith seat. You are also given the array students of length n, where students[j] is the position of the jth student.

You may perform the following move any number of times:

Increase or decrease the position of the ith student by 1 (i.e., moving the ith student from position x to x + 1 or x - 1)
Return the minimum number of moves required to move each student to a seat such that no two students are in the same seat.

Note that there may be multiple seats or students in the same position at the beginning.
*/

package minmoves

import "sort"

// Time Complexity: O(n)
// Space Complexity: O(1)

func minMovesToSeat(seats []int, students []int) int {
	var buckets = [101]int{}

	for i := 0; i < len(students); i++ {
		buckets[students[i]]++
		buckets[seats[i]]--
	}

	var (
		moves               = 0
		studentsWithoutSeat = 0
	)

	for _, v := range buckets {
		moves += max(studentsWithoutSeat, -studentsWithoutSeat)
		studentsWithoutSeat += v
	}

	return moves
}

// Time Complexity: O(n log n)
// Space Complexity: O(n) for sorting slices

func minMovesToSeat1(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	ans := 0

	for i := 0; i < len(students); i++ {
		if moves := seats[i] - students[i]; moves > 0 {
			ans += moves
		} else {
			ans -= moves
		}

	}

	return ans
}
