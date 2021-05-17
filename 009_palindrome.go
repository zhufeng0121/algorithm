/**
 leetcode 9 Palindrome Number

 Given an integer x, return true if x is palindrome integer.

 An integer is a palindrome when it reads the same backward as
 forward. For example, 121 is palindrome while 123 is not.

 */
package main

func reverse(x int) int{
	result := 0
	for x != 0 {
		result = result * 10 + x%10
		x = x/10
	}
	if result > 1 <<31 -1 || result < -(1<<31 -1) {
		return 0
	}
	return result
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rev := reverse(x)
	if rev != x {
		return false
	} else {
		return true
	}

}
