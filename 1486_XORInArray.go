/**
 leetcode 1486 XOR Operation in an Array

 Given an integer n and an integer start.

 Define an array nums where nums[i] = start + 2*i (0-indexed) and n == nums.length.

 Return the bitwise XOR of all elements of nums.

 */
package main

// 没看懂此题在考察什么
func xorOperation(n int, start int) int {
	nums := make([]int,0)
	for i := 0;i < n ; i++ {
		nums = append(nums,start + 2 * i)
	}
	result := nums[0]

	for i := 1; i < len(nums);i++  {
		result ^= nums[i]
	}
	return result

}
