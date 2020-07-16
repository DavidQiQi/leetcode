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

// 方法一：暴力法

//复杂度分析
//
//时间复杂度：{O}(N k)O(Nk)。其中 N 为数组中元素个数。
//
//空间复杂度：{O}(N - k + 1)O(N−k+1)，用于输出数组。


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

// 方法三： 双端队列

// 双端队列和普通队列最大的不同在于，它允许我们在队列的头尾两端都能在 O(1)的时间内进行数据的查看、添加和删除

// 可以利用一个双端队列来表示这个窗口。这个双端队列保存当前窗口中最大那个数的下标，
// 双端队列新的头总是当前窗口中最大的那个数。 同时，有了这个下标，
// 我们可以很快地知道新的窗口是否已经不再包含原来那个最大的数，如果不再包含， 我们就把旧的数从双端队列的头删除。
// 按照这样的操作，不管窗口的长度是多长，因为数组里的每个数都分别被压入和弹出双端队列一次，所以我们可以在 O(n) 的时间里完成任务。

func maxSlidingWindow(nums []int,k int) []int {
    if k < 1  {
        return nil
    }
    if k < 2  {
        return nums
    }
    windows , res := []int{},[]int{}
    for i,x  := range nums {
        // 判断滑动窗口头部下标是否在i-k差值以内，大于的话，就说明windows的头结点是正常的滑动窗口范围
        if i >= k && windows[0] <= i - k   {
            windows = windows[1:]
        }
        // 判断尾部元素和新入队元素的大小，如果新入队元素大于所有windows的元素，则windows 清理尾部元素
        for len(windows) > 0  && nums[windows[len(windows)-1]] < x {
            windows = windows[:len(windows)-1]
        }
        // 把下标信息入队
        windows = append(windows,i)

        // 获取每次的windows的开头元素
        if i >= k -1 {
            res = append(res,nums[windows[0]])
        }
    }
    return res
}


//方法四： 动态规划

//直觉
//
//这是另一个 {O}(N)O(N) 的算法。本算法的优点是不需要使用 数组 / 列表 之外的任何数据结构。
//
//算法的思想是将输入数组分割成有 k 个元素的块。
//若 n % k != 0，则最后一块的元素个数可能更少

// 开头元素为 i ，结尾元素为 j 的当前滑动窗口可能在一个块内，也可能在两个块中

//算法十分直截了当：
//
//从左到右遍历数组，建立数组 left。
//
//从右到左遍历数组，建立数组 right。
//
//建立输出数组 max(right[i], left[i + k - 1])，其中 i 取值范围为 (0, n - k + 1)。


// 复杂度分析
//
//时间复杂度：{O}(N)O(N)，我们对长度为 N 的数组处理了 3次。
//
//空间复杂度：{O}(N)O(N)，用于存储长度为 N 的 left 和 right 数组，以及长度为 N - k + 1的输出数组。

func maxSlidingWindow(nums []int,k int) []int {
    // 边界条件
    if k < 1 || k ==0 ||nums == nil   {
        return nil
    }
    if k < 2  {
        return nums
    }
    res := []int{}
    length := len(nums)
    // 创建left、right 数组
    left := make([]int,length)
    right := make([]int,length)

    // 从左到右，生成left数组
    for i:=0;i<length;i++ {
        if i % k == 0 {
            left[i] = nums[i]
        } else {
            left[i] = MaxInt(left[i-1],nums[i])
        }
    }
    // 从右到左，生成right数组
    right[length-1] = nums[length-1]

    for i:= length-2;i>=0;i-- {
        if (i+1) %k == 0 {
            right[i] = nums[i]
        } else {
            right[i] = MaxInt(right[i+1],nums[i])
        }
    }
    // 从0到length-k+1,进行比较left[j],right[i]
    for i := 0;i<length-k+1;i++{
        res =append(res,MaxInt(left[i+k-1],right[i]))
    }
    return res
}
func MaxInt(i,j int) int {
    if i < j {
        return j
    }
    return i
}






