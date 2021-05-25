/**
 leetcode 28 Implement strStr()

 Implement strStr().

 Return the index of the first occurrence of needle in haystack, or -1
 if needle is not part of haystack.

 Clarification:

 What should we return when needle is an empty string?
 This is a great question to ask during an interview.

 For the purpose of this problem, we will return 0 when needle is an empty
 string. This is consistent to C's strstr() and Java's indexOf().

*/
package main

// 想起了那篇论文，find a needle in haystack
//BF 算法
func strStr(haystack string, needle string) int {
	hay := []rune(haystack)
	nee := []rune(needle)
	if len(hay) < len(nee) {
		return -1
	}
	if len(nee) == 0 {
		return 0
	}
	for i := 0;i < len(hay) - len(nee);i++ {
		j := 0
		for j = 0;j < len(nee);j++ {
			if hay[i + j] != nee[j] {
				break
			}
		}
		if j == len(nee)  {
			return i
		}

	}
	return -1

}

//BF算法的另一种写法 ： 这种写法和书上见到的比较相同，采用回退
func strStrII(haystack string, needle string) int {
	//TODO:
	return -1
}