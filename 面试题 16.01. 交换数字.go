package main

// 编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。
//
//示例：
//
//输入: numbers = [1,2]
//输出: [2,1]
//提示：
//
//numbers.length == 2

// 方法1， 使用交换方式  a,b = b,a
// 方法2， 用交换数组下标方式
// 方法3，  用异或的方式  a ^ a =0  0 ^ a = a  a ^b ^a = b
// 方法4    加减法

func main() {
	
}
//方法1
func swapNumbers(numbers []int) []int {
	numbers[0],numbers[1] = numbers[1],numbers[1]
	return numbers
}



// 方法2
func swapNumbers(numbers []int) []int {
	result := []int{}
	result = append(result,numbers[1],numbers[0])
	return result
}


// 方法3
func swapNumbers(numbers []int) []int {
	numbers[0] ^= numbers[1]
	numbers[1] ^= numbers[0]
	numbers[0] ^= numbers[1]
	return numbers
}

// 方法4
func swapNumbers(numbers []int) []int {
	numbers[0] += numbers[1]
	numbers[1] = numbers[0]  -  numbers[1]
	numbers[0] = numbers[0]  - numbers[1]
	return numbers
}
