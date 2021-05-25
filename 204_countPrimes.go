/**
 leetcode 204 Count Primes

 Count the number of prime numbers less than a non-negative number, n.

 */
package main

/**
 此题要采用素数筛法来进行计算：
 基本思路是： 将n 开根号
 去掉 2 到 根号n 之间 所有素数的倍数，剩下的全部是素数


 只要小于或等于根号N的数（1除外）不能整除N，则N一定是素数；反过来说就是
 只要小于或等于根号N的数（1除外）能整除N，则N一定是合数

 假设一个数N是合数，那一定存在一个因数b和一个非1且非自身的因数a，即a*b=N

 等式两边同时开根号得出：(a*b)^0.5=a^0.5*b^0.5=N^0.5

 可以推出：若N为合数，则a和b中一定有一个大于或等于N^0.5，另一个小于或等于N^0.5

　按照这个结论，就只需计算小于等于N^0.5的数了，这样就大大提高了效率（要注意等于N^0.5的这个边界）：

 */

//
func countPrimes(n int) int {
	result := 0
	// 初始值均为false
	prime := make([]bool,n)
	for i := 2;i < n;i++ {
		if prime[i] == false {
			result++
		}
		for j := 2 * i;j < n;j += i {
			prime[j] = true
		}
	}
	return result
}


func countPrimesII(n int) int {
	if n <= 2 {
		return 0
	}

	notprime := make([]bool, n)
	//去掉1 ，去掉n
	count := n - 2
	for i := 2; i*i <= n; i++ {
		for j := i * i; j < n; j += i {
			if notprime[j] == false {
				notprime[j] = true
				count--
			}
		}
	}

	return count
}
