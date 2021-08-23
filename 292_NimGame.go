/**
 leetcode 292 Nim Game

 You are playing the following Nim Game with your friend:

 Initially, there is a heap of stones on the table.
 You and your friend will alternate taking turns, and you go first.
 On each turn, the person whose turn it is will remove 1 to 3 stones
 from the heap.
 The one who removes the last stone is the winner.
 Given n, the number of stones in the heap, return true if you can
 win the game assuming both you and your friend play optimally,
 otherwise return false.

 */
package main

//该题属于巴什博弈的一种，如果每个人只允许每次那最多m颗石子，最后取光者为胜，那么n %（m + 1）
func canWinNim(n int) bool {
	return n % 4 != 0

}

//此题还有一种变形，就是如果最后取光者为败
func canWinNimI(n int, m int ) bool {
	if n % (m + 1) == 1 {
		return false
	}
	return true

}