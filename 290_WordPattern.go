/**
 leetcode 290 Word Pattern

 Given a pattern and a string s, find if s follows the same pattern.

 Here follow means a full match, such that there is a bijection
 between a letter in pattern and a non-empty word in s.

 */
package main

import "strings"


//思路是对的,果然是需要用两个Map 来实现双射
func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s," ")
	pmap := make(map[byte]string)
	smap := make(map[string]byte)
	if len(pattern) != len(strs) {
		return false
	}
	i := 0
	for i < len(pattern){
		_, ok1 := pmap[pattern[i]]
		_, ok2 := smap[strs[i]]
		if !ok1 && !ok2 {
			pmap[pattern[i]] = strs[i]
			smap[strs[i]] = pattern[i]
			i++
		} else {
			if pmap[pattern[i]] == strs[i] && smap[strs[i]] == pattern[i]{
				i++
			} else {
				return false
			}
		}
	}
	return true
}

//还有其他的不错的实现方法