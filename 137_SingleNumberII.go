/**
 leetcode 137 Single Number II

 Given an integer array nums where every element appears three times except
 for one, which appears exactly once. Find the single element and return it.

 */

package main

func singleNumberII(nums []int) int {
	appear := make(map[int]int,0)
	for i := 0 ;i < len(nums);i++ {
			appear[nums[i]]++
	}
	for i, v := range appear {
		if v == 1 {
			return i
		}
	}
	return -1

}

// better method : use no extra space  也可以用位操作，但是还没想好
func singleNumberII_2(nums []int) {

}