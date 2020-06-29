package main

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
//
//示例 1:
//
//输入: 1->1->2
//输出: 1->2
//示例 2:
//
//输入: 1->1->2->3->3
//输出: 1->2->3
//

// 思考分析，由于是排好序的链表，重复元素是挨着的，只需要把相等的两个元素留一个  方法 迭代和递归
// 方法一： 迭代
// 方法二： 递归

// 方法一： 迭代
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil ||head.Next == nil  {
		return head
	}
	pre := head
	cur := head.Next
	for  cur.Next != nil {
		next := cur.Next
		if cur.Val == next.Val {
			pre.Next = next
			cur = next
		} else {
			pre = cur
			cur = next
		}
	}
	return head
}

//或者精简的迭代代码
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head;
	for cur != nil  && cur.Next != nil  {
		if (cur.Next.Val == cur.Val) {
			cur.Next = cur.Next.Next;
		} else {
			cur = cur.Next;
		}
	}
	return head;
}

// 方法二：递归   找出边界条件，递归下一次
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	if cur.Val == cur.Next.Val {
		cur.Next = cur.Next.Next
		cur = deleteDuplicates(cur)
	} else {
		cur.Next  = deleteDuplicates(cur.Next)
	}
	return head
}

// 别人写的递归
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head                             // 终止条件
	}
	head.Next  = deleteDuplicates(head.Next)     // 递归调用
	if head.Val == head.Next.Val {
		head = head.Next                      // 去重
	}
	return head
}

//作者：marvinysh
//链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/solution/go-di-gui-shi-xian-by-marvinysh/
