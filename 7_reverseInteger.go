package main

func reverseInt(x int) int{
	result := 0
	for x != 0 {
		result = result * 10 + x%10
		x = x/10
	}
	if result > 1 <<31 -1 || result < -(1<<31 -1) {
		return 0
	}
	return result
}
