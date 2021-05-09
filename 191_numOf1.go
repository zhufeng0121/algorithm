/**
 leetcode 191 Number of 1 bits

 Write a function that takes an unsigned integer and returns
 the number of '1' bits it has (also known as the Hamming weight).

 */

package main

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num & 1 == 1 {
			count++
		}
		num >>= 1
	}
	return count

}
