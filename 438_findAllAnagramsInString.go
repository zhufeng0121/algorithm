/**
 leetcode 438. Find All Anagrams in a String

 Given two strings s and p, return an array of all the start
 indices of p's anagrams in s. You may return the answer in any
 */
package main

//虽然通过了，可是算法比较慢，考虑一下是否有其他优化方法呢？
//利用滑动窗口，寻找异构词（滑动窗口和双指针比较像，因为都是用两个位置表示窗口大小）
func findAnagrams(s string, p string) []int {
	pFreq := make([]int,26)
	for _, v := range p {
		pFreq[v - 'a']++
	}
	res := make([]int,0)

	left, right := 0, 0
	for right < len(s) {
		if (right -left + 1) != len(p) {
			right++
		} else if !isAnagrams(s[left:right+1],pFreq) {
			left++
		} else {
			res = append(res,left)
			left++
		}
	}
	return res

}

//判断是否是同分异构词
func isAnagrams(window string, pFreq []int) bool {
	wFreq := make([]int,26)
	for _, v := range window {
		wFreq[v - 'a']++
	}
	for i := 0;i < 26;i++ {
		if wFreq[i] != pFreq[i] {
			return false
		}
	}
	return true
}

func findAnagramsII(s string, p string) []int {
	//TODO:
	return nil
}
