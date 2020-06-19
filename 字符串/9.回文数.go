package 字符串

import (
	"fmt"
	"strconv"
)

// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//  你能不将整数转为字符串来解决这个问题吗？

// 思路一 1， 负数肯定不是回文数 2，直接对值进行除10，反向相加的值和x是否相等
// 思路二 把数值转换为字符串，对比开头和结尾的字符是否相等，相等则继续下一下，不想等直接return

func main() {
	array := 121
	fmt.Println(isPalindrome(array))
}

//思路一：

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	sum := 0
	tmp := x
	for tmp != 0 {
		sum = sum*10 + tmp%10
		tmp /= 10
		fmt.Println(sum)
	}
	if sum == x {
		fmt.Println(true)
		return true
	}
	return false
}

//思路二

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xx := strconv.Itoa(x)
	xx_len := len(xx)
	for i := 0; i < xx_len/2; i++ {
		if xx[i] != xx[xx_len-1] {
			return false
		}
	}
	return true
}
