/**
 leetcode 215 Kth Largest Element in an Array

 Given an integer array nums and an integer k, return the kth largest
 element in the array.

 Note that it is the kth largest element in the sorted order, not the kth
 distinct element.

 */
package main

import (
	"sort"
)

//这是一道非常非常经典的算法题，是有很多种解决方法的
func findKthLargest(nums []int, k int) int {
	sort.Slice(nums,func(i int,j int ) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]

}
//利用堆，维护一个大小为k的最小堆，在遍历的数组中，最终的根节点即第K大的数

//维护一个容量为k的小顶堆，堆中的k个节点代表着当前最大的k个元素，而堆顶显然是这k个元素中的最小值。

//遍历原数组，每遍历一个元素，就和堆顶比较，如果当前元素小于等于堆顶，则继续遍历；如果元素大于堆顶，则把当前元素放在堆顶位置，并调整二叉堆（下沉操作）。
//
//遍历结束后，堆顶就是数组的最大k个元素中的最小值，也就是第k大元素。

func findKthLargestByHeap(nums []int, k int) int {
	buildHeap(nums,k)
	for i := k; i < len(nums);i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			sink(nums,0,k)
		}
	}
	return nums[0]

}
func sink(arr []int, index int, k int) {
	//这里的意义不仅仅在于只是为了交换而用，因为arr[index]会被重新赋值
	temp := arr[index]
	child := index << 1 + 1
	for child < k {
		if child + 1 < k && arr[child] > arr[child + 1] {
			child++
		}
		//这里的比较一定要用temp
		if temp <= arr[child] {
			break
		}
		arr[index] = arr[child]
		index = child
		child = child << 1 + 1
	}
	arr[index] = temp
}


func buildHeap(nums []int, length int) {
	//后半部分属于堆的叶子节点，不需要进行下沉操作
	for i := (length - 2) / 2 ; i >= 0; i-- {
		sink(nums,i,length)
	}

}

//




