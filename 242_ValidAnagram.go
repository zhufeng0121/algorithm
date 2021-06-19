/**
 leetcode 242. Valid Anagram

 Given two strings s and t, return true if t is an anagram of s,
 and false otherwise.

 */
package main

import "sort"
//sort
func isAnagramIII(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arrs := []byte(s)
	arrt := []byte(t)
	sort.Slice(arrs, func(i, j int) bool {
		return arrs[i] < arrs[j]
	})
	sort.Slice(arrt, func(i, j int) bool {
		return arrt[i] < arrt[j]
	})

	for i := 0; i < len(s);i++ {
		if arrs[i] != arrt[i] {
			return false
		}
	}
	return true

}
//Map
func isAnagramI(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countMap := make(map[byte]int)
	for i := 0; i < len(s);i++ {
		countMap[s[i]]++
	}
	for i := 0; i < len(t);i++ {
		if _, ok := countMap[t[i]]; !ok {
			return false
		}
		countMap[t[i]]--
	}
	for _, v := range countMap {
		if v != 0 {
			return false
		}
	}
	return true

}





