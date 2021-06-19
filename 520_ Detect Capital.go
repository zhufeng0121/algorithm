/**
 leetcode 520 Detect Capital

 We define the usage of capitals in a word to be right when one of the following cases holds:

	All letters in this word are capitals, like "USA".
	All letters in this word are not capitals, like "leetcode".
	Only the first letter in this word is capital, like "Google".
	Given a string word, return true if the usage of capitals in it is right.
 */
package main

func detectCapitalUse(word string) bool {
	if len(word) == 0 || len(word) == 1 {
		return true
	}
	if isCapital(word[0]) {
		//如果开头字母是小写,那么后面要么都是小写，要么都是大写
		if allCapital(word[1:]) || allLowercase(word[1:]) {
			return true
		}
		return false
	} else {
		//如果开头字母是小写,那么后面不能出现大写
		if allLowercase(word[1:]) {
			return true
		}
		return false
	}

}
func isCapital(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
func isLowercase(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false

}

func allCapital(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isCapital(s[i]) {
			return false
		}
	}
	return true
}
func allLowercase(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isLowercase(s[i]) {
			return false
		}
	}
	return true
}
//统计大小写字母的数量，与字符串的长度做比较，当大写字母或者小写
//字母的数量等于字符串的长度直接返回true
//当第一个字符为大写时，小写字母的数量等于字符串长度减1也满足条件直接返回即可;

func detectCapitalUseI(word string) bool {
	upper, lower := 0, 0
	for i := 0; i < len(word); i++ {
		if isCapital(word[i]) {
			upper++
		} else {
			lower++
		}


	}
	if upper == len(word) || lower == len(word) {
		return true
	}
	if isCapital(word[0]) && lower == len(word) - 1 {
		return true
	}
	return false

}