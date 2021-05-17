/**
 leetcode 48 Rotate Image

 You are given an n x n 2D matrix representing an image, rotate
 the image by 90 degrees (clockwise).

 You have to rotate the image in-place, which means you have to
 modify the input 2D matrix directly. DO NOT allocate another 2D
 matrix and do the rotation.

 */
package main

func rotate(matrix [][]int)  {
	width := len(matrix)
	for width > 1 {
		temp := make([]int,width)
		for i := 0;i < width;i++ {
			temp[i] = matrix[0][i]
		}
		for i := width -1 ;i >= 0 ;i -- {
			matrix[0][i] = matrix[width -1 -i][0]
		}
		for i := width -1 ;i >= 0 ;i -- {
			matrix[width - 1 - i][0] = matrix[width - 1][width - 1 - i]
		}
		for i := width -1 ;i >= 0 ;i -- {
			matrix[width -1][width -1 - i] = matrix[i][width -1]
		}
		for i := width -1 ;i >= 0 ;i -- {
			matrix[i][width -1] = temp[i]
		}

		width = width - 2
	}

}



