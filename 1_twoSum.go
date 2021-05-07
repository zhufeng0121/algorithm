
package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	length := len(nums)
	for i := 0;i < length;i++ {
		another := target - nums[i]
		if _ ,ok := m[another];ok {
			return []int{i,m[another]}
		}
		m[nums[i]] = i

	}
	return nil

}