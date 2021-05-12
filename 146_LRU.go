/**
 leetcode 146 LRU Cache

 Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

 Implement the LRUCache class:
 LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
 int get(int key) Return the value of the key if the key exists, otherwise return -1.

 void put(int key, int value) Update the value of the key if the key exists.
 Otherwise, add the key-value pair to the cache. If the number of keys exceeds
 the capacity from this operation, evict the least recently used key.

 Follow up:
 Could you do get and put in O(1) time complexity?

 */
package main

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

//没写完，这个太复杂了
//大致思路和设计思想总结如下：
//用双向链表

type Node struct {
	 pre     *Node
	 post    *Node
	 key     int
	 val     int
}
type DoubleList struct {
	head     *Node
	tail     *Node
}
// Push element to front of the list
func (this *DoubleList) Push(newTail *Node) {
	if this.head == nil {
		this.head = newTail
		this.tail = newTail
	} else {
		temp := this.tail
		this.tail = newTail
		temp.post = newTail
		this.tail.pre = temp

	}

}

// Pop element pointed by the head
func (this *DoubleList) Pop() {
	//仅有一个节点
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.post
		this.head.pre.post = nil
		this.head.pre = nil

	}

}
// Splice a node from the list and set prev/next pointers to nil
func (this *DoubleList) Splice(node *Node) {
	//头节点，弹出,
	if node == this.head {
		this.Pop()
	} else if node == this.tail {
		this.tail = this.tail.pre
		this.tail.post = nil
		this.tail.pre = node.pre.pre

	} else {
		node.pre.post = node.post
		node.post.pre = node.pre
		node.pre = nil
		node.post = nil
	}


}

type LRUCache struct {
	Len     int
	Cap     int
	Hash    map[int]*Node
	LRUList DoubleList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Len : 0,
		Cap: capacity,
		Hash: map[int]*Node{},
		LRUList: DoubleList{
			head: nil,
			tail: nil,
		},
	}

}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Hash[key]; !ok {
		return -1
	} else {
		node := this.Hash[key]
		this.LRUList.Splice(node)
		this.LRUList.Push(node)
		return node.val

	}

}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Hash[key]; !ok {
		if this.Len == this.Cap {
			delKey := this.LRUList.head.key
			delete(this.Hash,delKey)
			this.LRUList.Pop()
		}
		newHead := Node{
			pre: nil,
			post: nil,
			key :key,
			val: value,
		}
		this.LRUList.Push(&newHead)
		this.Hash[key] = this.LRUList.head
		//如果this.len == this.Cap len不用加也不用减
		if this.Len < this.Cap {
			this.Len += 1
		}
	} else {
		if this.LRUList.head.key == key {
			this.LRUList.head.val = value
		} else {
			nodePtr := this.Hash[key]
			nodePtr.val = value
			this.LRUList.Splice(nodePtr)
			this.LRUList.Push(nodePtr)
		}
	}

}


