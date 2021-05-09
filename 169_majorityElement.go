/**
 leetcode 169 Majority Element

 Given an array nums of size n, return the majority element.

 The majority element is the element that appears more than ⌊n / 2⌋
 times. You may assume that the majority element always exists in the array.

 */
package main

import "sort"

func majorityElement(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[len(nums)/2]

}

func majorityElement2(nums []int) int {
	temp := nums[0]
	count := 1
	for i := 1 ;i < len(nums);i++ {
		if temp == nums[i] {
			count += 1
		} else {
			if count > 1 {
				count -= 1
			} else {
				temp = nums[i]
			}
		}
	}
	return temp
}