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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodFill(t *testing.T) {
	type args struct {
		image [][]int
		sr    int
		sc    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{
				image: [][]int{
					{1, 1, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				sr:    1,
				sc:    1,
				color: 2,
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "two",
			args: args{
				image: [][]int{
					{0, 0, 0},
					{0, 0, 0},
				},
				sr:    0,
				sc:    0,
				color: 0,
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "three",
			args: args{
				image: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 1},
				},
				sr:    2,
				sc:    2,
				color: 2,
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 2},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := floodFill(test.args.image, test.args.sr, test.args.sc, test.args.color)

			assert.Equal(t, test.want, got)
		})
	}
}
