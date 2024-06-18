/*
57. Insert Interval
https://leetcode.com/problems/insert-interval/

You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Note that you don't need to modify intervals in-place. You can make a new array and return it.
*/

package insertinterval

// Time Complexity: O(N)
// Space Complexity: O(N)

func insert(intervals [][]int, newInterval []int) [][]int {
	var (
		index = 0
		cnt   = 0
		res   = make([][]int, len(intervals)+1)
	)

	for ; index < len(intervals) && intervals[index][1] < newInterval[0]; index++ {
		cnt++
	}

	copy(res, intervals[:index])

	start, end := newInterval[0], newInterval[1]

	for ; index < len(intervals) && intervals[index][0] <= newInterval[1]; index++ {
		start = min(start, intervals[index][0])
		end = max(end, intervals[index][1])
	}

	res[cnt] = []int{start, end}

	copy(res[cnt+1:], intervals[index:])

	cnt += len(intervals) - index

	return res[:cnt+1]
}

// Time Complexity: O(N)
// Space Complexity: O(N)

func insert1(intervals [][]int, newInterval []int) [][]int {
	var (
		index        = 0
		newIntervals = [][]int{}
	)

	for index < len(intervals) && intervals[index][0] < newInterval[0] {
		newIntervals = append(newIntervals, intervals[index])
		index++
	}

	// insert the interval
	if len(newIntervals) < 1 || newIntervals[len(newIntervals)-1][1] < newInterval[0] {
		newIntervals = append(newIntervals, newInterval)
	} else {
		newIntervals[len(newIntervals)-1][1] = max(newIntervals[len(newIntervals)-1][1], newInterval[1])
	}

	for index < len(intervals) {
		if newIntervals[len(newIntervals)-1][1] >= intervals[index][0] {
			newIntervals[len(newIntervals)-1][1] = max(newIntervals[len(newIntervals)-1][1], intervals[index][1])
		} else {
			newIntervals = append(newIntervals, intervals[index])
		}
		index++
	}

	return newIntervals
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
