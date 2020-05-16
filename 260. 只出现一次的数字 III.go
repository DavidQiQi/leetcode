package main

import "fmt"

//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
//
//示例 :
//
//输入: [1,2,1,3,2,5]
//输出: [3,5]
//注意：
//
//结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
//你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
//

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))

}

// 思路一： 根据map
func singleNumber(nums []int) []int {
	mm := make(map[int]int, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		mm[nums[i]] += 1
		if mm[nums[i]] == 2 {
			delete(mm, nums[i])
		}
	}
	sli := []int{}
	for k, _ := range mm {
		sli = append(sli, k)
	}
	return sli
}

// 思路二： 根据异或规则 相同的数字异或为0 ，所有数异或成一个结果result，根据这个结果二进制第一次出现1的位来划分数组

func singleNumber(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum ^= nums[i]
	}
	index := 0
	for {
		if (sum>>index)&1 == 1 {
			break
		}
		index += 1
	}
	one := 0
	two := 0
	for i := 0; i < len(nums); i++ {
		if ((nums[i] >> index) & 1) == 0 {
			one ^= nums[i]
		} else {
			two ^= nums[i]
		}
	}
	return []int{one, two}
}
