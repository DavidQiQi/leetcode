package main

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

// 思路0： 1->2->3->4->5->nil  使用栈的功能先进后出，全部都压入栈，再一次弹出，创建新的链表

// 思路一： 双指针迭代方法
//我们可以申请两个指针，第一个指针叫 pre，最初是指向 null 的。
//第二个指针 cur 指向 head，然后不断遍历 cur。
//每次迭代到 cur，都将 cur 的 next 指向 pre，然后 pre 和 cur 前进一位。
//都迭代完了(cur 变成 null 了)，pre 就是最后一个节点了。

// 思路二： 递归方法  不太好理解
//递归解法
//这题有个很骚气的递归解法，递归解法很不好理解，这里最好配合代码和动画一起理解。
//递归的两个条件：
//
//终止条件是当前节点或者下一个节点==null
//在函数内部，改变节点的指向，也就是 head 的下一个节点指向 head 递归函数那句
//
//head.next.next = head
//很不好理解，其实就是 head 的下一个节点指向head。
//递归函数中每次返回的 cur 其实只最后一个节点，在递归函数内部，改变的是当前节点的指向。

// 方法三 插入法：

// 方法一：双指针迭代方法
func reverseList(head *ListNode) *ListNode {
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
func reverseList(head *ListNode) *ListNode {
	//递归终止条件是当前为空，或者下一个节点为空
	if head == nil || head.Next == nil {
		return head
	}
	//这里的cur就是最后一个节点
	cur := reverseList(head.Next)
	head.Next.Next = cur
	head.Next = nil
	return cur
}

// 方法三：
func reverseList(head *ListNode) *ListNode {
	// 边界条件
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next // 获取第二个值
	head.Next = nil  // 重新利用head, 使头结点的next指向nil, 1-> nil 后面的值都插入到1的前面
	for cur != nil {
		next := cur.Next
		cur.Next = head // cur.Next cur 是第二个值，第二个值 指向head头
		head = cur      // head = cur head 等于cur 头节点
		cur = next
	}
	return head
}

// 其他方法

//妖魔化的双指针
//原链表的头结点就是反转之后链表的尾结点，使用 headhead 标记 .
//定义指针 curcur，初始化为 headhead .
//每次都让 headhead 下一个结点的 nextnext 指向 curcur ，实现一次局部反转
//局部反转完成之后，curcur 和 headhead 的 nextnext 指针同时 往前移动一个位置
//循环上述过程，直至 curcur 到达链表的最后一个结点

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	cur := head
	for head.Next != nil {
		t := head.Next.Next
		head.Next.Next = cur
		cur = head.Next
		head.Next = t
	}
	return cur
}
