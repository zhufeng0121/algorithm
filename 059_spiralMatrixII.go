/**
 leetcode 59 Spiral Matrix II

 Given a positive integer n, generate an n x n matrix filled with elements
 from 1 to n2 in spiral order.

 */
package main

//有精力可以关注一下其他的for循环遍历方式，实际上就是简单的for循环，理清边界条件即可
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int,n)
	}
	top, left := 0, 0
	bottom, right := n-1, n-1
	value := 1
	for top < bottom && left < right {
		//要对n为奇数是进行处理，最后剩下单行
		for i := left; i <= right; i++ {
			res[top][i] = value
			value++
		}
		for i := top + 1; i <= bottom; i++ {
			res[i][right] = value
			value++
		}
		for i := right-1; i >= left; i-- {
			res[bottom][i] = value
			value++
		}
		for i := bottom -1; i > top;i-- {
			res[i][left] = value
			value++
		}
		//break
		top++
		left++
		right--
		bottom--

	}
	if top == bottom && left == right {
		res[top][left] = n * n
	}
	return res
}