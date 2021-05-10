/**
 leetcode 1492 The kth Factor of n

 Given two positive integers n and k.

 A factor of an integer n is defined as an integer i where n % i == 0.

 Consider a list of all factors of n sorted in ascending order, return
 the kth factor in this list or return -1 if n has less than k factors.


 */
package main

import (
	"math"
	"sort"
)

// 注意要对完全平方数 做一下边界的处理
func kthFactor(n int, k int) int {
	factors := make([]int,0)
	factors = append(factors,1)
	for i := 2 ;i < int(math.Sqrt(float64(n))) + 1; i++ {
		if n % i == 0 {
			if i != n/i {
				factors = append(factors,i)
				factors = append(factors, n/i)
			} else {
				factors = append(factors,i)
			}

		}
	}
	factors = append(factors,n)
	sort.Ints(factors)
	if k > len(factors) {
		return -1
	}
	return factors[k-1]

}


