package main

//给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
//
//示例：
//
//给定一个链表: 1->2->3->4->5, 和 n = 2.
//
//当删除了倒数第二个节点后，链表变为 1->2->3->5.
//说明：
//
//给定的 n 保证是有效的。
//
//进阶：
//
//你能尝试使用一趟扫描实现吗？

// 思路分析: 1, 使用dummy节点，pre=dummy、cur=dummy.Next, pre、cur都往前走一步，计算n次 cur.next为nil时，把pre.next = cur.next
// 2, 使用dummy节点，fast,slow = dummy,dummy, fast 先走n步，然后fast\slow 同时都走一步

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := dummy.Next

	for {
		curr := cur
		for i := 0; i < n; i++ {
			curr = curr.Next
		}
		if curr == nil {
			break
		}
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = cur.Next

	return dummy.Next
}

// 方法二  改良版dummy

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
