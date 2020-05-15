package main

import "math"

// 633. 平方数之和
//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c。
//
//示例1:
//
//输入: 5
//输出: True
//解释: 1 * 1 + 2 * 2 = 5
// 
//
//示例2:
//
//输入: 3
//输出: False

func main() {
	n := 10
	fmt.Println(judgeSquareSum(n))

}

//方法一：
// 1 ，生成i值平方的map映射
// 2， 生成用c去减i * i 的值的map映射
// 3， 使用逆序循环，获取i的平方是否有对应的值


func judgeSquareSum(c int) bool {
	mm := make(map[int]int,int(math.Sqrt(float64(c))))
	mm1 :=  make(map[int]int,int(math.Sqrt(float64(c))))
	for i := 0 ;i<= c ;i++ {
		if i * i <= c {
			mm[i] = i * i
			mm1[c-mm[i]] = 0
		} else {
			break
		}
	}


	for i:= len(mm);i >= 0 ;i-- {
		if _, ok := mm1[mm[i]]; ok {
			return true
		}
	}
	return false
}

// 方法二
// 使用双指针  左右指针来获取对比值

func judgeSquareSum(c int) bool {
	low := 0
	high := int(math.Sqrt(float64(c)))
	for low <= high {
		sum := low * low +high * high
		if sum == c {
		return true
	}
		if sum < c {
		low += 1
	} else {
		high -=1
	}

	}
	return false
}


