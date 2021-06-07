/**
 leetcode 80. Remove Duplicates from Sorted Array

 Given a sorted array nums, remove the duplicates in-place such that duplicates
 appeared at most twice and return the new length.

 Do not allocate extra space for another array; you must do this by modifying
 the input array in-place with O(1) extra memory.

 */
package main
// two pointers
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1{
		return 1
	}
	res := 0
	slow, fast := 2, 2
	for fast

}
// 可以将此问题迁移至 删除有序数组重复项 保留两位置保留K位

//思路：
//首先我们先让前 2 位直接保留，得到 1,1
//对后面的每一位进行继续遍历，能够保留的前提是与当前位置的前面 k 个元素不同（答案中的第一个 1），因此我们会跳过剩余的
//1，将第一个 2 追加，得到 1,1,2
//继续这个过程，这时候是和答案中的第 2 个 1 进行对比，因此可以得到 1,1,2,2
//这时候和答案中的第 1 个 2 比较，只有与其不同的元素能追加到答案，因此剩余的 2 被跳过，3 被追加到答案：1,1,2,2,3

func removeDuplicatesK(nums []int,k int) int {
	res := 0
	for _, v := range nums {
		if res < k || nums[res - k] != v {
			nums[res] = v
			res++
		}
	}
	return res
}