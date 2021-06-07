/**
 leetcode 67 Add Binary

 Given two BinarySearch strings a and b,
 return their sum as a BinarySearch string.

 */

package main

import "strings"

// 思路是在短的数组前面不断补0,使长度和长数组相同，然后进行逐位相加
func addBinary(a string, b string) string {
	//将a字符串设置为其中较长的那一个
	if len(a) < len(b) {
		a, b = b, a
	}
	arr1, arr3 := []byte(a), []byte(b)
	res := make([]byte, len(a))
	//新的b字符串
	arr2 := make([]byte, len(a))
	for i := 0;i < len(a) - len(b);i++ {
		arr2[i] = '0'
	}
	//一定要注意这里的坐标转换
	for i := 0; i < len(b);i++ {
		arr2[i + len(a) - len(b)] = arr3[i]
	}

	var carry byte = 0
	for i := len(a) -1 ;i >= 0;i-- {
		res[i] = ((arr1[i] - '0' + arr2[i] - '0') + carry) % 2 + '0'
		if  (arr1[i] - '0' + arr2[i] - '0' + carry) / 2 == 1 {
			carry = 1
		} else {
			carry = 0
		}

	}
	if carry == 1 {
		return "1" + string(res)
	}
	return string(res)
}


//简洁明了的方法，我的思路和他差不多，但是太繁琐
func addBinaryIV(a string, b string) string {
	l1, l2 := len(a), len(b)
	l := l1
	if l1 > l2 {
		// 这个strings.Repeat真的好用
		b = strings.Repeat("0",l1 - l2) + b
	} else if l1 < l2 {
		a = strings.Repeat("0",l2 - l1) + a
		l = l2
	}
	s := []byte(a)
	var carry byte = 0
	for i := l - 1;i >= 0;i-- {
		s[i] = (a[i] - '0' + b[i] - '0' + carry) %2 + '0'
		carry = (a[i] - '0' + b[i] - '0' + carry) / 2
	}
	if carry == 1 {
		return "1" + string(s)
	}
	return string(s)

}

