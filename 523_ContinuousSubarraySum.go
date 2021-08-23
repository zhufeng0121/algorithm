/**
 leetcode 523 . Continuous Subarray Sum

 Given an integer array nums and an integer k, return true if nums has a
 continuous subarray of size at least two whose elements sum up to a multiple
 of k, or false otherwise.

 An integer x is a multiple of k if there exists an integer n such that
 x = n * k. 0 is always a multiple of k.
 */
package main

//先求前缀和数组

//利用同余的概念

//核心：当两个不同位置的前缀和对 kk 的取余相同时，
//我们看这两个位置的下标是否距离大于等于2. 如果满足以上条件，
//我们即找到了一个连续数组的和是 kk 的倍数。

func checkSubarraySum(nums []int, k int) bool {
	//建立map,key 存储当前前缀和的余数，value为下标
	m := make(map[int]int)
	//如果某个前缀和除k后余0，从map中取到的就是这个（0，-1），当前index+1大于等于2即直接返回true
	//这个就类似于哨兵，类似链表题在head前面加虚拟节点，可以不用处理边界条件
	//[3,6,1], k =3
	m[0] = -1
	presum := 0
	for i := 0; i < len(nums); i++ {
		presum += nums[i]
		remainder := presum % k
		if _, ok := m[remainder]; ok {
			if i - m[remainder] >= 2 {
				return true
			}
		} else {
			//这里只有不存在的时候才更新，存在的话，不要更新
			m[remainder] = i

		}

	}
	return false

}
