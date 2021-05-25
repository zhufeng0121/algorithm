/**
 leetcode 49 Group Anagrams

 Given an array of strings strs, group the anagrams together.
 You can return the answer in any order.

 An Anagram is a word or phrase formed by rearranging the letters
 of a different word or phrase, typically using all the original
 letters exactly once.

 */
package main

import "sort"

//将同分异构词进行分组
func groupAnagrams(strs []string) [][]string {
	//TODO:
	var res [][]string
	//strMap := make(map[string]bool,0)
	//for i := 0;i < len(strs);i++ {
	//	for j :=0 ;j < i;j++ {
	//		if isAnagram(strs[i],strs[j]) {
	//		}
	//	}
	//}
	return res

}

func isAnagram(s1 string, s2 string) bool {
	arr1 := make([]int, 26)
	for _, v := range s1 {
		arr1[v - 'a']++
	}
	arr2 := make([]int,26)
	for _, v := range s2 {
		arr2[v - 'a']++
	}
	for i:= 0;i < 26;i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}


//解法二： 利用排序

type sortByte []byte

func (s sortByte) Less(i,j int) bool {
	return s[i] < s[j]
}
func (s sortByte) Swap(i,j int) {
	s[i], s[j] = s[j], s[i]

}
func (s sortByte) Len() int {
	return len(s)
}
func SortString(str string) string {
	chars := []byte(str)
	sort.Sort(sortByte(chars))
	return string(chars)
}

func groupAnagramsI(strs []string) [][]string {
	var result [][]string
	if len(strs) == 0 {
		return result
	}

	dict := make(map[string]int)
	for _, s := range strs {
		sorted := SortString(s)
		if idx, ok := dict[sorted]; ok {
			result[idx] = append(result[idx], s)
		} else {
			dict[sorted] = len(result)
			result = append(result, []string{s})
		}
	}

	return result
}
