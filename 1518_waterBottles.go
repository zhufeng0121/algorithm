/**
 leetcode 1518. Water Bottles

 Given numBottles full water bottles, you can exchange numExchange empty water
 bottles for one full water bottle.
 The operation of drinking a full water bottle turns it into an empty bottle.
 Return the maximum number of water bottles you can drink.
 */
package main

func numWaterBottles(numBottles int, numExchange int) int {
	drink, empty := numBottles, numBottles
	for empty >= numExchange {
		drink += empty /numExchange
		empty = empty / numExchange + empty % numExchange


	}
	return drink

}
