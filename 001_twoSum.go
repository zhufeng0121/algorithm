package main

/**
 leetcode 001. Two Sum

 Given an array of integers nums and an integer target,
 return indices of the two numbers such that they add up to target.

 You may assume that each input would have exactly one solution,
 and you may not use the same element twice.

 You can return the answer in any order.

 */


func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	length := len(nums)
	for i := 0;i < length;i++ {
		another := target - nums[i]
		if _ ,ok := m[another];ok {
			return []int{i,m[another]}
		}
		m[nums[i]] = i

	}
	return nil

}

func twoSumI(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another],i}
		}
		m[nums[i]] = i
	}
	return nil
}