/**
 leetcode 541 Reverse String II

 Given a string s and an integer k, reverse the first k
 characters for every 2k characters counting from the
 start of the string.

 If there are fewer than k characters left, reverse all of
 them. If there are less than 2k but greater than or equal
 to k characters, then reverse the first k characters and
 left the other as original.


 */
package main
//每隔2k个字符，对前k个字符进行旋转， 如果不满2k个，做法太麻烦了
func reverseStr(s string, k int) string {
	var chunks []string
	arr := []byte(s)
	if len(arr) == 0 {
		return ""
	}
	for i := 0; i < len(arr); i += 2 * k {
		index := i + 2 * k
		if index > len(arr) {
			index = len(arr)
		}
		chunks = append(chunks, string(arr[i:index]))
	}
	var result string
	if len(chunks) > 1 {
		for i := 0; i < len(chunks) -1;i++ {
			result += reverseAll(chunks[i], k)
		}
	}
	last := len(chunks[len(chunks) - 1])
	if last < k {
		result += reverseAll(chunks[len(chunks) -1 ],last)
	} else if last <= 2 * k{
		result += reverseAll(chunks[len(chunks) -1],k)
	}
	return result

}

func reverseAll(s string,k int) string {
	arr := []byte(s)
	for i := 0; i < k/2 ;i++ {
		arr[i], arr[k - i - 1] = arr[k - i - 1], arr[i]
	}
	return string(arr)
}
//没看懂
func reverseStrI(s string,k int) string {
	arr := []byte(s)
	for start := 0; start < len(arr); start += 2 * k {
		i := start
		j := min(start + k -1,len(arr) -1)
		for i < j {
			arr[i] , arr[j] = arr[j], arr[i]
			i++
			j--
		}

	}
	return string(arr)

}
