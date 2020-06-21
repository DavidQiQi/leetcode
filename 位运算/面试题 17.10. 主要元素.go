package main

//数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。
//
//示例 1：
//
//输入：[1,2,5,9,5,9,5,5,5]
//输出：5
//
//示例 2：
//
//输入：[3,2]
//输出：-1
//
//示例 3：
//
//输入：[2,2,1,1,1,2,2]
//输出：2
//
//说明：
//你有办法在时间复杂度为 O(N)，空间复杂度为 O(1) 内完成吗？

// 思路分析  1， 位运算 计数每个位上的0，1的个数，大于一般则记录，否则为-1
// 2， 哈希方式 来获取
// 3 使用投票方法，如果位数大于一半的话，按照不同抵消的话，最后应该>=1才对

// 方法一：  时间复杂度O(64n) 时间复杂度O(1)
func majorityElement(nums []int) int {
	half := len(nums) / 2
	res := 0
	for i := 0; i < 64; i++ {
		count0 := 0
		count1 := 0
		ans := 0
		for _, v := range nums {
			if (v >> i & 1) == 1 {
				count1++
			} else {
				count0++
			}
		}
		if count0 > half || count1 > half {
			if count0 > half {
				ans = 0
			} else {
				ans = 1
			}
		} else {
			return -1
		}
		res |= ans << i
	}
	return res
}

// 方法二：  空间复杂度 O(n) 时间复杂度O(n)
func majorityElement(nums []int) int {
	half := len(nums) / 2
	numMap := make(map[int]int, len(nums))
	for _, v := range nums {
		numMap[v] += 1
	}
	for k, v := range numMap {
		if v > half {
			return k
		}
	}
	return -1
}

//方法三：  空间复杂度O(1), 时间复杂度O(n)
func majorityElement(nums []int) int {
	major := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if major == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			major = nums[i]
			count++
		}
	}
	if count < 1 {
		return -1
	} else {
		return major
	}
}
