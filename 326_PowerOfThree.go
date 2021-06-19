/**
 leetcode 326 Power of Three

 Given an integer n, return true if it is a
 power of three. Otherwise, return false.

 An integer n is a power of three, if there exists
 an integer x such that n == 3x.

 */
package main

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n % 3 == 0 {
		n /= 3
	}

	if n == 1 {
		return true
	}
	return false
}
