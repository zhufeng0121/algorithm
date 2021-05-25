/**
 leetcode 42 Trapping Rain Water

 Given n non-negative integers representing an elevation map where the width
 of each bar is 1, compute how much water it can trap after raining.

 */
package main

//利用单调栈的解法
func trap(height []int) int {
	if len(height) == 0 || len(height) == 1 || len(height) == 2{
		return 0
	}
	sum := 0
	// 这个单调栈的栈里面存的是数组的index
	stack := make([]int, 0)
	stack = append(stack, 0)
	for i:= 1;i < len(height);i++ {
		for len(stack) != 0 {
			// 获取栈顶元素对应的index
			if height[stack[len(stack) -1]] >= height[i] {
				stack = append(stack,i)
				break
			} else {
				//获取栈顶元素值
				temp := height[stack[len(stack) -1]]
				//对栈顶元素值进行出栈操作
				stack = stack[:len(stack)-1]
				sum += (minInt(height[stack[len(stack) -1]],height[i]) - temp )*(i - 1 - stack[len(stack) -1])

			}

		}
	}
	return sum

}
