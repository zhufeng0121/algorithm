/**
 leetcode 509 Fibonacci Number

 The Fibonacci numbers, commonly denoted F(n) form a sequence,
 called the Fibonacci sequence, such that each number is the
 sum of the two preceding ones, starting from 0 and 1. That is,

	F(0) = 0, F(1) = 1
	F(n) = F(n - 1) + F(n - 2), for n > 1.
	Given n, calculate F(n).


*/
package main

import "math"

//最简单的递归写法
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}
//更快的迭代写法
func fibI(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	arr := make([]int,n+1)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]

}
//利用差分方程，这个挺有意思的
func fibII(N int) int {
	var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2);
	return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)));
}
//不开辟数组迭代方法的
func fibIV(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	cur := 0
	a, b := 1, 1

	for i := 3; i <= n;i++ {
		cur = a + b
		b = a
		a = cur

	}
	return cur

}
//还有一种矩阵的写法，有时间再写
func fibIII(N int) int {
	if N <= 1 {
		return N
	}
	var A = [2][2]int{
		{1,1},
		{1,0},
	}
	A = matrixPower(A, N-1)
	return A[0][0]
}

func matrixPower(A [2][2]int, N int) [2][2]int {
	if N <= 1 {
		return A
	}
	A = matrixPower(A, N/2)
	A = multiply(A, A)

	var B = [2][2]int{
		{1,1},
		{1,0},
	}
	if N%2 != 0 {
		A = multiply(A, B)
	}

	return A
}
//矩阵相乘函数
func multiply(A [2][2]int, B [2][2]int) [2][2]int {
	x := A[0][0] * B[0][0] + A[0][1] * B[1][0];
	y := A[0][0] * B[0][1] + A[0][1] * B[1][1];
	z := A[1][0] * B[0][0] + A[1][1] * B[1][0];
	w := A[1][0] * B[0][1] + A[1][1] * B[1][1];

	A[0][0] = x;
	A[0][1] = y;
	A[1][0] = z;
	A[1][1] = w;

	return A
}
