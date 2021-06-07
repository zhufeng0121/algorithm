/**
 leetcode 88 Merge Sorted Array

 Given two sorted integer arrays nums1 and nums2, merge nums2 into
 nums1 as one sorted array.

 The number of elements initialized in nums1 and nums2 are m and n
 respectively. You may assume that nums1 has a size equal to m + n
 such that it has enough space to hold additional elements from nums2.

 */
package main


//因为这两个数组已经排好序,我们可以把两个指针分别放在两个数组的末尾,即 nums1 的
//m − 1 位和 nums2 的 n − 1 位。每次将较大的那个数字复制到 nums1 的后边,然后向前移动一位。
//因为我们也要定位 nums1 的末尾,所以我们还需要第三个指针,以便复制。
//在以下的代码里,我们直接利用 m 和 n 当作两个数组的指针,再额外创立一个 pos 指针,起
//始位置为 m + n − 1。每次向前移动 m 或 n 的时候,也要向前移动 pos。这里需要注意,如果 nums1
//的数字已经复制完,不要忘记把 nums2 的数字继续复制;如果 nums2 的数字已经复制完,剩余
//nums1 的数字不需要改变,因为它们已经被排好序。
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

