/**
 leetcode 56 Merge Intervals

 Given an array of intervals where intervals[i] = [starti, endi],
 merge all overlapping intervals, and return an array of the
 non-overlapping intervals that cover all the intervals in the input.

 */
package main

import "sort"

// 解法是 按照每个区间的左边界进行排序，然后分别对比，逐一合并即可

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
    // 存放第一个元素
	res := [][]int{intervals[0]}
	last := 0
	for i := 1 ;i < len(intervals);i++ {
		if res[last][1] >= intervals[i][0] {
			res[last][1] = max(res[last][1],intervals[i][1])
		} else {
			res = append(res,intervals[i])
			last++
		}
	}
	return res


}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


