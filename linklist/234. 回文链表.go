package main

// 请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

// 思路分析： 使用空间复杂度为n 可以使用数组，反转链表 方式 ，要想让复杂度为O(1) ，就只能在分拆链表

// 方法一：   从链表转换数组
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	nums := []int{}
	for head != nil {
		nums = append(nums,head.Val)
		head = head.Next
	}
	count := len(nums)
	for i:=0;i<count/2;i++{
		if nums[i] != nums[count-1-i]{
			return false
		}
	}
	return true
}

// 方法二:  反转链表，新旧相等就为TRUE
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next== nil {
		return true
	}
	// 赋值新列表
	cur := head
	newhead :=  &ListNode{Val:-1}
	nn := newhead
	for cur != nil {
		nn.Next = &ListNode{Val:cur.Val}
		cur = cur.Next
		nn = nn.Next
	}
	//  反转新列表
	var pre *ListNode = nil
	ccur := newhead.Next
	for ccur != nil {
		next := ccur.Next
		ccur.Next = pre
		pre = ccur
		ccur = next
	}
	//对比Val值
	for head != nil {
		if head.Val !=pre.Val {
			return false
		}
		head = head.Next
		pre = pre.Next
	}
	return true
}


// 方法三： 链表中间节点断开，后续链表再反转
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow,fast := head ,head
	// 取中间值slow
	for fast !=nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//当fast不为空时，为奇数，slow=slow.Next
	if fast !=nil {
		slow = slow.Next
	}
	//反转slow
	var pre *ListNode = nil
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 比较pre 和head的值
	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}

// 方法三：实现2
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	tail := reverseList(slow.Next)
	slow.Next = nil
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head, tail = head.Next, tail.Next
	}
	return true
}
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}
	return prev
}