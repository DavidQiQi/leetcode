package main


// 剑指 Offer 59 - I. 滑动窗口的最大值

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 
//
//提示：
//
//你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
//


// 思路分析  k个值 求最大值，可以用排序来获取最大值 ，可以用双指针滑动,两种方法都使用了比较Max函数获取最大值

// 构建k的大顶堆 找出最大值

// 方法一：

func maxSlidingWindow(nums []int, k int) []int {
    if k < 2  {
        return nums
    }
    num := []int{}
    head := 0
    tail := k -1
    length := len(nums)
    max := nums[0]
    for  tail < length {
        max = Max(nums[head:tail])
        num = append(num,max)
        head += 1
        tail += 1
    }
    return num

}


func Max (arr []int) int {
    max := arr[0]
    for i:=0;i<len(arr);i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}


// 方法二： 双指针指向两端，比较进入和取出的数的大小，进入的大，则max = 进入的，取出的大，则比较一次 所有的数据，比每次比较节省了时间

func maxSlidingWindow(nums []int, k int) []int {
    if k < 2  {
        return nums
    }
    num := []int{}
    head := 1
    tail := k
    length := len(nums)
    max := Max(nums[:k])
    num = append(num,max)
    for head >= 1 && tail < length {
        if nums[head-1] == max {
            max = Max(nums[head:tail])
        }
        if nums[tail] > max {
            max = nums[tail]
        }
        num = append(num,max)
        head += 1
        tail += 1
    }
    return num
}


// 方法三： 三指针+哨兵法

//三指针
//1. 使用左指针left和右指针right分别表示滑动窗口的左右边界
//2. 使用最大值指针max指向当前的最大值,比较max 指针指向的 与left和right的大小

// 哨兵
//将left, right, max分别初始化为-1, k-2, -1的方式实现
//如果将left, right, max分别初始化为0, k-1, 0, 是无法从始至终使用上述流程的


// 更新max的2种情况
//1. 随着左边界的迭代, max溜出滑动窗口的范围, 即left > max
//2. 随着右边界的迭代, 右边新加入的值大于max位置的值





// 方法四：  大顶堆+双指针

import "container/heap"

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
        return nil
    }
    intHeap := make(IntHeap, k)
    copy(intHeap, nums[:k])
    heap.Init(&intHeap)
    results := make([]int, 0, len(nums))
    results = append(results, intHeap[0])

    for left, right := 0, k; right < len(nums); left, right = left + 1, right + 1 {
        intHeap.Remove(nums[left])
        heap.Push(&intHeap, nums[right])
        results = append(results, intHeap[0])
    }

    return results
}

type IntHeap []int

func (h *IntHeap) Remove(x int) {
    for i := 0; i < len(*h); i++ {
        if (*h)[i] == x {
            *h = append((*h)[:i], (*h)[i + 1:]...)
            break
        }
    }

    heap.Init(h)
}

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i int, j int) bool {
    return h[i] >= h[j]
}

func (h IntHeap) Swap(i int, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    length := len(*h)
    val := (*h)[length - 1]
    *h = (*h)[: length - 1]

    return val
}

//作者：jihonghe
//链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/solution/zui-da-dui-shuang-zhi-zhen-by-jihonghe/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
