package main

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) —— 将元素 x 推入栈中。
//pop() —— 删除栈顶的元素。
//top() —— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。
// 
//
//示例:
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
// 
//
//提示：
//
//pop、top 和 getMin 操作总是在 非空栈 上调用。


// 思考方式 使用辅助栈来达到O(1) 获取GetMin,栈的实现是用切片实现的,
// 当栈用链表实现的时候，就可以增加一个Min的值
// 看到有使用一个栈来实现求最小值的，思路挺新奇的
// 对于栈来说，如果一个元素 a 在入栈时，栈里有其它的元素 b, c, d，那么无论这个栈在之后经历了什么操作，只要 a 在栈中，b, c, d 就一定在栈中，因为在 a 被弹出之前，b, c, d 不会被弹出。
//
//因此，在操作过程中的任意一个时刻，只要栈顶的元素是 a，那么我们就可以确定栈里面现在的元素一定是 a, b, c, d。
//
//那么，我们可以在每个元素 a 入栈时把当前栈的最小值 m 存储起来。在这之后无论何时，如果栈顶元素是 a，我们就可以直接返回存储的最小值 m。
//


// 方法一:    辅助栈

type MinStack struct {
	Array []int
	Min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	array := []int{}
	min := []int{}

	this := MinStack{
		Array: array,
		Min: min,
	}
	return this
}


func (this *MinStack) Push(x int)  {
	this.Array = append(this.Array,x)
	if len(this.Min) == 0 {
		this.Min = append(this.Min,x)
		return
	}
	if len(this.Min) > 0 && this.Min[len(this.Min)-1] >= x  {
		this.Min = append(this.Min,x)
	}
}

func (this *MinStack) Pop()  {
	array_last := this.Top()
	this.Array = this.Array[:len(this.Array)-1]
	min_last := this.Min[len(this.Min)-1]
	if min_last >= array_last {
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {

	return  this.Array[len(this.Array)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Min) == 0 {
		return -1
	}
	return this.Min[len(this.Min)-1]
}


// 方法二： 链表带有min属性

type MinStack struct {
	Value int
	Min   int
	Next *MinStack
}


/** initialize your data structure here. */
func Constructor() MinStack {
	head := MinStack{-1,-1,nil}
	return head
}

func (this *MinStack) Push(x int)  {
	head := this
	if head.Next == nil {
		node := &MinStack{}
		node.Value = x
		node.Min = x
		head.Next = node
		return
	}
	node := &MinStack{}
	node.Value = x
	next := head.Next
	node.Min = MinNum(x,next.Min)
	node.Next = next
	head.Next = node
}


func (this *MinStack) Pop()  {
	this.Next = this.Next.Next
}


func (this *MinStack) Top() int {
	head := this
	return head.Next.Value
}


func (this *MinStack) GetMin() int {
	head := this
	if head.Next == nil {
		return -1
	}
	return head.Next.Min
}

func MinNum(x,y int) int {
	if x > y {
		return y
	}
	return x
}