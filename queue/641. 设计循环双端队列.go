package main

// 设计实现双端队列。
//你的实现需要支持以下操作：
//
//MyCircularDeque(k)：构造函数,双端队列的大小为k。
//insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
//insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
//deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
//deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
//getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
//getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
//isEmpty()：检查双端队列是否为空。
//isFull()：检查双端队列是否满了。
//示例：
//
//MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
//circularDeque.insertLast(1);			        // 返回 true
//circularDeque.insertLast(2);			        // 返回 true
//circularDeque.insertFront(3);			        // 返回 true
//circularDeque.insertFront(4);			        // 已经满了，返回 false
//circularDeque.getRear();  				// 返回 2
//circularDeque.isFull();				        // 返回 true
//circularDeque.deleteLast();			        // 返回 true
//circularDeque.insertFront(4);			        // 返回 true
//circularDeque.getFront();				// 返回 4
// 
// 
//
//提示：
//
//所有值的范围为 [1, 1000]
//操作次数的范围为 [1, 1000]
//


// 分析，双端队列 就是两头都可以进出，非单一方向的入队和出对。 可以基于数组和链表进行构建双端队列


// 方法一：  双端队列 数组版
type MyCircularDeque struct {
	Val  []int
	Head  int
	Tail  int
	Size   int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	value := []int{}
	this := MyCircularDeque{Val:value,
		Head:-1,
		Tail:-1,
		Size:k,
	}
	return this
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.Val = append(this.Val,value)
		this.Head = 0
		this.Tail = 0
		return true
	}
	kk := []int{}
	kk = append(kk,value)
	kk = append(kk,this.Val[:]...)
	this.Val = kk
	this.Tail++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.Val = append(this.Val,value)
		this.Head = 0
		this.Tail = 0
		return true
	}
	this.Val = append(this.Val,value)
	this.Tail++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Tail == this.Head {
		this.Tail = -1
		this.Head = -1
		this.Val = []int{}
		return true
	}
	if this.Tail > this.Head {
		this.Val = this.Val[this.Head+1:this.Tail+1]
	}
	this.Tail -= 1
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Tail == this.Head {
		this.Tail = -1
		this.Head = -1
		this.Val = []int{}
		return true
	}
	this.Val = this.Val[this.Head:this.Tail]
	this.Tail -= 1
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Val[this.Head]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Val[this.Tail]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if  this.Head == -1 && this.Tail == -1  {
		return true
	}
	return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if  this.Tail - this.Head +1  == this.Size  {
		return true
	}
	return false
}


// 方法二： 双端队列  链表版  使用链表会更简单一些


type LinkNode struct {
	Val int
	Next *LinkNode
	Pre *LinkNode
}
type MyCircularDeque struct {
	Size   int
	Head  *LinkNode
	Tail  *LinkNode
	Length  int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	pre := &LinkNode{}
	next := &LinkNode{}
	pre.Next = next
	pre.Pre = nil
	next.Next = nil
	next.Pre = pre
	this := MyCircularDeque{
		Size: k,
		Head: pre,
		Tail: next,
		Length: 0 ,
	}
	return this
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		new_node := &LinkNode{Val:value}
		this.Head.Next = new_node
		new_node.Pre = this.Head
		this.Tail.Pre = new_node
		new_node.Next = this.Tail
		this.Length++
		return true
	}

	node := &LinkNode{Val:value}
	next := this.Head.Next
	this.Head.Next = node
	node.Next = next
	next.Pre = node
	node.Pre =this.Head
	this.Length++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		new_node := &LinkNode{Val:value}
		this.Head.Next = new_node
		new_node.Pre = this.Head
		this.Tail.Pre = new_node
		new_node.Next = this.Tail
		this.Length++
		return true
	}
	node := &LinkNode{Val:value}
	next := this.Tail.Pre
	next.Next = node
	node.Next = this.Tail
	node.Pre = next
	this.Tail.Pre =node
	this.Length++
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	next := this.Head.Next.Next
	this.Head.Next = next
	next.Pre = this.Head
	this.Length -= 1
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	pre := this.Tail.Pre.Pre
	pre.Next = this.Tail
	this.Tail.Pre = pre
	this.Length -= 1
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Head.Next.Val
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Tail.Pre.Val
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if  this.Head.Next == this.Tail {
		return true
	}
	return false
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.Size == this.Length {
		return true
	}
	return false
}
