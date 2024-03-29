package main

//80
//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
// 示例 1:
//
// 给定 nums = [1,1,1,2,2,3],
//
//函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。
//
//你不需要考虑数组中超出新长度后面的元素。
//
// 示例 2:
//
// 给定 nums = [0,0,1,1,1,1,2,3,3],
//
//函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。
//
//你不需要考虑数组中超出新长度后面的元素。
//
//
// 说明:
//
// 为什么返回数值是整数，但输出的答案是数组呢?
//
// 请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
// 你可以想象内部操作如下:
//
// // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
//int len = removeDuplicates(nums);
//
//// 在函数里修改输入数组对于调用者是可见的。
//// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
//for (int i = 0; i < len; i++) {
//    print(nums[i]);
//}
// Related Topics 数组 双指针

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates2(nums []int) int {
	//第一个指针遍历的位置，第二个指针有用元素的位置
	if len(nums) <= 2 {
		return len(nums)
	}
	tmp := nums[0]
	tmp1 := nums[1]
	idx := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] == tmp && nums[i] == tmp1 {
			tmp = tmp1
			tmp1 = nums[i]
			nums[i] = 0
		} else {
			tmp = tmp1
			tmp1 = nums[i]
			nums[i] = 0
			nums[idx] = tmp1
			idx++
		}
	}
	return idx

}

//leetcode submit region end(Prohibit modification and deletion)
