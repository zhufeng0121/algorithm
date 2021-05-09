/**
 leetcode 136 Single Number

 Given a non-empty array of integers nums, every element appears
 twice except for one. Find that single one.

 Follow up: Could you implement a solution with a linear runtime
 complexity and without using extra memory?

 */
package main

func singleNumber(nums []int) int {
	appear := make(map[int]int,0)

	for i := 0 ;i < len(nums);i++ {
		appear[nums[i]]++

	}
	for i, v := range appear {
		if v == 1 {
			return i
		}
	}
	return 0

}
// better method  use xor operation
func singleNumber2(nums[] int) int {
	result := 0
	for i := 0 ;i < len(nums);i++ {
		result ^= nums[i]
	}
	return result

}
