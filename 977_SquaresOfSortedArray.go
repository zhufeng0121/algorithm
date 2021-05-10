/**
 leetcode 977 Squares of a Sorted Array

 Given an integer array nums sorted in non-decreasing order, return
 an array of the squares of each number sorted in non-decreasing order.


 */
package main

import "sort"

//在O(n)时间内完成
func sortedSquares(nums []int) []int {
	increase := make([]int,len(nums))
	start, end := 0, len(nums) - 1
	index := len(nums) - 1
	for start <= end {
		if nums[start] * nums[start] < nums[end] * nums[end] {
			increase[index] = nums[end] * nums[end]
			index--
			end--
		} else {
			increase[index] = nums[start] * nums[start]
			index--
			start++
		}
	}
	return increase

}

func sortedSquaresII(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	sort.Ints(nums)
	return nums
}

func sortedSquaresIII(nums []int) []int {
	front := 0
	for front < len(nums) && nums[front] < 0 {
		front++
	}
	back := front - 1
	increase := make([]int,0)

	for back >= 0 && front < len(nums) {
		if nums[back] * nums[back] < nums[front] * nums[front] {
			increase = append(increase, nums[back] * nums[back])
			back --
		} else {
			increase = append(increase, nums[front] * nums[front])
			front++
		}
	}

	for back >= 0 {
		increase = append(increase, nums[back] * nums[back])
		back --

	}
	for front < len(nums) {
		increase = append(increase, nums[front] * nums[front])
		front ++
	}
	return increase

}