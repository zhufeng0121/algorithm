package main

func main() {
	m := []int{4,5,6,0,0,0}
	n := []int{1,2,3}
	mergeI(m,3,n,3)

}
func mergeI(nums1 []int, m int, nums2 []int, n int)  {
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
	for n > 0 {
		nums1[index] = nums2[n-1]
		n--
		index--
	}

}
