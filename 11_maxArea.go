/**
  leetcode 11. Container With Most Water

  Given n non-negative integers a1, a2, ..., an ,
  where each represents a point at coordinate (i, ai).
  n vertical lines are drawn such that the two endpoints
  of the line i is at (i, ai) and (i, 0). Find two lines,
  which, together with the x-axis forms a container,
  such that the container contains the most water.
 */
package main

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	max := 0
	for start < end {
		if max < (end -start) * min(height[start], height[end]) {
			max = (end -start) * min(height[start], height[end])
		}
		if height[start] < height[end] {
			start++
		} else {
			end --
		}
	}
	return max
}

func min (a, b int) int {
	if a > b {
		return b
	}
	return a
}

