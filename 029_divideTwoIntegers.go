/**
 leetcode 29. Divide Two Integers

 Given two integers dividend and divisor, divide two integers without
 using multiplication, division, and mod operator.

 Return the quotient after dividing dividend by divisor.

 The integer division should truncate toward zero, which means losing
 its fractional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.

 Note: Assume we are dealing with an environment that could only store
 integers within the 32-bit signed integer range: [−231, 231 − 1]. For
 this problem, assume that your function returns 231 − 1 when the division
 result overflows.

 */
package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}
	if math.Abs(float64(divisor)) > math.Abs(float64(dividend)) {
		return 0
	}
	//越界
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	// 记录两个数的符号
	f1, f2 := false, false

	if dividend < 0 {
		dividend = - dividend
		f1 = true
	}
	if divisor < 0 {
		divisor = - divisor
		f2 = true
	}
	//使用快速加法
	res := 0
	base := 1
	x := divisor
	for x <= dividend {
		base = 1
		res += base
		for x + x <= dividend {
			x += x
			res += base
			base += base
		}
		dividend -= x
		x = divisor

	}
	if f1 != f2 {
		res = -res
	}
	return res

}
