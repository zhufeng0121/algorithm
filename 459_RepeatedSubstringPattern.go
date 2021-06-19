/**
 leetcode 459  Repeated Substring Pattern

 Given a string s, check if it can be constructed by taking a substring of
 it and appending multiple copies of the substring together.

 */

package main

import "go/ast"

//思路：
func repeatedSubstringPattern(s string) bool {
	arr := []byte(s)
	for i := 1;i < len(s);i++ {
		tmp := circleShift(i,arr)
		if tmp == arr {
			return true
		}
	}
	return false
}
func circleShift(n int, arr []byte) []byte {
	reverseI(arr[:n])
	reverseI(arr[n+1:])
	reverseI(arr)
	return arr

}
func reverseI(arr []byte) []byte {
	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[len(arr) - 1 - i] = arr[len(arr) - 1 - i], arr[i]
	}
}
func bruteforce(s string, t string) bool {

}