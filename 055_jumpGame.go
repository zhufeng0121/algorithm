/**
 leetcode 55 Jump Game

 Given an array of non-negative integers nums, you are initially positioned
 at the first index of the array.

 Each element in the array represents your maximum jump length at that position.

 Determine if you are able to reach the last index.

 */
package main

//Greedy
func canJump(nums []int) bool {
	length := len(nums)
	//向右可以跳跃的最远的位置
	rightmost := 0
	// index + nums[index]
	for i := 0; i < len(nums);i++ {
		if i <= rightmost {
			rightmost = max(rightmost, i + nums[i])
			if rightmost >= length -1 {
				return true
			}
		}
	}
	return false

}
func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
