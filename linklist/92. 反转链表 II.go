package main

//反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
//说明:
//1 ≤ m ≤ n ≤ 链表长度。
//
//示例:
//
//输入: 1->2->3->4->5->NULL, m = 2, n = 4
//输出: 1->4->3->2->5->NULL
//

//思考分析。 1，迭代法，根据输出的内容，如果m=n 的话，不反转，当m>n,差值是1 的话，1，2，3，4，5  m=2,n=3 从2开始反转，到3结束，结果为1，3，2，4，5
// 当m=2 n=4  4继续反转 我们利用这种后面的值 放到前面来 写代码

// 方法一 迭代  使用双指针，使用dummy节点 当m=1 n=5时，相当于全部反转一遍，
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m >= n {
		return head
	}
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	count := 1 // 计数
	pre := dummy
	cur := head
	for cur != nil {
		if count >= m && count < n {
			ccur := pre.Next
			next := cur.Next
			cur.Next = cur.Next.Next
			pre.Next = next
			next.Next = ccur
		} else {
			pre = cur
			cur = cur.Next
		}
		count++
	}
	return dummy.Next
}
