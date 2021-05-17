/**
 leetcode 64 Minimum Path Sum

 Given a m x n grid filled with non-negative numbers, find a path from top
 left to bottom right, which minimizes the sum of all numbers along its path.

 Note: You can only move either down or right at any point in time.

 */
package main

func minPathSum(grid [][]int) int {
	height := len(grid)
	width  := len(grid[0])
   // arr := grid
	var arr [][]int
	for x := 0; x < height;x++ {
		ar := make([]int, width)
		arr = append(arr,ar)
	}
	arr[0][0] = grid[0][0]
	for i := 1;i < width;i ++ {
		arr[0][i] = grid[0][i] + arr[0][i-1]
	}
	for i := 1;i < height;i++ {
		arr[i][0] = grid[i][0] + arr[i-1][0]
	}
	for i := 1;i < height;i++ {
		for j := 1;j < width;j++ {
			arr[i][j] = minInt(arr[i-1][j], arr[i][j-1]) + grid[i][j]
		}
	}
	return arr[height-1][width -1]

}
