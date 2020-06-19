package 位运算

import "math/bits"

//
//颠倒给定的 32 位无符号整数的二进制位。
//
//
//
//示例 1：
//
//输入: 00000010100101000001111010011100
//输出: 00111001011110000010100101000000
//解释: 输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
//     因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。
//示例 2：
//
//输入：11111111111111111111111111111101
//输出：10111111111111111111111111111111
//解释：输入的二进制串 11111111111111111111111111111101 表示无符号整数 4294967293，
//     因此返回 3221225471 其二进制表示形式为 10111111111111111111111111111111 。
//

// 思路分析：  逆序二进制串   bits.reverseBit32  可以直接翻转 ，也可以根据左右位移、与运算进行翻转
// 思路一：  使用官方提供的库
// 思路二：  对二进制串 进行右位移后 与1 得到的值 再左位移
// 思路三：  对二进制串 进行左右位移16位，然后相与   左右位移8位 翻转8位，然后再4位  再2位

// 思路一：
func reverseBits(num uint32) uint32 {
	return bits.Reverse32(num)
}

// 思路二：

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		ans := (num >> i) & 1
		res |= ans << (31 - i)
	}
	return res
}

func reverseBits(num uint32) uint32 {
	bits, count := uint32(0), uint32(32)
	for count > 0 {
		bits |= (num & 1) << (count - 1)
		num >>= 1
		count--
	}
	return bits
}

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		x := num % 2
		res = res*2 + x
		num = num / 2
	}
	return res
}

// 思路三：
func reverseBits(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}
