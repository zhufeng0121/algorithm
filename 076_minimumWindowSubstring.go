/**
 leetcode 76  Minimum Window Substring

 Given two strings s and t of lengths m and n respectively, return the
 minimum window in s which will contain all the characters in t. If
 there is no such window in s that covers all characters in t, return
 the empty string "".

 Note that If there is such a window, it is guaranteed that there will
 always be only one unique minimum window in s.
 */
package main
// Sliding Window
func minWindow(s string, t string) string {
	//TODO:
	return ""
}
func include(s []rune, t []rune) bool {
	stringMap := make(map[rune]bool,len(s))
	for _, v := range s {
		if _,ok := stringMap[v]; !ok {
			stringMap[v] = true
		}
	}
	for _, v := range t {
		if _,ok := stringMap[v]; !ok {
			return false
		}
	}
	return true
}
