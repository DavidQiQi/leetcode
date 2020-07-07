package main

// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
//
//你应当保留两个分区中每个节点的初始相对位置。
//
//示例:
//
//输入: head = 1->4->3->2->5->2, x = 3
//输出: 1->2->2->4->3->5


// 分析： 可以把以x为中心，左边小于的放在一个链表里，右边大于等于的放在一个链表里面，再合并两个链表

// 方法一：

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{}
	dummy1 := &ListNode{}
	cur := dummy
	last := dummy1
	for head != nil {
		if head.Val < x {
			cur.Next = head
			cur = cur.Next
		} else {
			last.Next = head
			last = last.Next
		}
		head = head.Next
	}
	last.Next = nil   // 注意最后链表的Next是nil
	cur.Next = dummy1.Next
	return dummy.Next
}
