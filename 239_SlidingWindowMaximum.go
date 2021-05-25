/**
 leetcode 239 Sliding Window Maximum

 You are given an array of integers nums, there is a sliding window of
 size k which is moving from the very left of the array to the very right.
 You can only see the k numbers in the window. Each time the sliding window
 moves right by one position.

 Return the max sliding window.

 */
package main
// 利用队列来存，单调队列，队列中存储下标


//一定要好好理解这道题
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k < 1 || len(nums) < k {
		return nil
	}
	//单调队列，用于存放下标
	queue := make([]int,0)
	res := make([]int,0)
	for i := 0;i < len(nums);i++ {
		//pop elements from queue tail if the element is smaller than nums[i]
		//so the queue is always in desc order
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[0 : len(queue)-1]
		}
		//如果队列为空，或者队列不为空，但是队列的最后一个元素值要小于等于当前nums[i],将下标入队
		queue = append(queue, i)
		for len(queue) > 0 && queue[0] < i - k + 1 {
			queue = queue[1:]
		}
		//pop element from queue head if the element is going to be moved out of sliding window
		if queue[0] == i - k {
			queue = queue[1:]
		}
		//sliding window size is equal to k
		if i - k + 1 >= 0 {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}


func maxSlidingWindowII(nums []int, k int) []int {
	res := make([]int, 0)

	//the queue is used to keep sliding window indexes
	queue := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		//pop elements from queue tail if the element is smaller than nums[i]
		//so the queue is always in desc order
		for len(queue) > 0 && nums[i] > nums[queue[len(queue) - 1]] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, i)

		//pop element from queue head if the element is going to be moved out of sliding window
		if queue[0] == i - k {
			queue = queue[1:]
		}

		//sliding window size is equal to k
		if i - k + 1 >= 0 {
			res = append(res, nums[queue[0]])
		}
	}


	return res
}

