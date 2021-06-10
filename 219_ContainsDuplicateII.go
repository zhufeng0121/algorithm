/**
 leetcode 219 Contains Duplicate II

 Given an integer array nums and an integer k, return true if
 there are two distinct indices i and j in the array such that
 nums[i] == nums[j] and abs(i - j) <= k.

 */
package main

//检查长度为k + 1的窗口中是否有相同的元素
func containsNearbyDuplicate(nums []int, k int) bool {
	//思路应该考虑双指针 + map
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	left, right := 0, 0
	unique := make(map[int]bool)
	for right < len(nums) {
		if _, ok := unique[nums[right]]; ok {
			return true
		}
		//注意更新right
		right++
		unique[nums[right]] = true
		if right - left > k {
			delete(unique,nums[left])
			left++
		}
	}
	return false

}
