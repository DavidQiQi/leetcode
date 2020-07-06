package main

import "fmt"

// 归并排序  自顶向下
// 归并排序算法的使用情景
//归并排序算法和快速排序算法是java.util.Arrays中使用的排序算。
// 对于一般的基本数据类型，Arrays.sort函数使用双轴快速排序算法，而对于对象类型使用归并排序（准确的说使用的是TimSort排序算法，
// 它是归并排序的优化版本）。这样做的原因有两点，第一个原因，归并排序是稳定的，而快速排序不是稳定的。第二个原因，
// 对于基本数据类型，排序的稳定性意义不大，但对于复合数据类型（比如对象）排序的稳定性就能帮助我们保持排序结果的某些性质

//####自顶向下的归并排序
// 自顶向下的排序算法就是把数组元素不断的二分，直到子数组的元素个数为一个，因为这个时候子数组必定是已有序的，
// 然后将两个有序的序列合并成一个新的有序的序列，两个新的有序序列又可以合并成另一个新的有序序列，以此类推，直到合并成一个有序的数组

//####自底向上的归并排序
// 自底向上的归并排序算法的思想就是数组中先一个一个归并成两两有序的序列，两两有序的序列归并成四个四个有序的序列，
// 然后四个四个有序的序列归并八个八个有序的序列，以此类推，直到，归并的长度大于整个数组的长度，此时整个数组有序。
// 需要注意的是数组按照归并长度划分，最后一个子数组可能不满足长度要求，
// 这个情况需要特殊处理。自顶下下的归并排序算法一般用递归来实现，而自底向上可以用循环来实现。

func main() {
	arr := []int{4,7,0,5,6,9,1,3,2}
	result := mergeSort(arr)
	fmt.Println(result)
}


// 自顶向下的归并排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	return merge(mergeSort(arr[:i]),mergeSort(arr[i:]))
}

func merge(left ,right []int) []int {
	result := make([]int,0)
	i, j := 0, 0
	l, r := len(left),len(right)

	for i < l && j < r {
		if left[i] < right[j] {
			result = append(result,left[i])
			i++
			continue
		}
		result = append(result,right[j])
		j++
	}
	result = append(result,left[i:]...)
	result = append(result,right[j:]...)
	return result
}


// 自底向上排序
// 自底向上的归并排序算法的思想就是数组中先一个一个归并成两两有序的序列，两两有序的序列归并成四个四个有序的序列，
// 然后四个四个有序的序列归并八个八个有序的序列，以此类推，直到，归并的长度大于整个数组的长度，此时整个数组有序。
// 需要注意的是数组按照归并长度划分，最后一个子数组可能不满足长度要求，
// 这个情况需要特殊处理。自顶下下的归并排序算法一般用递归来实现，而自底向上可以用循环来实现。



// 自底向上归并排序
func MergeSort2(arr []int , begin,end  int )  {
	// 步数为1开始，step长度的数组表示一个有序的数组
	step := 1
	// 范围大于 step 的数组才可以进入归并
	for end - begin > step {
		// 从头到尾对数组进行归并操作
		// step << 1 = 2 * step 表示偏移到后两个有序数组将它们进行归并
		for i := begin;i<end;i+=step <<1 {
			lo := i         // 第一个有序数组的上界
			mid := lo +step    // 第一个有序数组的下界，第二个有序数组的上界
			hi := lo + (step <<1 )  // 第二个有序数组的下界

			// 不存在第二个数组，直接返回
			if mid > end {
				return
			}
			// 第二个数组长度不够
			if hi > end {
				hi = end
			}
			// 两个有序数组进行合并
			merge2(arr,lo,mid,hi)
		}
		// 上面的 step 长度的两个数组都归并成一个数组了，现在步长翻倍
		step <<= 1
	}
}

func merge2(array []int,begin,mid,end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 array[begin,mid),array[mid,end)
	leftSize := mid - begin    // 左边数组的长度
	rightSize := end - mid     // 右边数组的长度
	newSize := leftSize +rightSize  // 辅助数组的长度
	result  :=  make([]int,0,newSize)
	l, r := 0,0

	for l < leftSize && r < rightSize {
		lValue := array[begin+l] // 左边数组的元素
		rValue := array[mid+r]   // 右边数组的元素
		// 小的元素先放进辅助数组里
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}

	// 将剩下的元素追加到辅助数组后面
	result = append(result, array[begin+l:mid]...)
	result = append(result, array[mid+r:end]...)

	// 将辅助数组的元素复制回原数组，这样该辅助空间就可以被释放掉
	for i := 0; i < newSize; i++ {
		array[begin+i] = result[i]
	}
	return
}
