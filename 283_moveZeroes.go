/**
 leetcode 283. Move Zeroes

 Given an integer array nums, move all 0's to the end of it while
 maintaining the relative order of the non-zero elements.

 Note that you must do this in-place without making a copy of the array.

 */
package main

// 哇！ 这个题给了一个很好的提示，不要想着交换，而是想着把不是0的元素不断向前移
//那么自然 是0的元素就都在后面了 非常tricky
func moveZeroes(nums []int) {
	index := 0
	for i := 0;i < len(nums);i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index++
	}

}

//仔细体会这个做法的原因
func moveZeroesII(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i := 0; i < len(nums);i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++

		}
	}
}


// 还能继续改进，参见摩尔投票算法