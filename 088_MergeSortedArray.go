/**
 leetcode 88 Merge Sorted Array

 Given two sorted integer arrays nums1 and nums2, merge nums2 into
 nums1 as one sorted array.

 The number of elements initialized in nums1 and nums2 are m and n
 respectively. You may assume that nums1 has a size equal to m + n
 such that it has enough space to hold additional elements from nums2.

 */
package main

//从后向前遍历即可
func mergeII(nums1 []int, m int, nums2 []int, n int)  {
	index := m + n -1
	for m > 0 && n > 0  {
		if nums1[m-1] >= nums2[n-1] {
			nums1[index] = nums1[m-1]
			m--
		} else {
			nums1[index] = nums2[n-1]
			n--
		}
		index--
	}
	if n > 0 {
		nums1[index] = nums2[n-1]
		n--
		index--
	}


}

