/**
 leetcode 84  Largest Rectangle in Histogram

 Given an array of integers heights representing the histogram's bar height where
 the width of each bar is 1, return the area of the largest rectangle in the histogram.

*/
package MonotoneStack


//仍然利用单调栈来解决 单调栈仍然存放下标
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	stack := make([]int,0)
    result := 0
	for i := 0;i < len(heights);i++ {
		for len(stack) != 0 && heights[stack[len(stack) -1]] >= heights[i] {
			//临时存放栈顶元素
			temp := heights[stack[len(stack) -1]]
			stack = stack[:len(stack) -1]
			sidx := -1
			if len(stack) > 0 {
				sidx = stack[len(stack) -1]
			}
			result = maxInt(result, temp * (i- sidx -1))

		}
		stack = append(stack,i)
	}

	return result

}
func maxInt(a,b int) int {
	if a > b {
		return a
	}
	return b
}