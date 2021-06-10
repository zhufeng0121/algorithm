/**
 leetcode 125 Valid palindrome

 Given a string s, determine if it is a palindrome, considering only
 alphanumeric characters and ignoring cases.

 */
package main

//double pointers
func isPalindrome(s string) bool {
	i, j := 0, len(s) -1
	//
	for i < j {
		//如果s[i]既不是大写字母，也不是小写字母
		for !('a' <= s[i] && s[i] <= 'z') && !('A' <= s[i] && s[i] <= 'Z') {
			i++
		}
		fmt.Println(s[i])
		for !(('a' <= s[j] && s[j] <= 'z') || ('A' <= s[j] && s[j] <= 'Z')) {
			j--
		}
		fmt.Println(s[j])
		if s[i] != s[j] && s[i] != s[j] - 32 && s[i] != s[j] +32 {
			return false
		}
		i++
		j--
	}
	return true

}
