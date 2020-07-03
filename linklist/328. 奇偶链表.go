package main


//  给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
//
//请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
//
//示例 1:
//
//输入: 1->2->3->4->5->NULL
//输出: 1->3->5->2->4->NULL
//示例 2:
//
//输入: 2->1->3->5->6->4->7->NULL
//输出: 2->3->6->7->1->5->4->NULL
//说明:
//
//应当保持奇数节点和偶数节点的相对顺序。
//链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
//


// 思路分析：  1，拆分出奇偶节点，重新拼接在一起    2， 把偶数节点摘出来，挂到head后面


// 方法一： 分离出奇偶两个链表，拼接
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	evendummy, odddummy := &ListNode{},&ListNode{}
	even := evendummy  // 偶数
	odd := odddummy    // 奇数

	count := 1
	for head != nil {
		if count % 2 == 0 {
			even.Next = head
			even = even.Next
		} else {
			odd.Next = head
			odd = odd.Next
		}
		count++
		head = head.Next
	}
	even.Next = nil
	odd.Next = evendummy.Next
	return odddummy.Next
}

// 方法二：  思路 当count 为偶数时，摘除节点，放到dummy.Next节点
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 1
	dummy := &ListNode{}
	d := dummy
	cur := head

	for cur != nil && cur.Next != nil {
		// 奇数节点
		if count % 2  == 1 {
			next := cur.Next
			cur.Next = cur.Next.Next
			next.Next = nil
			d.Next = next
			d = d.Next
			count++
		}
		cur = cur.Next
		count++
	}
	hh := head
	for hh.Next != nil {
		hh = hh.Next
	}
	hh.Next = dummy.Next
	return head
}