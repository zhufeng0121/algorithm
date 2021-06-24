/**
 leetcode Range Addition II

 You are given an m x n matrix M initialized with all 0's
 and an array of operations ops, where ops[i] = [ai, bi]
 means M[x][y] should be incremented by one for all 0 <= x < ai
 and 0 <= y < bi.

 Count and return the number of maximum integers in the matrix
 after performing all the operations.

 */
package main

//如何不开辟数组，得到加法数量最多的元素个数呢
//求交集，需要对ops 进行排序
func maxCountI(m int, n int, ops [][]int) int {
	width, height := m, n
	for i := 0; i < len(ops); i++ {
		if ops[i][0] < width {
			width = ops[i][0]
		}
		if ops[i][1] < height {
			height = ops[i][1]
		}
	}
	return width * height

}


//这么做内存占用很高，想想别的方法
func maxCount(m int, n int, ops [][]int) int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int,n)
	}
	for i := 0; i < len(ops);i++ {
		for j := 0; j < ops[i][0]; j++ {
			for k := 0; k < ops[i][1]; k++ {
				matrix[j][k]++
			}
		}
	}
	max := 0
	for i := 0; i < m;i++ {
		for j := 0; j < n; j++ {
			if max < matrix[i][j] {
				max =matrix[i][j]
			}
		}
	}
	res := 0
	for i := 0; i < m;i++ {
		for j := 0; j < n; j++ {
			if max ==matrix[i][j] {
				res++
			}
		}
	}

	return res

}


