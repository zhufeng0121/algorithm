/**
 leetcode 74 Search a 2D Matrix

 Write an efficient algorithm that searches for a value
 in an m x n matrix. This matrix has the following properties:

 Integers in each row are sorted from left to right.
 The first integer of each row is greater than the last integer of the previous row.

 */

package main

import(
	_ "fmt"
)
// 将二维数组看做是一维数组，并进行二分查找
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, low, high := len(matrix[0]),0, len (matrix[0]) * len(matrix) -1
	for low <= high {
		mid := low + (high - low) >> 1
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}


// 1. 从矩阵最右上角的数开始寻找 2. 比较当前数matrix[row][col] 和 target的关系
// 2. 如果和target相等，已找到，如果比target大，col -- 如果比target小，row++
func searchMatrix2(matrix [][]int, target int) bool {
	row := 0
	column := len(matrix[0]) - 1

	for row < len(matrix) && column > -1 {
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] > target {
			column--
		} else {
			row++
		}
	}
	return false
}
