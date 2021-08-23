/**
 leetcode 718. Maximum Length of Repeated Subarray

 Given two integer arrays nums1 and nums2, return the maximum
 length of a subarray that appears in both arrays.

*/
package main

//若当前两个元素值相同，即 A[i] == B[j]，则说明当前元素可以构成公共子数组，
//所以还要加上它们的前一个元素构成的最长公共子数组的长度(在原来的基础上加 1)，
//此时状态转移方程：dp[i][j] = dp[i - 1][j - 1] + 1。
//若当前两个元素值不同，即 A[i] != B[j]，则说明当前元素无法构成公共子数组
//(就是：当前元素不能成为公共子数组里的一员)。因为公共子数组必须是连续的，
//而此时的元素值不同，相当于直接断开了，此时状态转移方程：dp[i][j] = 0。
func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
    //dp[i][j] 以 nums1[i]和nums2[j]为结尾的两个数组中公共的、长度最长的子数组的长度
	dp := make([][]int,n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	max := 0
	for i := 1 ; i < n; i++ {
		for j := 1; j < m ;j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	return max

}

//dp[i][j] ：长度为i，末尾项为A[i-1]的子数组，
//与长度为j，末尾项为B[j-1]的子数组，二者的最大公共后缀子数组长度。
//如果 A[i-1] != B[j-1]， 有 dp[i][j] = 0
//如果 A[i-1] == B[j-1] ， 有 dp[i][j] = dp[i-1][j-1] + 1
//base case：如果i==0||j==0，则二者没有公共部分，dp[i][j]=0
//最长公共子数组以哪一项为末尾项都有可能，求出每个 dp[i][j]，找出最大值。

func findLengthI(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	//dp[i][j] 以 nums1[i]和nums2[j]为结尾的两个数组中公共的、长度最长的子数组的长度
	dp := make([][]int,n + 1)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m + 1)
	}
	max := 0
	for i := 1 ; i <= n; i++ {
		for j := 1; j <= m ;j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}

	return max

}
