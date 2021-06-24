/**
 leetcode 442 Find All Duplicates In Array

 Given an integer array nums of length n where all the integers
 of nums are in the range [1, n] and each integer appears once
 or twice, return an array of all the integers that appears twice.

 You must write an algorithm that runs in O(n) time and uses only
 constant extra space.

 */
package main

//这个题有点意思啦
func findDuplicates(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums);i++ {
		for nums[i] != i + 1 && nums[i] != nums[nums[i] -1] {
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		}

	}
	for i := 0;i < len(nums);i++ {
		if nums[i] != i + 1 {
			res = append(res,nums[i])
		}
	}
	return res
}