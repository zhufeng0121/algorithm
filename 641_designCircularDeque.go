/**
leetcode 641 Design Circular Deque

Design your implementation of the circular double-ended queue (deque).

Your implementation should support following operations:

MyCircularDeque(k): Constructor, set the size of the deque to be k.
insertFront(): Adds an item at the front of Deque. Return true if the operation is successful.
insertLast(): Adds an item at the rear of Deque. Return true if the operation is successful.
deleteFront(): Deletes an item from the front of Deque. Return true if the operation is successful.
deleteLast(): Deletes an item from the rear of Deque. Return true if the operation is successful.
getFront(): Gets the front item from the Deque. If the deque is empty, return -1.
getRear(): Gets the last item from Deque. If the deque is empty, return -1.
isEmpty(): Checks whether Deque is empty or not.
isFull(): Checks whether Deque is full or not.

*/
package main
// 设计一个双端循环队列
// 这个队列看起来设计很简单，但是确实是有很多坑在里面的
// 首先就是 front == rear  会出现表示队空和队满的两种情况，出现二义性
// 因此需要将rear表示为队尾的下一个元素，这样，需要牺牲一个存储单元
// 因此申请空间时，需要申请 K + 1的 数组空间
// 数组rear 指向的位置是空的，为了能够让队列容纳k个元素，需要 申请数组的空间要多一个
type MyCircularDeque struct {
	size int
	queue []int
	front int
	rear  int

}


/** Initialize your data structure here. Set the size of the deque to be k. */
func ConstructorI(k int) MyCircularDeque {
	queue := make([]int,k + 1)
	return MyCircularDeque{
		size: k + 1,
		queue: queue,
		front: 0,
		rear: 0,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	} else {
		//时刻注意循环
		this.front = (this.front - 1 + this.size) % this.size
		this.queue[this.front] = value
		return true
	}

}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	// 判断队列是否已满，如果未满 rear = rear + 1 % k
	if this.IsFull() {
		return false
	} else {
		this.queue[this.rear] = value
		this.rear = (this.rear + 1) % this.size

		return true
	}

}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	//判断队列是否为空
	if this.IsEmpty() {
		return false
	} else {
		this.front = (this.front + 1 ) % this.size
		return true
	}

}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.rear = (this.rear - 1 + this.size) % this.size
		return true
	}

}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	//不要忘记判空
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.front]

}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	//不要忘记判空
	if this.IsEmpty() {
		return -1
	}
	return this.queue[(this.rear -1 + this.size) % this.size]

}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.rear

}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.rear + 1) % this.size  == this.front

}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

