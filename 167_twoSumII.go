/**
 leetcode 167 Two Sum II - Input array is sorted

 Given an array of integers numbers that is already sorted in
 ascending order, find two numbers such that they add up to a
 specific target number.

 Return the indices of the two numbers (1-indexed) as an integer
 array answer of size 2, where 1 <= answer[0] < answer[1] <= numbers.length.

 You may assume that each input would have exactly one solution
 and you may not use the same element twice.

 */
package main


//本题应该还有很多种解法
func twoSumII(numbers []int, target int) []int {

	start := 0
	end := len(numbers) -1

	for start < end {
		if numbers[start] + numbers[end] == target {
			return []int{start + 1,end + 1}
		} else if numbers[start] + numbers[end] > target {
			end--
		} else {
			start++
		}
	}
	return []int{0,0}

}


