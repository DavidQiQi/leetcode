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

// 1 以x作为值的边界 分别创建2个链表，一个比x值大与或者等于，一个比x值小
//2 需要把大的链表放到小的链表后面，最好的方式可以创建2个哨兵节点，作为两个链表的开始节点
//3 遍历当前指针，然后将当前值分别连接到大值链表或者小值链表
//4 需要特殊处理的是 需要将大值链表的尾节点设置为空，因为可能当前大值链表的尾节点next 值不为空 造成链表死循环

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
