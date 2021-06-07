/**
 leetcode 135  Candy

 There are n children standing in a line. Each child is assigned a rating
 value given in the integer array ratings.

 You are giving candies to these children subjected to the following requirements:

 Each child must have at least one candy.
 Children with a higher rating get more candies than their neighbors.
 Return the minimum number of candies you need to have to distribute the candies to the children.
 */
package main

//但我们只需要简单的两次遍历即可:把所有孩子的糖果数初始化为 1;
//先从左往右遍历一遍,如果右边孩子的评分比左边的高,则右边孩子的糖果数更新为左边孩子的
//糖果数加 1;再从右往左遍历一遍,如果左边孩子的评分比右边的高,且左边孩子当前的糖果数
//不大于右边孩子的糖果数,则左边孩子的糖果数更新为右边孩子的糖果数加 1。
func candy(ratings []int) int {
	childNum := len(ratings)
	candy := make([]int, childNum)
	for i := 0 ;i < childNum;i++ {
		candy[i] = 1
	}
	for i := 0; i < len(ratings) -1;i++ {
		if ratings[i] < ratings[i + 1] {
			candy[i+1] = candy[i] + 1
		}
	}
	for i := len(ratings) - 1; i> 0;i-- {
		if ratings[i] < ratings[i-1] && candy[i] >= candy[i-1] {
			candy[i-1] = candy[i] + 1
		}
	}
	sum := 0
	for i := 0;i < len(candy);i++ {
		sum += candy[i]
	}
	return sum

}
