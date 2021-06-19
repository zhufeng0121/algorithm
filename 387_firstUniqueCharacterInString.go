/**
 leetcode 387.First Unique Character in a String

 Given a string s, return the first non-repeating character in it and return its index.
 If it does not exist, return -1.
 */
package main

//先写常规做法吧
func firstUniqChar(s string) int {
	countMap := make(map[byte]int)
	for i := 0 ;i < len(s);i++ {
		countMap[s[i]]++
	}
	for i := 0 ;i < len(s);i++ {
		if countMap[s[i]] == 1 {
			return i
		}
	}
	return -1

}

func firstUniqCharI(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
//我们也可以借助队列找到第一个不重复的字符。队列具有「先进先出」的性质，
//因此很适合用来找出第一个满足某个条件的元素。
//
//具体地，我们使用与方法二相同的哈希映射，并且使用一个额外的队列，
//按照顺序存储每一个字符以及它们第一次出现的位置。当我们对字符串进行遍历时，
//设当前遍历到的字符为 cc，如果 cc 不在哈希映射中，我们就将 cc 与它的索引
//作为一个二元组放入队尾，否则我们就需要检查队列中的元素是否都满足「只出现一次」
//的要求，即我们不断地根据哈希映射中存储的值（是否为 -1−1）选择弹出队首的元素，
//直到队首元素「真的」只出现了一次或者队列为空。
//
//在遍历完成后，如果队列为空，说明没有不重复的字符，返回 -1−1，否则队首的元素
//即为第一个不重复的字符以及其索引的二元组。

//在维护队列时，我们使用了「延迟删除」这一技巧。也就是说，即使队列中有一些字符出现了
//超过一次，但它只要不位于队首，那么就不会对答案造成影响，我们也就可以不用去删除它。
//只有当它前面的所有字符被移出队列，它成为队首时，我们才需要将它移除。


type pair struct {
	ch  byte
	pos int
}

//利用队列
func firstUniqCharIV(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	q := []pair{}
	for i := range s {
		ch := s[i] - 'a'
		if pos[ch] == n {
			pos[ch] = i
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = n + 1
			for len(q) > 0 && pos[q[0].ch] == n+1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return q[0].pos
	}
	return -1
}


