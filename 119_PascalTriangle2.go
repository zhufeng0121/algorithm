/**
 leetcode 119. Pascal's Triangle II

 Given an integer rowIndex, return the rowIndexth (0-indexed) row of
 the Pascal's triangle.

 In Pascal's triangle, each number is the sum of the two numbers directly
 above it as shown:

*/
package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1,1}
	}
	result := make([]int,rowIndex + 1)
	result[0] = 1
	result[1] = 1
	for i := 2; i <= rowIndex; i++ {
		for j := i - 1; j >= 1; j-- {
			result[j] = result[j] + result[j-1]
		}
		result[i] = 1
	}
	return result

}
