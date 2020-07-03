package main

// 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
//将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//示例 1:
//
//给定链表 1->2->3->4, 重新排列为 1->4->2->3.
//示例 2:
//
//给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.


// 思路分析 这道题跟回文链表相似  1,可以使用数组存储节点，交换数组元素，输出链表，
// 2 根据题给出的信息，中间节点后面的反转，和前面的拼接到一块    3， 也可以用递归的方法

// 方法一： 简单，自行尝试

// 方法二：
func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	// 第一步 找出中间节点，分奇偶个数
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newhead := slow.Next
	slow.Next = nil
	// 反转slow
	var pre *ListNode = nil
	cur := newhead
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 轮流添加节点
	dummy := &ListNode{}
	d := dummy
	for head  != nil || pre != nil {
		if head != nil {
			d.Next = head
			head = head.Next
			d = d.Next
		}
		if pre != nil {
			d.Next = pre
			pre = pre.Next
			d = d.Next
		}

	}
	head = dummy.Next
}

//  方法三：
//  参考：https://leetcode-cn.com/problems/reorder-list/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-34/
// 思路:
//从外向内旋转
//实现:
//从内向外实现
//终止条件：到达中心 len == 1 或 len == 2
//递归返回：返回外一层的尾结点
//递归实现：
//输入：当前层头结点，当前层及内层剩余节点数；返回：外一层为节点数
//递归：由外箱内，逐层剥开 len - 2 ； 终止条件：剥到了中心 len == 1 len == 2
//递归中的连接：连接本层首尾 head 与 tail，同时将本层与内一层连接 tail 与 head.next
//

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	var len int
	// 计算链表长度
	tmp := head
	for tmp != nil {
		len++
		tmp = tmp.Next
	}
	// 递归
	reorderListHelper(head, len)
}

func reorderListHelper(head *ListNode, len int) *ListNode{
	var upTail *ListNode  // 返回上一层(外一层)的 尾结点
	// 终止条件  旋转到中心
	if len == 1 {
		upTail = head.Next
		head.Next = nil
		return upTail
	}
	if len == 2 {
		upTail = head.Next.Next
		head.Next.Next = nil
		return upTail
	}
	tail := reorderListHelper(head.Next, len - 2) // head.Next 为往里一层  及剩余节点数量  返回值为外一层的尾结点
	// 当前层的头结点 为head
	// 先记录要返回当前层 的 外一层的尾结点  就是当前层的尾结点 向后一个
	upTail = tail.Next
	// 记录里一层的头结点 用于连接
	inHead := head.Next
	// 当前层的连接
	head.Next = tail
	// 当前层尾结点 连接里面一层 头结点
	tail.Next = inHead
	return upTail
}

