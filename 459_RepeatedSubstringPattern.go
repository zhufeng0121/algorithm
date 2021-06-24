/**
 leetcode 459  Repeated Substring Pattern

 Given a string s, check if it can be constructed by taking a substring of
 it and appending multiple copies of the substring together.

 */

package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i != 0 {
			continue
		}
		if substring(s, s[:i]) {
			return true
		}
	}

	return false
}

func substring(s, sub string) bool {
	if len(s) == 0 {
		return true
	}

	if s[:len(sub)] != sub {
		return false
	}

	return substring(s[len(sub):], sub)
}


func repeatedSubstringPatternI(s string) bool {
	return strings.Contains(strings.Repeat(s, 2)[1:2*len(s)-1],s)

}
