/**
 leetcode 485 Max Consecutive Ones

 Given a binary array nums, return the maximum number of
 consecutive 1's in the array.

*/
package main

func findMaxConsecutiveOnes(nums []int) int {
	max, cur := 0, 0
	for i := 0; i < len(nums);i++ {
		if nums[i] == 1 {
			cur++
		} else {
			max = maxInt(max,cur)
			cur = 0
		}
		if i == len(nums) - 1 {
			max = maxInt(max,cur)
		}

	}
	return max

}
