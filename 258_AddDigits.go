/**
 leetcode 258. Add Digits

 Given an integer num, repeatedly add all its digits until the result
 has only one digit, and return it.

*/
package main

//这个居然写不出来，要好好反思反思了
func addDigits(num int) int {
	for {
		tmp := 0
		for num > 0 {
			tmp += num %10
			num = num / 10
		}
		if tmp < 10 {
			return tmp
		}
		num = tmp
	}

}
//The original number is divisible by 9 if and only if
//the sum of its digits is divisible by 9
//这道题的第二种解法很有意思，试试理解一下
func addDigitsIV(num int) int {
	if num == 0 {
		return 0
	}
	if num % 9 == 0 {
		return 9
	}
	return num % 9

}


