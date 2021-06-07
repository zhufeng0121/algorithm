/**
 leetcode 367. Valid Perfect Square

 Given a positive integer num, write a function which returns True if
 num is a perfect square else False.

 Follow up: Do not use any built-in library function such as sqrt.

 */
package main

func isPerfectSquare(num int) bool {
	start, end := 0, num
	for start <= end {
		mid := start + (end - start) >> 1
		if mid * mid == num {
			return true
		} else if mid * mid > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false

}

//第二种解法也可以考虑牛顿迭代法


