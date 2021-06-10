/**
 leetcode 231 Power of Two

 Given an integer n, return true if it is a power of two.
 Otherwise, return false.An integer n is a power of two, if there exists
 an integer x such that n == 2x.

 Follow up: Could you solve it without loops/recursion?

 */
package main


func isPowerOfTwo(n int) bool {
	count := 0
	for n != 0 {
		if n & 1 == 1 {
			count++
		}
		n >>= 1
	}
	return count == 1

}

// n & (n - 1) == 0
func isPowerOfTwoI(n int) bool {
	if n == 0 {
		return false
	}
	return n & (n - 1) == 0
}
