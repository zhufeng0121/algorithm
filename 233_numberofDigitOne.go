/**
 leetcode 233 Number of Digit One

 Given an integer n, count the total number of digit 1 appearing in all
 non-negative integers less than or equal to n.


 */
package main
//最简单的一种做法，但超时
func countDigitOne(n int) int {
	num := 0
	for i := 0 ;i <= n;i++ {
		num += numberOne(i)
	}
	return num

}

func numberOne(n int) int {
	count := 0
	for n != 0 {
		if n % 10 == 1{
			count++
		}
		n /= 10
	}
	return count
}

func countDigitOneII(n int) int {

}
