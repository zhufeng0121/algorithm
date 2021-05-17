/**
 leetcode 50 Pow(x,n)

 Implement pow(x, n), which calculates x raised to the power n (i.e., xn).


 */
package main

//根据n的奇偶性，可以用递归方法来进行
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	t := myPow(x, n/2)
	if n % 2 == 1 {
		return t * t * x
	} else {
		return t * t
	}
}

//快速幂算法
func myPowII(x float64 , n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	sign := 1
	if n < 0 {
		sign = -1
	}
	ans := 1.0
	n = abs(n)
	for n != 0 {
		if n & 1 == 1 {
			ans = ans * x
		}
		x = x * x
		n >>= 1
	}
	if sign == -1 {
		return 1/ans
	}
	return ans

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}


