/*
733. Flood Fill
https://leetcode.com/problems/flood-fill/

An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

You are also given three integers sr, sc, and color. You should perform a flood fill on
the image starting from the pixel image[sr][sc].

To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally
to the starting pixel of the same color as the starting pixel,
plus any pixels connected 4-directionally to those pixels (also with the same color),
and so on. Replace the color of all of the aforementioned pixels with color.

Return the modified image after performing the flood fill.
*/

package floodfill

var (
	delRow = [4]int{0, 0, 1, -1}
	delCol = [4]int{1, -1, 0, 0}
)

// Recursion DFS
// Time Complexity: O(N*M)
// Space Complexity: O(N*M)

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	oldColor := image[sr][sc]

	if oldColor == color {
		return image
	}

	image[sr][sc] = color

	for i := range delRow {
		nrow := sr + delRow[i]
		ncol := sc + delCol[i]

		if nrow > -1 && nrow < len(image) && ncol > -1 && ncol < len(image[0]) && image[nrow][ncol] == oldColor {
			floodFill(image, nrow, ncol, color)
		}
	}
	return image
}
