/**
 leetcode 70 climbling stairs

 You are climbing a staircase. It takes n steps to reach the top.

 Each time you can either climb 1 or 2 steps. In how many distinct
 ways can you climb to the top?

 */
package main

//递归
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)

}

func climbStairs2(n int) int {
	a, b := 0, 1
	count := 0
	for n > 0 {
		count = a + b
		a = b
		b = a + b
		n--
	}
	return count
}


func climbStairs3(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2;i <= n;i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs4(number int) int {
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	a, b := 1, 2
	res := 0
	for i := 3; i <= number; i++ {
		res = a + b
		a = b
		b = res
	}
	return res

}