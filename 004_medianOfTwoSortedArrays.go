/**
 leetcode 4 Median of Two Sorted Arrays

 Given two sorted arrays nums1 and nums2 of size m and n respectively,
 return the median of the two sorted arrays.

 The overall run time complexity should be O(log (m+n)).

 */
package main

// nums1 [1,3,5,7] nums2 [2,4,6]

//当 m+n 是奇数时，中位数是两个有序数组中的第 (m+n)/2 个元素，
//当 m+n 是偶数时，中位数是两个有序数组中的第 (m+n)/2 个元素和
//第 (m+n)/2+1 个元素的平均值。因此，这道题可以转化成寻找两个有
//序数组中的第 k 小的数，其中 k 为 (m+n)/2 或 (m+n)/2+1。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total & 1 == 1 {
		mid := total / 2
		return float64(getKthElement(nums1,nums1,mid + 1))
	} else {
		mid1, mid2 := total/2 -1 , total/2
		return (float64(getKthElement(nums1,nums1,mid1 + 1)) + float64(getKthElement(nums1,nums1,mid2 + 1)))/2.0
	}

}
//
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}
		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k/2
		newIndex1 := min(index1 + half, len(nums1)) - 1
		newIndex2 := min(index2 + half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}




