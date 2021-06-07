/**
 leetcode 76  Minimum Window Substring

 Given two strings s and t of lengths m and n respectively, return the
 minimum window in s which will contain all the characters in t. If
 there is no such window in s that covers all characters in t, return
 the empty string "".

 Note that If there is such a window, it is guaranteed that there will
 always be only one unique minimum window in s.
 */
package DoublePointers

//滑动窗口的思想：
//用i,j表示滑动窗口的左边界和右边界，通过改变i,j来扩展和收缩滑动窗口，可以想象成一个窗口在字符串上游走
//当这个窗口包含的元素满足条件，即包含字符串T的所有元素，记录下这个滑动窗口的长度j-i+1，这些长度中的
//最小值就是要求的结果。
//步骤一
//不断增加j使滑动窗口增大，直到窗口包含了T的所有元素
//
//步骤二
//不断增加i使滑动窗口缩小，因为是要求最小字串，所以将不必要的元素排除在外，使长度减小，
//直到碰到一个必须包含的元素，这个时候不能再扔了，再扔就不满足条件了，记录此时滑动窗口的长度，并保存最小值
//
//步骤三
//让i再增加一个位置，这个时候滑动窗口肯定不满足条件了，那么继续从步骤一开始执行，
//寻找新的满足条件的滑动窗口，如此反复，直到j超出了字符串S范围。
func minWindowI(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	tmap := make(map[rune]int)
	arrt := []rune(t)
	for _, v := range arrt {
		if _, ok := tmap[v]; !ok {
			tmap[v] = 1
		}
		tmap[v]++
	}
	// Number of unique characters in t, which need to be present in the desired window.
	required := len(tmap)
	left, right := 0, 0

	window := make(map[rune]int)
	// formed is used to keep track of how many unique characters in t
	// are present in the current window in its desired frequency.
	formed := 0

	// ans list of the form (window length, left, right)
	ans := []int{-1, 0, 0};

	arr := []rune(s)

	for right < len(arr) {
		c := arr[right]
		if _, ok := window[c]; !ok {
			window[c] = 1
		}
		window[c]++
		if _,ok := tmap[c]; ok && window[c] == tmap[c] {
			formed++
		}
		for left <= right && formed == required {
			c = arr[left]
			if (ans[0] == -1 || right- left + 1 < ans[0]){
				ans[0] = right - left + 1
				ans[1] = left
				ans[2] = right
			}
			// The character at the position pointed by the
			// `Left` pointer is no longer a part of the window.
			// 开始缩小左窗口
			window[c] = window[c] - 1
			if _,ok := tmap[c]; ok && window[c] < tmap[c] {
				formed--
			}
			// Move the left pointer ahead, this would help to look for a new window.
			left++

		}
		// Keep expanding the window once we are done contracting.
		right++
	}

	if ans[0] == -1 {
		return ""
	}
	return string(arr[ans[1]: ans[2] + 1])

}


