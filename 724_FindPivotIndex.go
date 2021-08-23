/**
 leetcode 724 Find Pivot Index

 Given an array of integers nums, calculate the pivot
 index of this array.

 The pivot index is the index where the sum of all the numbers
 strictly to the left of the index is equal to the sum of all
 the numbers strictly to the index's right.

 If the index is on the left edge of the array, then the left
 sum is 0 because there are no elements to the left. This also
 applies to the right edge of the array.

 Return the leftmost pivot index. If no such index exists, return -1.
 */
package main

//利用前缀和
func pivotIndex(nums []int) int {
	presum := make([]int,len(nums))
	//获取前缀和
	presum[0] = nums[0]
	for i := 1; i < len(nums);i++ {
		presum[i] = presum[i-1] + nums[i]
	}
	for i := 0; i < len(nums);i++ {
		if i == 0 {
			if presum[len(presum) - 1] - presum[0] == 0 {
				return i
			} else {
				continue
			}
		}
	   if presum[i-1] == presum[len(presum) - 1] - presum[i] {
			return i
		}
	}
	return -1

}


//利用前缀和更好的写法
func pivotIndexI(nums []int) int {
	presum := 0
	for _, v := range nums {
		presum += v
	}
	leftsum := 0
	for i := 0; i < len(nums); i++ {
		if leftsum == presum - nums[i] - leftsum {
			return i
		}
		leftsum += nums[i]
	}
	return -1

}
