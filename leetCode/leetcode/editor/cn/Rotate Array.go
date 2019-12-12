package main

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
// 示例 1:
//
// 输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//
//
// 示例 2:
//
// 输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]
//
// 说明:
//
//
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 要求使用空间复杂度为 O(1) 的 原地 算法。
//
// Related Topics 数组

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int) {
	if k == len(nums) || len(nums) == 1 {
		return
	}
	e := k % len(nums)
	if e <= len(nums)/2 {
		for j := 0; j < e; j++ {
			c := nums[len(nums)-1]
			for i := len(nums) - 2; i >= 0; i-- {
				nums[i+1] = nums[i]
			}
			nums[0] = c
		}
	} else {
		for j := 0; j < len(nums)-e; j++ {
			c := nums[0]
			for i := 1; i <= len(nums)-1; i++ {
				nums[i-1] = nums[i]
			}
			nums[len(nums)-1] = c
		}
	}
}

func rotate2(nums []int, k int) {
	if len(nums) == 1 || k == 0 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	var tmp []int
	for i := len(nums) - k; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}
	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < len(tmp); i++ {
		nums[i] = tmp[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
