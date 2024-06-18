/*
542. 01 Matrix
https://leetcode.com/problems/01-matrix

Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.
*/
package matrix

import (
	"math"
)

// bfs
// Time Complexity: O(N*M)
// Space Complexity: O(N*M)

func updateMatrix(mat [][]int) [][]int {

	var q = make([][]int, 0)

	for row := range mat {
		for col := range mat[0] {
			if mat[row][col] == 0 {
				q = append(q, []int{row, col})
			} else {
				mat[row][col] = math.MaxInt
			}
		}
	}
	var directions = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for len(q) > 0 {
		row, col := q[0][0], q[0][1]
		q = q[1:]

		for _, dir := range directions {
			nRow := row + dir[0]
			nCol := col + dir[1]

			if nRow > -1 && nRow < len(mat) && nCol > -1 && nCol < len(mat[0]) && mat[nRow][nCol] > mat[row][col]+1 {
				q = append(q, []int{nRow, nCol})
				mat[nRow][nCol] = mat[row][col] + 1
			}
		}
	}

	return mat
}
