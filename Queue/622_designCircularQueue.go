/**
 leetcode 622 Design Circular Queue

 Design your implementation of the circular queue. The circular queue is a
 linear data structure in which the operations are performed based on FIFO
 (First In First Out) principle and the last position is connected back to
 the first position to make a circle. It is also called "Ring Buffer".

 One of the benefits of the circular queue is that we can make use of the
 spaces in front of the queue. In a normal queue, once the queue becomes
 full, we cannot insert the next element even if there is a space in front
 of the queue. But using the circular queue, we can use the space to store new values.

 Implementation the MyCircularQueue class:

 MyCircularQueue(k) Initializes the object with the size of the queue to be k.
 int Front() Gets the front item from the queue. If the queue is empty, return -1.
 int Rear() Gets the last item from the queue. If the queue is empty, return -1.
 boolean enQueue(int value) Inserts an element into the circular queue. Return true if the operation is successful.
 boolean deQueue() Deletes an element from the circular queue. Return true if the operation is successful.
 boolean isEmpty() Checks whether the circular queue is empty or not.
 boolean isFull() Checks whether the circular queue is full or not.
 You must solve the problem without using the built-in queue data structure in your programming language.

*/
package Queue

type MyCircularQueue struct {
	size   int
	queue  []int
	front  int
	rear   int

}

func ConstructorI(k int) MyCircularQueue {
	queue := make([]int, k + 1)
	return MyCircularQueue{
		size: k + 1,
		queue: queue,
		front: 0,
		rear : 0,
	}

}


func (this *MyCircularQueue) EnQueue(value int) bool {
	//入队要先判断队列是否满
	if this.IsFull() {
		return false
	} else {
		//队尾入队
		this.queue[this.rear] = value
		this.rear = (this.rear + 1) % this.size
		return true
	}

}


func (this *MyCircularQueue) DeQueue() bool {
	//出队要先判断队列是否为空
	if this.IsEmpty() {
		return false
	} else {
		this.front = (this.front + 1) % this.size
		return true
	}


}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.queue[this.front]
	}

}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.queue[(this.rear - 1 + this.size) % this.size]
	}

}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear

}


func (this *MyCircularQueue) IsFull() bool {
	return (this.rear + 1) % this.size == this.front

}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
