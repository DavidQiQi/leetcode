package main

// 693. 交替位二进制数
//给定一个正整数，检查他是否为交替位二进制数：换句话说，就是他的二进制数相邻的两个位数永不相等。
//
//示例 1:
//
//输入: 5
//输出: True
//解释:
//5的二进制数是: 101
//示例 2:
//
//输入: 7
//输出: False
//解释:
//7的二进制数是: 111
//示例 3:
//
//输入: 11
//输出: False
//解释:
//11的二进制数是: 1011
// 示例 4:
//
//输入: 10
//输出: True
//解释:
//10的二进制数是: 1010


// 分析：1  01 ，2 10，  5  101 ， 10 1010  格式交替二进制
// 方法1：  数学方法 根据1，2，5，10，21，42，85  变换公式
// 方法2：  数组二进制位 比较
// 方法3：  二进制异或方式  2 >> 1 变成01  10 ^ 01 变成了全 11 再加1 就变成了100 ， 100 & 011 == 0

func main() {
	
}

// 方法1：
// 交替位二进制 规则1 10 101 1010 10101 101010 1010101 这种格式 根据这种格式获取的值位1,2，5，10，21，42，85
//  奇数位 为前一个数的2n 偶数位 为前一个数的2n+1,然后比较num 与n 是否相等
func hasAlternatingBits(n int) bool {
	num := 1
	for num < n {
		if num %2 == 1 {
			num = 2 * num
		} else {
			num = 2 * num +1
		}
	}
	if num == n {
		return true
	}
	return false
}


// 方法二：用数组比较法
func hasAlternatingBits(n int) bool {
	if n == 1 {
		return true
	}
	nums := make([]int,32)
	nums[0] = n & 1
	n = n >>1
	count := 1
	for n != 0 {
		nums[count] = n & 1
		if nums[count] == nums[count-1] {
			return false
		}
		count++
		n = n >>1
	}
	return true
}

// 方法三：  二进制异或方式  2 >> 1 变成01  10 ^ 01 变成了全 11 再加1 就变成了100 ， 100 & 011 == 0
// 对二进制111 形式不敏感

func hasAlternatingBits(n int) bool {
	nn := n ^ (n >>1) +1
	return nn & (nn-1) == 0
}