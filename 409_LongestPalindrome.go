/**
 leetcode 409. Longest Palindrome

 Given a string s which consists of lowercase or uppercase
 letters, return the length of the longest palindrome that
 can be built with those letters.

 Letters are case sensitive, for example, "Aa" is not
 considered a palindrome here.
 */
package main

func longestPalindrome(s string) int {
	smap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		smap[s[i]]++
	}
	sum := 0
	flag := false
	for _, v := range smap {
		if v & 1 == 1 {
			flag = true
		}
		sum += v / 2
	}
	if flag {
		return sum * 2 + 1
	}
	return sum * 2

}
