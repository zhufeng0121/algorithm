/**
 leetcode 205 Isomorphic Strings

 Given two strings s and t, determine if they are isomorphic.

 Two strings s and t are isomorphic if the characters in s can be replaced
 to get t.

 All occurrences of a character must be replaced with another character
 while preserving the order of characters. No two characters may map to
 the same character, but a character may map to itself.

 */
package main

// isomorphic 同源异构
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[byte]byte)
	map2 := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		value, ok := map1[s[i]]
		if ok {
			if value != t[i]{
				return false
			}
		} else {
			if _, ok := map2[t[i]] ; ok {
				return false
			}
			map1[s[i]] = t[i]
			map2[t[i]] = true
		}
	}
	return true

}
