/**
 leetcode 455 Assign Cookies

 Assume you are an awesome parent and want to give your children some cookies.
 But, you should give each child at most one cookie.

 Each child i has a greed factor g[i], which is the minimum size of a cookie
 that the child will be content with; and each cookie j has a size s[j]. If
 s[j] >= g[i], we can assign the cookie j to the child i, and the child i will
 be content. Your goal is to maximize the number of your content children and
 output the maximum number.

 */
package main

import "sort"
// GreedyAlgorithm algorithm
//对 饼干大小和 孩子的饿程度进行排序
func findContentChildren(g []int, s []int) int {
	result := 0
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			result += 1
			i++
		}
		j++
	}
	return result

}

