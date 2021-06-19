/**
 leetcode 412 Fizz Buzz

 Given an integer n, return a string array answer (1-indexed) where:

	answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
	answer[i] == "Fizz" if i is divisible by 3.
	answer[i] == "Buzz" if i is divisible by 5.
	answer[i] == i if non of the above conditions are true.

 */
package main

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string,n)
	for i := 1; i <= n ;i++ {
		if i % 15 == 0 {
			res[i-1] = "FizzBuzz"
		} else if i % 3 == 0 {
			res[i-1] = "Fizz"
		} else if i % 5 == 0 {
			res[i-1] = "Buzz"
		} else {
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res

}
