/**
 leetcode 48 Rotate Image

 You are given an n x n 2D matrix representing an image, rotate
 the image by 90 degrees (clockwise).

 You have to rotate the image in-place, which means you have to
 modify the input 2D matrix directly. DO NOT allocate another 2D
 matrix and do the rotation.

 */
package main

//这种题的做法就是考循环，还是要熟练

//最好是一圈一圈的旋转，写一个圈旋转的函数，然后再主函数中调用
func rotate(matrix [][]int)  {
	top, bottom := 0, len(matrix) - 1
	left, right := 0, len(matrix[0]) - 1
	for top < bottom {
		rotateLayer(matrix,top,left,bottom,right)
		top++
		left++
		bottom--
		right--
	}


}
//只需要开辟一个tmp的临时数组就可以了
func rotateLayer(matrix [][]int, top int, left int, bottom int,right int) {
	tmp, times := 0, right - left
	for i := 0;i < times;i++ {
		tmp = matrix[top][left + i]
		matrix[top][left + i] = matrix[bottom - i][left]
		matrix[bottom -i][left] = matrix[bottom][right -i]
		matrix[bottom][right-i] = matrix[top +i][right]
		matrix[top + i][right] = tmp

	}
}



