package main

// 中等难度

//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
//你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
//进阶：
//
//如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
//
//
//
//示例：
//
//输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 8 -> 0 -> 7

//思路分析  根据题意, 一，可以先使用链表反转，再相加   二，可以使用栈存储两个链表，并依次pop栈顶元素，相加，输出到链表中， 方法三：递归

// 方法一： 链表反转
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	// 反转两个列表
	L1 := Reverse(l1)
	L2 := Reverse(l2)
	ll1 := L1
	ll2 := L2
	flag := false
	dummy := &ListNode{}
	d := dummy

	for ll1 != nil || ll2 != nil {
		value := 0
		if flag == true {
			value += 1
			flag = false
		}
		if ll1 != nil {
			value += ll1.Val
			ll1 = ll1.Next
		}
		if ll2 != nil {
			value += ll2.Val
			ll2 = ll2.Next
		}
		if value >= 10 {
			value -= 10
			flag = true
		}
		d.Next = &ListNode{Val: value}
		d = d.Next
	}
	if flag == true {
		neww := &ListNode{Val: 1}
		d.Next = neww
	}
	head := Reverse(dummy.Next)
	return head
}

func Reverse(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 方法二：
// 思路  用两个栈存放两组信息,先都压入栈中，再依次出站，相加，计算的结构再放到链表d.Next
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	nums1 := []int{}
	nums2 := []int{}
	for l1 != nil {
		nums1 = append(nums1, l1.Val)
		l1 = l1.Next

	}
	for l2 != nil {
		nums2 = append(nums2, l2.Val)
		l2 = l2.Next

	}
	flag := false
	dummy := &ListNode{}
	d := dummy
	i := len(nums1) - 1
	j := len(nums2) - 1
	for i != -1 || j != -1 {
		value := 0
		if flag == true {
			value += 1
			flag = false
		}
		if i != -1 {
			value += nums1[i]
			i--
		}
		if j != -1 {
			value += nums2[j]
			j--
		}
		if value >= 10 {
			value %= 10
			flag = true
		}
		pre := &ListNode{Val: value}
		pre.Next = d.Next
		d.Next = pre
	}
	if flag == true {
		next := &ListNode{Val: 1}
		next.Next = d.Next
		d.Next = next
	}
	return dummy.Next
}

//  方法三： 没有实现出来，抄于网上
// 如何实现两个长度不同的链表尾部对齐?
//计算链表长度作为,递归的深度,通过深度控制是否next指针
//还有什么方法?
//最简单的方式是 将链表反转然后直接遍历.
//在不引起副作用的前提下,递归方法也可以改进成双栈实现.
//从中学到什么?
//对于此题,双栈实现要比递归实现简单易懂,可见 如果递归理顺不清,可以从辅助栈的思路来思考,这对限制时间的思考,较为有用

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var d1, d2 int
	a, b := l1, l2
	for a != nil {
		d1++
		a = a.Next
	}
	for b != nil {
		d2++
		b = b.Next
	}
	res, c := addTwoNumbers2(l1, l2, d1, d2)
	if c > 0 {
		return &ListNode{Val: c, Next: res}
	}
	return res
}

func addTwoNumbers2(l1, l2 *ListNode, d1, d2 int) (res *ListNode, c int) {
	if l1 != nil && l2 != nil {
		if l1.Next == nil && l2.Next == nil {
			n := l1.Val + l2.Val
			c = n / 10
			res = &ListNode{Val: n % 10, Next: nil}
			return
		}
	}
	var a *ListNode
	var b, n int
	if d1 > d2 {
		a, b = addTwoNumbers2(l1.Next, l2, d1-1, d2)
		n = l1.Val + b
	} else if d1 < d2 {
		a, b = addTwoNumbers2(l1, l2.Next, d1, d2-1)
		n = l2.Val + b
	} else {
		a, b = addTwoNumbers2(l1.Next, l2.Next, d1-1, d2-1)
		n = b + l1.Val + l2.Val
	}
	res = &ListNode{Val: n % 10, Next: a}
	c = n / 10
	return
}
