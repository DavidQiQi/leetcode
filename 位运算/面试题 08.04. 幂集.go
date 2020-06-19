package main

//幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
// 输入： nums = [1,2,3]
// 输出：
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//

// 思考分析：这道题 标记 回溯算法，暂时不会回溯， 用二进制方式来实现，
//
// 方法一： 二进制串来实现:  nums 有3个数  1，2，3  总共有1 << len(nums) 中排列方法，
// num = 3 二进制串为011  对应的 三个数 1 不取 2，3 取 ，，   （num >> 2）& 1  第一位是否取    （num >> 1）& 1 第二位为1 是否取
// （num >> 0）& 1 第三位为1 是否取

// 方法二： 回溯 ，暂时不会写， 等把回溯弄明白了，再回来补充

// 方法一：
func subsets(nums []int) [][]int {
	// lenn := 1 << len(nums)
	array := [][]int{}
	for num := 0; num < 1<<len(nums); num++ {
		varray := []int{}
		for i := 0; i < len(nums); i++ {
			a := num >> i & 1
			if a == 1 {
				varray = append(varray, nums[i])
			}
		}
		array = append(array, varray)
	}
	return array
}
