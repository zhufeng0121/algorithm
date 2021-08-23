/**
 leetcode 583 Delete Operation for Two Strings

 Given two strings word1 and word2, return the minimum number
 of steps required to make word1 and word2 the same.

 In one step, you can delete exactly one character in
 either string.
 */
package main

//本题是求删除次数最少，使得两个字符串相等，实际上就是求两个串的最长公共子序列的长度，然后用两个串的长度相减即可
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m + 1; i++ {
		dp[i] = make([]int,n + 1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1],dp[i-1][j])
			}
		}
	}
	return m + n - 2 * dp[m][n]

}
