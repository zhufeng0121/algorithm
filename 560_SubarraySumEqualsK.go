/**
 leetcode 560 Subarray Sum Equals K

 Given an array of integers nums and an integer k, return
 the total number of continuous subarrays whose sum equals to k.

*/
package main

//利用前缀和
func subarraySum(nums []int, k int) int {
	presum := make([]int,len(nums) + 1)
	presum[0] = 0
	for i := 0; i < len(nums); i++ {
		presum[i+1] = presum[i] + nums[i]
	}
	//fmt.Println(presum)
	res := 0
	for i := 0; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if presum[i] - presum[j] == k {
				// fmt.Println(j,i)
				res++
			}
		}
	}
	return res

}


func subarraySumI(nums []int, k int) int {
	presum := make([]int,len(nums))
	for i := 0; i < len(nums); i++ {
		presum[i+1] = nums[i] + presum[i]
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if presum[j+1] - presum[i] == k {
				count++
			}
		}
	}
	return count
}

//参考两数之和，重点就是找到和为k的连续子数组，假设在区间[left, right]中的数组元素的和为k，
//那么利用前缀和中的sum数组，即sum[right] - sum[left] = k，
//然后扫描整个数组，利用map存储 sum[left] 出现的个数，
//加上sum[right]- k作为出现过的前n项和的个数，然后更新统计量即可。

func subarraySumII(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	smap := make(map[int]int)
	//初始化一对键值对0,1
	smap[0] = 1
	count := 0
	presum := 0
	for _, v := range nums {
		presum += v
		if _, ok := smap[presum - k]; ok {
			//看 presum - k 出现过几次
			count += smap[presum - k]
		}
		smap[presum]++
	}
	return count
}