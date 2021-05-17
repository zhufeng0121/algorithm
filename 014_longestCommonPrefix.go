/**
 leetcode 14 Longest Common Prefix

 Write a function to find the longest common prefix
 string amongst an array of strings.

 If there is no common prefix, return an empty string "".

 */

package main

func longestCommonPrefix(strs []string) string {
	min,index := 200,0
	for i := 0;i < len(strs);i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
			index = i
		}
	}
	var temp string
    temp = strs[0]
    strs[0] = strs[index]
    strs[index] = temp
	length, count := 0,0
	bench := strs[0]
    for i := 0;i < len(strs[0]);i++ {
    	for j := 1;j < len(strs);j++ {
			if []rune(strs[j])[i] == []rune(bench)[i] {
				count++
			}
		}
		if count == len(strs) -1 {
			length++
			count = 0
		} else {
			break
		}

	}
	return bench[:length]

}
