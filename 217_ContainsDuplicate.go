/**
 leetcode 217 Contains Duplicate

 Given an integer array nums, return true if any value appears at
 least twice in the array, and return false if every element is distinct.

 */
package main

import "sort"

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	uniqueMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := uniqueMap[v]; ok {
			return true
		}
		uniqueMap[v] = true
	}
	return false

}

//也可以使用排序
func containsDuplicateI(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
