/**
 leetcode 96 Unique Binary Search Trees

 Given an integer n, return the number of structurally unique BST's
 (binary search trees) which has exactly n nodes of unique values from 1 to n.


 */
package main

//numTrees(n) = numTrees(n-1) + numTrees(n-2) * numTrees(1) + ...
//就是catalan数
func numTrees(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	sum := 0
	for i := 1; i < n - 1; i++ {
		sum += numTrees(i) * numTrees(n-1-i)
	}
	sum += 2* numTrees(n-1)
	return sum

}
// 计算catalan数
func numTreesIV(n int) int {
	arr := make([]int,n + 1)
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] * (4 * i -2)/(i + 1)
	}
	return arr[n]
}

func numTreesIII(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			//不用递归，记录下这种形式
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}