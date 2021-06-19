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

//我们判断 s 和 t 每个位置上的字符是否都一一对应， 即 s 的任意一个字符被 t 中唯一的字符对应，同时 t
//的任意一个字符被 s 中唯一的字符对应。这也被称为「双射」的关系 ,t 中的字符 a 和 r 虽然有唯一的映射
//o，但对于 s 中的字符 o 来说其存在两个映射,故不满足条件

//因此，我们维护两张哈希表，第一张哈希表 以 s 中字符为键，映射至 t 的字符为值，
//第二张哈希表以t中字符为键，映射至 s 的字符为值。从左至右遍历两个字符串的字符，
//不断更新两张哈希表，如果出现冲突（即当前下标index 对应的字符 s[index] 已经
//存在映射且不为 t[index] 或当前下标 index 对应的字符 t[index] 已经存在映射
//且不为 s[index]时说明两个字符串无法构成同构，返回false。
//如果遍历结束没有出现冲突，则表明两个字符串是同构的，返回 true 即可

func isIsomorphicIV(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		value1, ok := s2t[s[i]]
		if ok {
			if value1 != t[i] {
				return false
			}
		} else {
			if _, ok := t2s[t[i]]; ok {
				return false
			}
			s2t[s[i]] = t[i]
			t2s[t[i]] = s[i]
		}
	}
	return true

}

