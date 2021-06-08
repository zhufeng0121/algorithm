/**
 leetcode 97 Interleaving String

 Given strings s1, s2, and s3, find whether s3 is formed by an interleaving of s1 and s2.

 An interleaving of two strings s and t is a configuration where they are divided into
 non-empty substrings such that:

 s = s1 + s2 + ... + sn
 t = t1 + t2 + ... + tm
 |n - m| <= 1
 The interleaving is s1 + t1 + s2 + t2 + s3 + t3 + ... or t1 + s1 + t2 + s2 + t3 + s3 + ...
 Note: a + b is the concatenation of strings a and b.
 */
package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	//TODO:
	n, m := len(s1), len(s2)
	if len(s3) != n + m {
		return false
	}
	//声明一个n +1 * m + 1 bool 类型的数组， dp[i][j] = true 代表的含义是s3 中前长度为 i + j + 1
	//由s1[i] 和 s2[j] 组成，
	dp := make([][]bool,n + 1)
	for i := 0;i < len(dp);i++ {
		dp[i] = make([]bool,m + 1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j-1] == s3[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i-1] == s3[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || dp[i][j-1] && s2[j-1] == s3[i+j-1]
			}

		}
	}

	//所以返回dp[n][m]即可，那么dp[n][m] = dp[n][m -1] (if s3[n + m] == s1[n] || s3[n + m] == s2[m])
	//那么如何确定边界条件呢
	return dp[n][m]

}

//还可以对空间进行优化 将二维的动态规划的数组转换成一维的
func isInterleaveI(s1 string, s2 string, s3 string) bool {
	return false
}
