/**
 leetcode 492 Construct the Rectangle

 A web developer needs to know how to design a web page's size. So,
 given a specific rectangular web page’s area, your job by now is to
 design a rectangular web page, whose length L and width W satisfy the
 following requirements:

 The area of the rectangular web page you designed must equal to the
 given target area.
 The width W should not be larger than the length L, which means L >= W.
 The difference between length L and width W should be as small as possible.
 Return an array [L, W] where L and W are the length and width of the web
 page you designed in sequence.

 */
package main

import "math"

func constructRectangle(area int) []int {
	gap := area - 1
	length, width := area, 1
	for i := 2; i < int(math.Sqrt(float64(area))) + 1; i++ {
		if area % i == 0 && gap > area/i - i {
			gap = area / i - i
			length = area / i
			width  = i
		}
	}
	return []int{length,width}

}
func constructRectangleIV(area int) []int {
	for l := int(math.Sqrt(float64(area))); l <= area; l++ {
		if area % l == 0 && l >= area/l {
			return []int{l, area/l}
		}
	}
	return []int{0,0}
}
