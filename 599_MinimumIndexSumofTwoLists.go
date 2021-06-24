/**
 leetcode 599 Minimum Index Sum of Two Lists

 Suppose Andy and Doris want to choose a restaurant for dinner,
 and they both have a list of favorite restaurants represented by strings.

 You need to help them find out their common interest with the least
 list index sum. If there is a choice tie between answers, output
 all of them with no order requirement. You could assume there always
 exists an answer.
 */
package main


//此题解法麻烦，考虑优化
func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) < len(list2) {
		list1, list2 = list2, list1
	}
	res := make([]string,0)
	min := len(list1) + len(list2) - 2
	for i := 0; i < len(list1);i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				if i + j < min {
					min = i + j
					res = res[:0]
					res = append(res,list1[i])
				} else if i + j == min {
					res = append(res,list1[i])
				}
			}
		}
	}
	return res

}
