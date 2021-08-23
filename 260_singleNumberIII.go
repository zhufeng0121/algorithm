/**
 leetcode 260 Single Number III

 Given an integer array nums, in which exactly two elements appear
 only once and all the other elements appear exactly twice. Find
 the two elements that appear only once. You can return the answer in any order.

 You must write an algorithm that runs in linear runtime complexity and uses
 only constant extra space.

 */
package main

func singleNumberIII(nums []int) []int {
	res, div := 0, 1
	for i := 0;i < len(nums);i++ {
		res ^= nums[i]
	}
	for (res & div) == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, v := range nums {
		if div & v != 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a,b}
}


