/**
 leetcode 435 Non-overlapping Intervals

 Given an array of intervals intervals where intervals[i] = [starti, endi],
 return the minimum number of intervals you need to remove to make the rest
 of the intervals non-overlapping.

 */
package main

import (
	"sort"
)

// greedy algorithm  按照区间的右边界大小进行排序
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return 0
	}
	erase, i := 0, 1
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	//利用一个变量记录下每次的end，防止进行数组中元素的删除操作
	end := intervals[0][1]
	for i < len(intervals) {
		if end > intervals[i][0] {
			erase++
		} else {
			end = intervals[i][1]
		}
		i++
	}
	return erase
}
