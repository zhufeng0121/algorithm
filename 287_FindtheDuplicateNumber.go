/**
 leetcode 287. Find the Duplicate Number

 Given an array of integers nums containing n + 1 integers where each
 integer is in the range [1, n] inclusive.

 There is only one repeated number in nums, return this repeated number.

 You must solve the problem without modifying the array nums and uses
 only constant extra space.

 */
package main

//利用二分法，如何利用二分法？
func findDuplicate(nums []int) int {
	//TODO:
	start := 1
	end := len(nums) - 1
	ans := 0
	for start <= end {
		mid := start + (end - start) >> 1
		cnt := 0
		for i := 0 ; i < len(nums);i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if mid > cnt {
			end = mid - 1
			ans = mid
		} else {
			//mid 是符合条件的
			start = mid + 1

		}
	}
	return ans


}

func findDuplicateI(nums []int) int {
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[fast]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	pre1 := 0
	pre2 := slow
	for slow != fast {
		pre1 = nums[pre1]
		pre2 = nums[pre2]
	}
	return pre1
}


