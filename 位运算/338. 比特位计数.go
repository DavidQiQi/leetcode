package 位运算

import "fmt"

// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
//示例 1:
//
//输入: 2
//输出: [0,1,1]
//示例 2:
//
//输入: 5
//输出: [0,1,1,2,1,2]
//进阶:
//
//给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
//要求算法的空间复杂度为O(n)。
//你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。

// 方法1，先求出每个数的二进制1的个数，再轮询
// 方法2， 使用递归方式获取
// 方法3， 对递归方法改进，记录生成的内容到数组,效率会高一些

func main() {
	fmt.Println(countBits(5))
}

// 方法一：
func countBits(num int) []int {
	result := []int{}
	for i := 0; i <= num; i++ {
		result = append(result, numBits(i))
	}
	return result
}

func numBits(num int) int {
	count := 0
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}

// 方法二：

func countBits(num int) []int {
	result := []int{}
	for i := 0; i <= num; i++ {
		result = append(result, numBits(i))
	}
	return result
}

func numBits(num int) int {
	if num == 0 {
		return 0
	}
	if num == 2 || num == 1 {
		return 1
	}
	return numBits(num&(num-1)) + 1
}

// 方法3：
func countBits(num int) []int {
	nums := make([]int, num+1)
	for i := 1; i <= num; i++ {
		nums[i] = nums[i&(i-1)] + 1
	}
	return nums
}
