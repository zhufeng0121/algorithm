/**
 leetcode 507 Perfect Number

 A perfect number is a positive integer that is equal to the sum
 of its positive divisors, excluding the number itself. A divisor
 of an integer x is an integer that can divide x evenly.

 */
package main

import "math"

//可能还有其他解法，时间复杂度过高
func checkPerfectNumber(num int) bool {
	sum := 0
	for i := 1 ;i <= num/2 ;i++ {
		if num % i == 0 {
			sum += i

		}
	}
	return sum == num

}

func checkPerfectNumberII(num int) bool {
	if num == 4 || num <= 1 {
		return false
	}
	sum := 0
	for i := 1 ;i < int(math.Sqrt(float64(num))) ;i++ {
		if num % i == 0 {
			sum = sum + i + num/i
		}
	}
	return num == sum

}





