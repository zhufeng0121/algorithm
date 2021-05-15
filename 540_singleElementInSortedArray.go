/**
 leetcode 540 Single Element in a Sorted Array

 You are given a sorted array consisting of only integers where
 every element appears exactly twice, except for one element which
 appears exactly once. Find this single element that appears only once.

 Follow up: Your solution should run in O(log n) time and O(1) space.

 */
package main

//要求以 O(logN) 时间复杂度进行求解，因此不能遍历数组并进行异或操作来求解，这么做的时间复杂度为 O(N)。

//令 index 为 Single Element 在数组中的位置。在 index 之后，数组中原来存在的成对状态被改变。如果 m 为偶数，并且 m + 1 < index，那么 nums[m] == nums[m + 1]；m + 1 >= index，那么 nums[m] != nums[m + 1]。

//从上面的规律可以知道，如果 nums[m] == nums[m + 1]，那么 index 所在的数组位置为 [m + 2, h]，此时令 l = m + 2；如果 nums[m] != nums[m + 1]，那么 index 所在的数组位置为 [l, m]，此时令 h = m。

//因为 h 的赋值表达式为 h = m，那么循环条件也就只能使用 l < h 这种形式。

//此题有一种解法 用异或操作，但是不太符合出题者的意思，要求O(logn)可以试试binary
func singleNonDuplicate(nums []int) int {
	//TODO:
	return 0
}

