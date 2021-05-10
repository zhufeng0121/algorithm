/**
 leetcode 1748 Sum of Unique Elements

 You are given an integer array nums. The unique elements of an
 array are the elements that appear exactly once in the array.

 Return the sum of all the unique elements of nums.

 */
package main

//此题应该还有其他的解法，这个解法应该不是题目想要的
func sumOfUnique(nums []int) int {
	countMap := make(map[int]int,0)
	sum := 0
	for _, v := range nums {
		countMap[v]++
	}
	for i, v := range countMap {
		if v == 1 {
			sum += i
		}
	}
	return sum

}

func sumOfUnique2(nums []int) int {
	return 0
}



