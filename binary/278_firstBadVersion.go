/**
 leetcode 278 First Bad Version

 You are a product manager and currently leading a team to develop a new product.
 Unfortunately, the latest version of your product fails the quality check. Since
 each version is developed based on the previous version, all the versions after a
 bad version are also bad.

 Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad
 one, which causes all the following ones to be bad.

 You are given an API bool isBadVersion(version) which returns whether version is bad.
 Implement a function to find the first bad version. You should minimize the number of
 calls to the API.

 */
package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// 二分查找
func firstBadVersion(n int) int {
	start, end := 1, n
	for start < end {
		mid := start + (end -start) >> 1
		if isBadVersion(mid) == false {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start

}
func isBadVersion(n int) bool {
	return false
}
