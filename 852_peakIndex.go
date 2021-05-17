/**
 leetcode 852 Peak Index in a Mountain Array

 Let's call an array arr a mountain if the following properties hold:

   arr.length >= 3
   There exists some i with 0 < i < arr.length - 1 such that:
   arr[0] < arr[1] < ... arr[i-1] < arr[i]
   arr[i] > arr[i+1] > ... > arr[arr.length - 1]
 Given an integer array arr that is guaranteed to be a mountain, return any
 i such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... >
 arr[arr.length - 1].
 */
package main

func peakIndexInMountainArray(arr []int) int {
	start, end := 0, len(arr) - 1
	if len(arr) < 3 {
		return -1
	}
	for start <= end {
		mid := start + (end - start) >> 1
		// 注意条件不要写错，这个题太简单了
		if (arr[mid] > arr[mid -1]) && (arr[mid] > arr[mid + 1]) {
			return mid
		} else if arr[mid] > arr[mid + 1] && arr[mid] < arr[mid -1]{
			start = mid + 1
		} else if arr[mid] < arr[mid -1] && arr[mid] > arr[mid + 1] {
			end = mid - 1
		}
	}
	return -1

}
