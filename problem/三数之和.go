/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出
*/

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	n := len(nums)

	for first := 0; first < n; first++ {
		//保证不枚举同一个数
		if first != 0 && nums[first] == nums[first-1] {
			continue
		}

		targt := -1 * nums[first]

		for second := first + 1; second < n; second++ {
			//保证不枚举同一个数
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			third := n - 1
			//保证第三个指针在第二个的左边
			for third > second && nums[second]+nums[third] > targt {
				third--
			}
			// 这是一个优化
			if second == third {
				break
			}

			if nums[second]+nums[third] == targt {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}