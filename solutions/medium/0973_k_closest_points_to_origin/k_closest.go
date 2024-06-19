/*
973. K Closest Points to Origin
https://leetcode.com/problems/k-closest-points-to-origin/

Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).
*/

package kclosest

// Time Complexity: O(N log N)
// Space Complexity: O(N)

func kClosest(points [][]int, k int) [][]int {
	var res = make([][]int, 0, k)

	var heap = BuildHeap(points)

	for i := 0; i < k; i++ {
		res = append(res, heap.Pop())
	}

	return res
}
