/**
 leetcode 434. Number of Segments in a String

 You are given a string s, return the number of segments in the string.

 A segment is defined to be a contiguous sequence of non-space characters.
 */
package main

func countSegments(s string) int {
	res := 0
	for i := 0; i < len(s) ;i++ {
		if (s[i] != ' ') && (i == 0 || s[i-1] == ' ') {
			res++
		}

	}
	return res
}

//计算单词的数量，等同于计数单词开始的下标个数

//若该下标前为空格（或者为初始下标），且自身不为空格，
//则其为单词开始的下标
func countSegmentsII(s string) int {
	res := 0
	for i := 0; i < len(s);i++ {
		if (i== 0 || s[i-1] == ' ') && s[i] != ' ' {
			res++
		}
	}
	return res

}