/**
 leetcode 134 Gas Station

 There are n gas stations along a circular route, where the amount
 of gas at the ith station is gas[i].

 You have a car with an unlimited gas tank and it costs cost[i] of
 gas to travel from the ith station to its next (i + 1)th station.
 You begin the journey with an empty tank at one of the gas stations.

 Given two integer arrays gas and cost, return the starting gas station's index
 if you can travel around the circuit once in the clockwise direction,
 otherwise return -1. If there exists a solution, it is guaranteed to be unique

在一条环路上N个加油站，其中第i个加油站有汽油gas[i]升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1个加油站需要消耗汽油cost[i]升。
你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

*/
package main

//
func canCompleteCircuit(gas []int, cost []int) int {
	gasSum := sum(gas)
	costSum := sum(cost)
	if costSum > gasSum {
		return -1
	}


}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}