/**
 leetcode 263 Ugly Number

 An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

 Given an integer n, return true if n is an ugly number.

*/
package main

func isUgly(n int) bool {
	for n != 1 {
		if n % 2 == 0 {
			n /= 2
			continue
		} else if n % 3 == 0 {
			n /= 3
			continue

		} else if n % 5 == 0 {
			n /= 5
			continue
		}
		return false
	}
	return true

}
//
func isUglyII(num int) bool {
	if num <= 0 {
		return false
	}

	primes := []int{2, 3, 5}
	for _, prime := range primes {
		for num != 1 && num/prime*prime == num {
			num = num / prime
		}

		if num == 1 {
			return true
		}
	}

	return false
}