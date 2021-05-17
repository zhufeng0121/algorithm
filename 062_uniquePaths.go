/**
 leetcode 62 Unique Paths

 A robot is located at the top-left corner of a m x n grid
 (marked 'Start' in the diagram below).

 The robot can only move either down or right at any point in
 time. The robot is trying to reach the bottom-right corner
 of the grid (marked 'Finish' in the diagram below).

 How many possible unique paths are there?

 */
package main

func uniquePaths(m int, n int) int {
	var path [][]int
	for x := 0; x < m; x++ {
		arr := make([]int,n)
		path = append(path,arr)
	}
	for i := 0 ;i < m;i++ {
		path[i][0] = 1
	}
	for i := 0; i < n;i++ {
		path[0][i] = 1
	}
	for i := 1;i < m;i++ {
		for j := 1;j < n;j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}

	return path[m-1][n-1]

}

// 看不懂？？
func uniquePathsII(m int, n int) int {
	dp := make([]int,n)
	for i := 0;i < n ;i++ {
		dp[i] = 1
	}
	for  i := 1;i < m;i++ {
		for  j:= 1;j < n;j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]

}

