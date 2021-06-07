/**
 leetcode 45 Jump Game II

 Given an array of non-negative integers nums, you are initially positioned
 at the first index of the array.

 Each element in the array represents your maximum jump length at that position.

 Your goal is to reach the last index in the minimum number of jumps.

 You can assume that you can always reach the last index.

 */
package main

//Greedy 返回到达末尾的最小跳跃次数
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	jump := 0
	cur := 0
	next := 0
	for i := 0;i < len(nums);i++ {
		if cur < i {
			jump++
			cur = next
		}
		next = max(next, i + nums[i])
	}
	return jump

}

//如果我们贪心地进行正向查找，每次找到可到达的最远位置，就可以在线性时间内得到最少的跳跃次数。

//例如，对于数组 [2,3,1,2,4,2,3]，初始位置是下标 0，从下标 0 出发，最远可到达下标 2。
//下标 0 可到达的位置中，下标 1 的值是 3，从下标 1 出发可以达到更远的位置，因此第一步到达下标 1。

//从下标 1 出发，最远可到达下标 4。下标 1 可到达的位置中，下标 4 的值是 4 ，从下标 4 出发可以达
//到更远的位置，因此第二步到达下标 4。
//在遍历数组时，我们不访问最后一个元素，这是因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置，
//否则就无法跳到最后一个位置了。如果访问最后一个元素，在边界正好为最后一个位置的情况下，我们会增加一次
//「不必要的跳跃次数」，因此我们不必访问最后一个元素。


func jumpI(nums [] int) int {
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < len(nums) - 1; i++ {
		maxPosition = max(maxPosition,i + nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps



}

