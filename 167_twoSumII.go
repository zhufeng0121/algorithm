/**
 leetcode 167 Two Sum II - Input array is sorted

 Given an array of integers numbers that is already sorted in
 ascending order, find two numbers such that they add up to a
 specific target number.

 Return the indices of the two numbers (1-indexed) as an integer
 array answer of size 2, where 1 <= answer[0] < answer[1] <= numbers.length.

 You may assume that each input would have exactly one solution
 and you may not use the same element twice.

 */
package main

//可以证明,对于排好序且有解的数组,双指针一定能遍历到最优解。证明方法如下:假设最
//优解的两个数的位置分别是 l 和 r。我们假设在左指针在 l 左边的时候,右指针已经移动到了 r;
//此时两个指针指向值的和小于给定值,因此左指针会一直右移直到到达 l。同理,如果我们假设
//在右指针在 r 右边的时候,左指针已经移动到了 l;此时两个指针指向值的和大于给定值,因此
//右指针会一直左移直到到达 r。所以双指针在任何时候都不可能处于 (l, r) 之间,又因为不满足条
//件时指针必须移动一个,所以最终一定会收敛在 l 和 r。

//本题应该还有很多种解法
func twoSumII(numbers []int, target int) []int {

	start := 0
	end := len(numbers) -1

	for start < end {
		if numbers[start] + numbers[end] == target {
			return []int{start + 1,end + 1}
		} else if numbers[start] + numbers[end] > target {
			end--
		} else {
			start++
		}
	}
	return []int{0,0}

}


