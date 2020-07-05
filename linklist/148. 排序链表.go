package main

//在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
//
//示例 1:
//
//输入: 4->2->1->3
//输出: 1->2->3->4
//示例 2:
//
//输入: -1->5->3->4->0
//输出: -1->0->3->4->5

// 分析: O(nlogn)这个排序,有快排、归并，在排序中，先想到快排，因为空间复杂度是常数级的，测试了一下，快排好像不行。就试了一下归并，
// 之前排序的时候没有好好写归并，在这个题上面有些思路

//方法一: 归并法
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head1 := slow.Next
	slow.Next = nil
	head2 := head
	return mergerList(sortList(head1), sortList(head2))
}

func mergerList(s1 *ListNode, s2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	d := dummy
	for s1 != nil && s2 != nil {
		if s1.Val <= s2.Val {
			d.Next = s1
			s1 = s1.Next
		} else {
			d.Next = s2
			s2 = s2.Next
		}
		d = d.Next
	}
	if s1 != nil {
		d.Next = s1
	}
	if s2 != nil {
		d.Next = s2
	}
	return dummy.Next
}
