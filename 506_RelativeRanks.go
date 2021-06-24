/**
 leetcode 506 Relative Ranks

 You are given an integer array score of size n, where score[i]
 is the score of the ith athlete in a competition. All the scores
 are guaranteed to be unique.

 The athletes are placed based on their scores, where the 1st place
 athlete has the highest score, the 2nd place athlete has the 2nd
 highest score, and so on. The placement of each athlete determines
 their rank:

	The 1st place athlete's rank is "Gold Medal".
	The 2nd place athlete's rank is "Silver Medal".
	The 3rd place athlete's rank is "Bronze Medal".
	For the 4th place to the nth place athlete, their rank is their
    placement number (i.e., the xth place athlete's rank is "x").

 Return an array answer of size n where answer[i] is the rank of the
 ith athlete.
 */

package main

import (
	"sort"
	"strconv"
)

//
func findRelativeRanks(score []int) []string {
	order := make([]int,len(score))
	copy(order,score)
	sort.Slice(order, func(i, j int) bool {
		return order[i] > order[j]
	})
	res := make([]string,len(score))
	for i := 0 ; i < len(score);i++ {
		rank := BinarySearch(order,score[i]) + 1
		if rank == 1 {
			res[i] = "Gold Medal"
		} else if rank == 2 {
			res[i] ="Silver Medal"
		} else if rank == 3 {
			res[i] = "Bronze Medal"
		} else {
			res[i] = strconv.Itoa(rank)
		}

	}
	return res

}

//要注意这个数组是倒序的，所以二分查找也需要做些改变
func BinarySearch(nums []int, target int) int  {
	start, end := 0, len(nums) - 1
	for start <= end {
		mid := start + (end - start) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}


//
type MyInt struct {
	Val int
	Index int
}

func findRelativeRanksI(nums []int) []string {
	athletes := make([]MyInt, len(nums))
	for i, num := range nums {
		athletes[i].Index=i
		athletes[i].Val=num
	}
	sort.Slice(athletes, func(i, j int) bool {
		return athletes[i].Val>athletes[j].Val
	})
	rank := make([]string, len(nums))
	for i, athlete := range athletes {
		switch i {
		case 0:
			rank[athlete.Index]="Gold Medal"
		case 1:
			rank[athlete.Index]="Silver Medal"
		case 2:
			rank[athlete.Index]="Bronze Medal"
		default:
			rank[athlete.Index] = strconv.Itoa(i+1)
		}
	}
	return rank

}