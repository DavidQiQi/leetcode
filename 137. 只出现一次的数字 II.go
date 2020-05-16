package main

import "fmt"

// 137. 只出现一次的数字 II
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,3,2]
//输出: 3
//示例 2:
//
//输入: [0,1,0,1,0,1,99]
//输出: 99
//

// 思路0： 逻辑有问题，不成功， 三个相同的数，之和再模3 余数为0，但是要求单独的数不为3的倍数才可以

// 思路一：根据相同位 三个1为3，再模3 重复的1就被消掉，剩下的就是单独数字的状态

func main() {
	nums := []int{2, 2, 2, 4}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	res := 0
	for index := 0; index < 64; index++ {
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum += (nums[i] >> index) & 1
		}
		res |= sum % 3 << index
	}
	return res
}

// 思路二  来自网上，自己没有想到 通过数学方式来解决  3 *(a+b+c) = 3a + 3b +c + 2c 需要消耗集合空间

// 思路三  来自网上，大牛的解题思路 one = 0,two=0  two=(two ^ next) & ^one    one =(one ^ next) & ^two  解法相当牛X

func singleNumber(nums []int) int {
	one, two := 0, 0
	for i := 0; i < len(nums); i++ {
		two = (two ^ nums[i]) & ^one
		one = (one ^ nums[i]) & ^two
	}
	return two
}

// 思路四：  使用map映射来解决，需要消耗空间
func singleNumber(nums []int) int {
	nums_map := make(map[int]int, len(nums)-1)
	for _, v := range nums {
		nums_map[v] += 1
		if nums_map[v] == 3 {
			delete(nums_map, v)
		}
	}
	for k, _ := range nums_map {
		return k
	}
	return -1
}
