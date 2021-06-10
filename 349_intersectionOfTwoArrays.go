/**
 leetcode 349 Intersection of Two Arrays

 Given two integer arrays nums1 and nums2, return an array of their intersection.
 Each element in the result must be unique and you may return the result in any order.

 */
package main

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int,0)
	unique := make(map[int]bool)
	for _, value := range nums1 {
		if _, ok := unique[value]; !ok {
			unique[value] = true
		}
	}
	for _, value := range nums2 {
		if v, ok := unique[value]; ok {
			if v && !inSlice(value,res) {
				res = append(res, value)
			}
		}
	}
	return res

}
func inSlice(target int, nums []int) bool {
	for _, v := range nums {
		if target == v {
			return true
		}
	}
	return false
}


//排序 + 双指针
func intersectionDoublePointers(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			// 如果两个数字相等，且该数字不等于pre,添加
			if res == nil || x > res[len(res) - 1] {
				res = append(res,x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return res
}