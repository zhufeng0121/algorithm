/**
 leetcode 345. Reverse Vowels of a string

 Given a string s, reverse only all the vowels in the string and return it.

 The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both cases.

 */
package main

func reverseVowels(s string) string {
	left, right := 0, len(s) - 1

	tmp := []byte(s)

	vowel := make(map[byte]bool)
	vowel['a'] = true
	vowel['e'] = true
	vowel['i'] = true
	vowel['o'] = true
	vowel['u'] = true
	vowel['A'] = true
	vowel['E'] = true
	vowel['I'] = true
	vowel['O'] = true
	vowel['U'] = true

	for left < right {
		if vowel[tmp[left]] && vowel[tmp[right]] {
			tmp[left], tmp[right] = tmp[right], tmp[left]
			left++
			right--

		} else if _, ok := vowel[tmp[left]]; !ok {
			left++
		} else if _, ok := vowel[tmp[right]]; !ok {
			right--
		}

	}
	return string(tmp)

}
