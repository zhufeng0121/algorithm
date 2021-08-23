/**
 leetcode 20 Valid Parentheses

 Given a string s containing just the characters '(', ')',
 '{', '}', '[' and ']', determine if the input string is
 valid

 */
package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune,0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack,v)
		} else if (v == ')') && len(stack) > 0 && stack[len(stack) - 1] == '(' ||
			(v == ']') && len(stack) > 0 && stack[len(stack) - 1] == '[' ||
			(v == '}') && len(stack) > 0 && stack[len(stack) - 1] == '{' {
			stack = stack[:len(stack) - 1]
		} else {
			//要处理那种没有左和右的情况
			return false
		}

	}
	return len(stack) == 0

}


