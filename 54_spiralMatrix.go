/**
 leetcode 54 Spiral Matrix

 Given an m x n matrix, return all elements of the matrix in spiral order.

 */
package main

// 转圈打印矩阵
func spiralOrder(matrix [][]int) []int {
	height := len(matrix)
	width  := len(matrix[0])
	result := make([]int,0)
	row1, column1 := 0, 0
	row2, column2 := height-1 ,width-1

	for row1 <= row2 && column1 <= column2 {
		if row1 == row2 {
			//矩阵只有一行
			for i := column1;i <= column2;i++ {
				result = append(result,matrix[row1][i])
			}
		} else if column1 == column2 {
			//矩阵只有一列
			for i := row1; i<= row2;i++ {
				result = append(result,matrix[i][column1])
			}
		} else {

			for i:= column1 ;i <= column2;i++ {
				result = append(result,matrix[row1][i])
			}
			for i := row1;i <= row2;i++ {
				result = append(result,matrix[i][column2])
			}
			for i := column2 ;i >= column1;i-- {
				result = append(result,matrix[row2][i])

			}
			for i := row2; i >= row1;i-- {
				result = append(result,matrix[i][column1])
			}
			row2--
			column2--
			row1++
			column1++
		}

	}
	return result

}

// 旋转打印  终于过了！！ 第一版留着吧，时刻告诉我哪里出了问题
func spiralOrderII(matrix [][]int) []int {
	top, left := 0, 0
	bottom, right := len(matrix) - 1,len(matrix[0]) - 1
	result := make([]int,0)

	for top <= bottom && left <= right {
		if top == bottom {
			//矩阵只有一行
			for i := left; i <= right;i++ {
				result = append(result,matrix[top][i])

			}
			top++
		} else if left == right {
			//矩阵只有一列
			for i := top; i <= bottom;i++ {
				result = append(result,matrix[i][left])
			}
			left++
		} else {
			for i := left;i <= right;i++ {
				result = append(result,matrix[top][i])
			}
			for i := top + 1; i <= bottom;i++ {
				result = append(result,matrix[i][right])
			}
			for i := right -1; i >= left;i-- {
				result = append(result,matrix[bottom][i])
			}
			for i := bottom -1; i > top;i-- {
				result = append(result,matrix[i][left])
			}
			top++
			left++
			right--
			bottom--
		}
	}
	return result

}


