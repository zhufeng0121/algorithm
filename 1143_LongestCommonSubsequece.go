/**
 leetcode 1143  Longest Common Subsequence

 Given two strings text1 and text2, return the length
 of their longest common subsequence. If there is no
 common subsequence, return 0.

 A subsequence of a string is a new string generated
 from the original string with some characters (can be none)
 deleted without changing the relative order of the remaining
 characters.

 For example, "ace" is a subsequence of "abcde".
 A common subsequence of two strings is a subsequence that
 is common to both strings.

*/
package main

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	//dp[i][j] 以 nums1[i]和nums2[j]为结尾的两个数组中公共的、长度最长的子数组的长度
	dp := make([][]int,n + 1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m + 1)
	}
	for i := 1 ; i <= n; i++ {
		for j := 1; j <= m ;j++ {
			dp[i][j] = max(dp[i-1][j],dp[i][j-1])

		}
	}
	return dp[n][m]

}



