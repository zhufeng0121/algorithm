/**
 leetcode 448. Find All Numbers Disappeared in an Array

 Given an array nums of n integers where nums[i] is in the range
 [1, n], return an array of all the integers in the range [1, n]
 that do not appear in nums.

	Input: nums = [4,3,2,7,8,2,3,1]
	Output: [5,6]
 */

package main
//不开辟额外空间，将nums当成Map
func findDisappearedNumbers(nums []int) []int {
	res := make([]int,0)
	for i := 0; i < len(nums);i++ {
		nums[nums[i] -1] += len(nums)
	}
	for i := 0; i < len(nums);i++ {
		if nums[i] < len(nums) {
			res = append(res, i)
		}
	}
	return res

}
//
func findDisappearedNumbersIV(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		//fmt.Printf("i=%d\n",i)
		for nums[i]-1 != i && nums[i] != nums[nums[i]-1] {
			// fmt.Printf("nums[i]=%d\n",nums[i])
			// fmt.Printf("nums[nums[i]-1]=%d\n",nums[nums[i]-1])
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			//fmt.Println(nums)
		}
	}

	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			ans = append(ans, i+1)
		}
	}

	return ans
}
