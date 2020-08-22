package interview

/*
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:

输入: [1,3,4,2,2]
输出: 2
示例 2:

输入: [3,1,3,4,2]
输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-the-duplicate-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

解题思路：

1. hash 表。 统计重复次数我们可以采用hash表来计算。
2. 二分查找法
   解的值域范围为： [1, n], 因此我们可以进行二分搜索值域： mid = (l + r)
   我们可以统计小于等于mid数cnt[mid]:
	   如果cnt[mid] <= mid， 那么重复的数字可能存在结果为[mid+1, r]
	   如果cnt[mid] > mid, 那么重复的数字可能存在的结果为[l, mid]
*/

//二分法
func findDuplicate(nums []int) int {
	l, r := 1, len(nums)
	for l != r {
		mid := (l + r) / 2
		cnt := 0

		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

//这也可以看做是一个链表
