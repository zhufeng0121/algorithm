/**
 leetcode 697 Degree of An Array

 Given a non-empty array of non-negative integers nums,
 the degree of this array is defined as the maximum frequency
 of any one of its elements.

 Your task is to find the smallest possible length of a (contiguous)
 subarray of nums, that has the same degree as nums.

 */
package main

//思路是先寻找众数，然后确定该众数第一次和最后一次出现的位置即可
func findShortestSubArray(nums []int) int {
	countMap := make(map[int]int,0)
	mode := make([]int,0)
	max := 0
	for _, v := range nums {
		countMap[v]++
	}
	for _, v := range countMap {
		if v > max {
			max = v
		}
	}
	for i, v := range countMap {
		if v == max {
			mode = append(mode,i)
		}
	}
	max = len(nums)
	for _, v := range mode {
		first, last := 0, len(nums) -1
		for i := 0; i < len(nums) - 1; i++{
			if nums[i] == v {
				first = i
				break
			}
		}
		for j :=len(nums) -1; j >=0; j-- {
			if nums[j] == v {
				last = j
				break
			}
		}
		if max > last - first + 1 {
			max = last -first + 1
		}
	}
	return max

}

//leetcode题解 暂时还没有来的及看
type entry struct{ cnt, l, r int }

func findShortestSubArrayIV(nums []int) (ans int) {
	mp := map[int]entry{}
	for i, v := range nums {
		if e, has := mp[v]; has {
			e.cnt++
			e.r = i
			mp[v] = e
		} else {
			mp[v] = entry{1, i, i}
		}
	}
	maxCnt := 0
	for _, e := range mp {
		if e.cnt > maxCnt {
			maxCnt, ans = e.cnt, e.r-e.l+1
		} else if e.cnt == maxCnt {
			ans = min(ans, e.r-e.l+1)
		}
	}
	return
}