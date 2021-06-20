/**
 leetcode 371 Sum of Two Integers

 Given two integers a and b, return the sum of the two
 integers without using the operators + and -.


 */
package main

func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a

}

func getSumByBit(a int ,b int) int {
	for {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
		if carry == 0 {
			break
		}
	}
	return a

}


//递归调用
func add(a int32, b int32) int32 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return add(2*(a&b), a^b)
}

