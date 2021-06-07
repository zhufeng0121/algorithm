/**
 leetcode 295 Find Median from Data Stream

 The median is the middle value in an ordered integer list. If the size of
 the list is even, there is no middle value and the median is the mean
 of the two middle values.

 For example, for arr = [2,3,4], the median is 3.
 For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.
 Implement the MedianFinder class:

 MedianFinder() initializes the MedianFinder object.
 void addNum(int num) adds the integer num from the data stream to the data structure.
 double findMedian() returns the median of all elements so far. Answers within 10-5
 of the actual answer will be accepted.

 */
package main
/**
 MedianFinder 中有两个堆,一个是大根堆,另一个是小根堆。大根堆中含有接收的所有数中较小的一半,
 并且按大根堆的方式组织起来,那么这个堆的堆顶就是较小一半的数中最大的那个。小根堆中含有接收的
 所有数中较大的一半,并且按小根堆的方式组织起来,那么这个堆的堆顶就是较大一半的数中最小的那个。
 */
/**
 1)第一个出现的数直接进入大根堆。
 2)以后对每一个新出现的数 cur,判断 cur 是否小于或等于大根堆的堆顶。如果是,cur 进
 入大根堆;如果不是,cur 进入小根堆。
 3)每一个数加入完成后,判断大根堆和小根堆的大小。如果个数较多的堆比个数较少的堆
 拥有数的个数超过 1,从个数较多的堆中弹出堆顶,放入另一个堆。
 4)任何时候想得到所有数字的中位数,一定可以由两个堆的堆顶得到。
 */

type MedianFinder struct {


}


/** initialize your data structure here. */
func Constructor() MedianFinder {

}


func (this *MedianFinder) AddNum(num int)  {

}


func (this *MedianFinder) FindMedian() float64 {

}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
