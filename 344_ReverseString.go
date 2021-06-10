/**
 leetcode 344 Reverse String

 Write a function that reverses a string. The input string is
 given as an array of characters s.
 */
package main

func reverseString(s []byte)  {
	if len(s) == 0 || len(s) == 1 {
		return
	}
	for i := 0 ;i < len(s)/2; i++ {
		s[i], s[len(s) -1 - i] = s[len(s) -1 - i], s[i]
	}

}
