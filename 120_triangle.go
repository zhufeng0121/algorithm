/**
 leetcode 120 Triangle

 Given a triangle array, return the minimum path sum from top to bottom.

 For each step, you may move to an adjacent number of the row below. More
 formally, if you are on index i on the current row, you may move to either
 index i or index i + 1 on the next row.

 */
package main

//一道简单的动态规划的题目
func minimumTotal(triangle [][]int) int {
	height := len(triangle)
	//width  := len(triangle[len(triangle)])
	// 状态转移方程 min[i][j] 代表的含义为 到 triangle[i][j] 的最小代价和
	min := triangle
	min[0][0] = triangle[0][0]

	for i := 1 ;i < height;i++ {
		min[i][0] = triangle[i][0] + min[i-1][0]
		min[i][len(triangle[i]) - 1] = triangle[i][len(triangle[i]) - 1] + min[i-1][len(triangle[i-1]) - 1]
	}

	for i := 1; i < height; i++ {
		for j := 1 ;j < len(triangle[i]) - 1;j++ {
			min[i][j] = minInt(min[i-1][j], min[i-1][j-1]) + triangle[i][j]
		}
	}

	return minInSlice(min[height-1])


}

func minInSlice(nums []int) int {
	min := nums[0]
	for i := 1;i < len(nums) ;i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min

}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

