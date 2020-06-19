package 位运算

// 461. 汉明距离

// 两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
//
//给出两个整数 x 和 y，计算它们之间的汉明距离。
//
//注意：
//0 ≤ x, y < 231.
//
//示例:
//
//输入: x = 1, y = 4
//
//输出: 2
//
//解释:
//1   (0 0 0 1)
//4   (0 1 0 0)
//       ↑   ↑
//
//上面的箭头指出了对应二进制位不同的位置。

// 思考 汉明距离，就是求出相同位 值不通
// 方法1：  可以把x,y 作为二进制，按位比较
// 方法2：  使用异或，再用n&(n-1)求出1的个数

func main() {

}

// 方法1：
func hammingDistance(x int, y int) int {
	count := 0
	for x != y {
		a := x & 1
		b := y & 1
		if a != b {
			count++
		}
		x = x >> 1
		y = y >> 1
	}
	return count
}

//方法2：  使用异或，再用n&(n-1)求出1的个数
func hammingDistance(x int, y int) int {
	count := 0
	res := x ^ y
	for res != 0 {
		res &= res - 1
		count++
	}
	return count
}
