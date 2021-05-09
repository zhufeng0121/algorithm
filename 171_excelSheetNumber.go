/**
 leetcode 171 Excel Sheet Column Number

 Given a string columnTitle that represents the column title as
 appear in an Excel sheet, return its corresponding column number.

 */
package main

var letter  = map[rune]int {
	'A' : 1,
	'B' : 2,
	'C' : 3,
	'D' : 4,
	'E' : 5,
	'F' : 6,
	'G' : 7,
	'H' : 8,
	'I' : 9,
	'J' : 10,
	'K' : 11,
	'L' : 12,
	'M' : 13,
	'N' : 14,
	'O' : 15,
	'P' : 16,
	'Q' : 17,
	'R' : 18,
	'S' : 19,
	'T' : 20,
	'U' : 21,
	'V' : 22,
	'W' : 23,
	'X' : 24,
	'Y' : 25,
	'Z' : 26,

}
// 注意 秦九韶算法的 程序写法
func titleToNumber(columnTitle string) int {
	strs := []rune(columnTitle)
	sum := letter[strs[0]]
	for i := 1;i < len(strs) ;i++ {
		sum = sum * 26 + letter[strs[i]]
	}
	return sum

}

// 通过下标访问字符串时，返回的是一个字节.所以，如果你一定要通过下标访问字符串，可以先将其转换为[]rune类型
func titleToNumberII(s string) int {
	val, res := 0, 0
	for i := 0 ;i < len(s);i++ {
		val = int(s[i] - 'A' + 1)
		res = res * 26 + val
	}
	return res
}

