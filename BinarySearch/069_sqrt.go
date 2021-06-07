/**
 leetcode 69 Sqrt(x)

 Given a non-negative integer x, compute and return the square root of x.

 Since the return type is an integer, the decimal digits are truncated,
 and only the integer part of the result is returned.

 Note: You are not allowed to use any built-in exponent function or operator,
 such as pow(x, 0.5) or x ** 0.5.
 */
package main

import "math"

func mySqrt(x int) int {
	start, end := 0, x
	for start <= end {
		mid := start + (end - start) >> 1
		if mid * mid == x {
			return mid
		} else if mid * mid > x {
			end = mid -1
		} else {
			if (mid + 1) * (mid +1) == x {
				return mid + 1
			} else if (mid + 1) * (mid +1) < x {
				start = mid + 1
			} else {
				return mid
			}
		}
	}
	return 0

}

func mySqrtI(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
//牛顿迭代法
func mySqrtNewton(x int) int {
	if x == 0 {
		return 0
	}
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0 - xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}


