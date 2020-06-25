package main

//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//
//说明：不允许修改给定的链表

// 思路分析： 1，map方式，找到第一次重复出现的listnode
// 2， 使用141的快慢指针方式模板，找出第一次快慢指针相遇的点，再设置cur=head 和slow 每次都走一步，相遇的时候就是入环点

// 方法一： 使用额外空间，代码会简洁
func detectCycle(head *ListNode) *ListNode {
	Cycle := make(map[*ListNode]bool)
	hh := head
	for hh != nil {
		_, ok := Cycle[hh]
		if !ok {
			Cycle[hh] = true
		} else {
			return hh
		}
		hh = hh.Next
	}
	return nil
}

// 方法二：快慢指针找到相遇点，再求出入环点， 使用快慢指针的模板

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 找到快慢指针的相遇点
			cur := head
			for cur != slow { //从head 开始和slow 再次相遇时，则为入环点
				slow = slow.Next
				cur = cur.Next
			}
			return cur
		}
	}
	return nil
}
