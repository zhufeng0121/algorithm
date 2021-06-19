/**
 leetcode 125 Valid palindrome

 Given a string s, determine if it is a palindrome, considering only
 alphanumeric characters and ignoring cases.

 */
package main

import "strings"

//double pointers
func isPalindromeIV(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s) -1
	//
	for i < j {
		//如果s[i]既不是大写字母，也不是小写字母
		for i < j && !isalnum(s[i]) {
			i++
		}
		//时时刻刻要比较i和j的大小关系
		for i < j && !isalnum(s[j]) {
			j--
		}
		if i < j {
			if s[i] != s[j] {
				return false
			}
		}
		i++
		j--
	}
	return true

}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') ||(ch >= 'a' && ch <= 'z') || (ch >= '1' && ch <= '9')

}
