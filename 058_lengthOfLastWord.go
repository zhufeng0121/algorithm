/**
 leetcode 58 Length of Last Word
 */
package main

import "strings"

func lengthOfLastWord(s string) int {
	//这里用了Trim 不知道合理不合理，不过题过了
	s = strings.Trim(s," ")
	strs := strings.Split(s," ")
	if len(strs) == 0 {
		return 0
	}
	if strs[len(strs) - 1] == "" {
		return 0
	} else  {
		return len(strs[len(strs) - 1])
	}

}


