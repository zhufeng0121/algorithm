/**
 leetcode 202 : Happy Number

 Write an algorithm to determine if a number n is happy.

 A happy number is a number defined by the following process:

 Starting with any positive integer, replace the number by the
 sum of the squares of its digits.

 Repeat the process until the number equals 1 (where it will stay),
 or it loops endlessly in a cycle which does not include 1.

 Those numbers for which this process ends in 1 are happy.
 Return true if n is a happy number, and false if not.

 */
package main

//如果 平方和相加之后的和 出现过，那么就说明 陷入循环，永远也不可能归为1
func isHappy(n int) bool {
	visited := make(map[int]bool)
	for n != 1 {
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = true
		m := 0
		temp := n
		for temp > 0 {
			last := temp % 10
			m = m + last * last
			temp = temp /10
		}
		n = m

	}
	return true

}
