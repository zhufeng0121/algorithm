/**
 leetcode 389 Find the Difference

 You are given two strings s and t.

 String t is generated by random shuffling string s and
 then add one more letter at a random position.

 Return the letter that was added to t.

 */
package main

func findTheDifference(s string, t string) byte {
	res := byte(0)
	for i := 0 ;i < len(s); i++ {
		res ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		res ^= t[i]
	}
	return res
}
