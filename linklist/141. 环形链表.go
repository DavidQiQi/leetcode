package main

//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环

// 思路： 1,用map方式来测试是否有重复项出现
// 2， 快慢指针方式模板
// 3, 快慢指针另外一种形式 slow := head fast := head.Next

//type ListNode struct {
//     Val int
//     Next *ListNode
// }
// 方式一： map[*ListNode]bool 方式，之前想到的是map[int]bool,如果出现重复数字的话，就会出现问题
func hasCycle(head *ListNode) bool {
	Cycle := make(map[*ListNode]bool)
	hh := head
	for hh != nil {
		_, ok := Cycle[hh]
		if !ok {
			Cycle[hh] = true
		} else {
			return true
		}
		hh = hh.Next
	}
	return false
}

// 方式二： 快慢指针,链表快慢指针模板
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 快慢指针 另外一种写法：参看leetcode的某一位作者
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
