/**
 leetcode 118. Pascal's Triangle

 Given an integer numRows, return the first numRows of Pascal's triangle.

 In Pascal's triangle, each number is the sum of the two numbers directly
 above it as shown
 */
package main

//杨辉三角  理解一下
func generate(numRows int) [][]int {
	res := make([][]int,numRows)
	for i := 0;i < numRows;i++ {
		//创建二维数组 注意这个创建二维数组的长度，容易写错
		arr := make([]int,i + 1)
		res[i] = arr
	}
	for i := 0;i < numRows;i++ {
		res[i][0] = 1
	}
	for i := 1; i < numRows;i++ {
		res[i][i] = 1
	}
	for i := 1; i < numRows;i++ {
		for j := 1;j < i;j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}

	}
	return res

}


func generateII(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 0 {
		return result
	}

	result[0] = []int{1}
	if numRows == 1 {
		return result
	}

	result[1] = []int{1, 1}
	if numRows == 2 {
		return result
	}

	for i := 3; i <= numRows; i++ {
		temp := make([]int, i)
		temp[0] = 1
		temp[i-1] = 1
		for j := 1; j <= i-2; j++ {
			temp[j] = result[i-2][j] + result[i-2][j-1]
		}
		result[i-1] = temp
	}

	return result
}
