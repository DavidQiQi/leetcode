package 位运算

import "fmt"

//136. 只出现一次的数字
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,1]
//输出: 1
//示例 2:
//
//输入: [4,1,2,1,2]
//输出: 4

func main() {
	nums := []int{4, 4, 5, 7, 7, 9, 9}
	fmt.Println(singleNumber(nums))
}

//思路一： 位运算 异或 相同的两个数相异或为0  0与任何数a异或为a     a^a = 0   a^ 0 = a

//复杂度分析
// 时间复杂度：O(n) ，其中 n 是数组长度。只需要对数组遍历一次。
//
//空间复杂度：O(1) 。

func singleNumber(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	return sum
}

// 思路二： 使用空间解决，
//复杂度分析
// 时间复杂度O(n)
// 空间复杂度O(n)

func singleNumber(nums []int) int {
	mm := make(map[int]int, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		_, ok := mm[nums[i]]
		if ok {
			delete(mm, nums[i])
		} else {
			mm[nums[i]] += 1
		}
	}
	for k, _ := range mm {
		return k
	}
	return -1
}
