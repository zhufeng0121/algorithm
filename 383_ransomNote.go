/**
 leetcode 383. Ransom Note

 Given two stings ransomNote and magazine, return true if ransomNote
 can be constructed from magazine and false otherwise.

 Each letter in magazine can only be used once in ransomNote.

 */
package main

func canConstruct(ransomNote string, magazine string) bool {
	rmap := make(map[byte]int)
	mmap := make(map[byte]int)
	for i := 0;i < len(ransomNote);i++ {
		rmap[ransomNote[i]]++
	}
	for i := 0;i < len(magazine); i++ {
		mmap[magazine[i]]++
	}
	for key, value := range rmap {
		if _, ok := mmap[key]; !ok {
			return false
		}
		if value > mmap[key] {
			return false
		}
	}
	return true

}
