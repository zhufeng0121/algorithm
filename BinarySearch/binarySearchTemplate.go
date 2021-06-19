package main

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return 0


}

//查找需要访问数组中当前索引以及直接右邻居索引的元素
func binarySearchII(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left != len(nums) && nums[left] == target {
		return left
	}
	return -1
}
