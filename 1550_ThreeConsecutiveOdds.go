/**
 leetcode 1550 Three Consecutive Odds

 Given an integer array arr, return true if there are three
 consecutive odd numbers in the array. Otherwise, return false.

 */
package main

func threeConsecutiveOdds(arr []int) bool {
	count, index := 0, 0
	for index < len(arr) {
		if arr[index] & 1 == 1 {
			count++
			index++
		} else {
			count = 0
			index ++
		}
		if count >= 3 {
			return true
		}

	}
	return false
}
