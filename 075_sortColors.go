/**
 leetcode 75. sort Colors

 Given an array nums with n objects colored red,
 white, or blue, sort them in-place so that objects
 of the same color are adjacent, with the colors in the
 order red, white, and blue.

 We will use the integers 0, 1, and 2 to represent
 the color red, white, and blue, respectively.


 */

package main

func sortColors(nums []int)  {
	red := 0
	blue := len(nums) - 1
	i, temp := 0, 0
	for i <= blue {
		if nums[i] == 0 {
			temp = nums[red]
			nums[red] = nums[i]
			nums[i] = temp
			red ++
			i++
		} else if nums[i] == 1 {
			i++
		} else if nums[i] == 2 {
			temp = nums[blue]
			nums[blue] = nums[i]
			nums[i] = temp
			blue --
		}

	}

}
