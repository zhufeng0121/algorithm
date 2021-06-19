/**
 leetcode 374 Guess Number Higher or Lower

 We are playing the Guess Game. The game is as follows:

 I pick a number from 1 to n. You have to guess
 which number I picked.

 Every time you guess wrong, I will tell you whether the number
 I picked is higher or lower than your guess.

 You call a pre-defined API int guess(int num), which returns 3 possible results:

 -1: The number I picked is lower than your guess (i.e. pick < num).
 1: The number I picked is higher than your guess (i.e. pick > num).
 0: The number I picked is equal to your guess (i.e. pick == num).
 Return the number that I picked.
 */
package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	start, end  := 1, n
	for start <= end {
		mid := start + (end - start) >> 1
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			end = mid - 1
		} else if guess(mid) == -1 {
			start = mid + 1
		}
	}
	return -1

}
func guess(mid int) int {
	return 0
}
