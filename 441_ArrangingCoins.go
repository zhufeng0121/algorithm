/**
 leetcode 441. Arranging Coins

 You have n coins and you want to build a staircase
 with these coins. The staircase consists of k rows
 where the ith row has exactly i coins. The last row
 of the staircase may be incomplete.

 Given the integer n, return the number of complete rows
 of the staircase you will build.


 */
package main

func arrangeCoins(n int) int {
	for i :=0;i <= n;i++ {
		if i * (i + 1) <= 2 * n && 2 * n < (i + 1) * (i + 2) {
			return i
		}
	}
	return -1
}
