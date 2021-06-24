/**
 leetcode 566 Reshape Matrix

 In MATLAB, there is a handy function called reshape which
 can reshape an m x n matrix into a new one with a different
 size r x c keeping its original data.

 You are given an m x n matrix mat and two integers r and c
 representing the row number and column number of the wanted
 reshaped matrix.

 The reshaped matrix should be filled with all the elements of
 the original matrix in the same row-traversing order as they were.

 If the reshape operation with given parameters is possible and
 legal, output the new reshaped matrix; Otherwise, output the
 original matrix.
 */

package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	row, col := len(mat), len(mat[0])
	if row * col != r * c {
		return mat
	}
	res := make([][]int,r)
	for i := 0; i < r; i++ {
		res[i] = make([]int,c)
	}
	//这个坐标转换 一定要算对
	for i := 0; i < r ;i++ {
		for j := 0; j < c; j++ {
			pos := i * c + j
			res[i][j] = mat[pos /col][pos % col]

		}
	}
	return res

}
