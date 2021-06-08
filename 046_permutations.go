/**
 leetcode 46  Permutations

 Given an array nums of distinct integers, return all the possible permutations.
 You can return the answer in any order.


 */
package main

func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	used := make([]bool,len(nums))


}
