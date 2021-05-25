/**
 leetcode 3 Longest Substring Without Repeating Characters

 Given a string s, find the length of the longest substring without repeating characters.
 */
package DoublePointers

import "math"

//最长无重复子串的长度


//思路： 采用双指针（或者滑动窗口）
//1.遍历字符串，如果当前字符不在map之中,说明从未出现过，那么直接加入到map之中，同时更新长度
//2.如果出现在map之中，假设 index < start 说明出现的 不在 [start,end]窗口中，那么将该字符出现的位置进行更新，同时增加长度
//3.如果index >= start,说明出现在了[start,end]窗口中，那么将start = index + 1,同时将该字符串出现的位置进行更新


//这里有一个问题，假如输入为 “txxmast”，当右指针right移动到最后一位t时，此时的子串为"xmas"，
//按理说 t 并没有出现在这个子串当中，此时的字符串长度应为5。但是由于之前的"t"在哈希表中的索引还是0，
//因此最后一个"t"会因为已经出现在哈希表里，而被当成重复字符不进入计算，导致bug。
// 因此，针对此边界问题，我们需要再加入一个判断：若当前字符已经存在哈希表中，但它是上一个字符串里的字符
//（即它在哈希表中的索引小于当前right），则当前字符不属于当前子串，仍计算长度。

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := math.MinInt32
	start := 0
	// map 中存放下标
	stringMap := make(map[rune]int,0)
	str := []rune(s)
	for end := 0;end < len(str);end++ {
		index, ok := stringMap[str[end]]
		if !ok {
			res = maxInt(res, end - start + 1)
		} else  {
			if index < start {
				res = maxInt(res, end - start + 1)
			} else {
				start = index + 1
			}
		}
		stringMap[str[end]] = end

	}
	return res

}

func maxInt(a,b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstringII(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := math.MinInt32
	start := 0
	// map 中存放下标
	stringMap := make(map[rune]int,0)
	str := []rune(s)
	for end := 0;end < len(str);end++ {
		index, ok := stringMap[str[end]]
		if !ok || ok && index < start {
			res = maxInt(res, end - start + 1)
		} else if ok && index >= start  {
			start = index + 1
		}
		stringMap[str[end]] = end

	}
	return res

}
