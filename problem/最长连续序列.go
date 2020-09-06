package preoblem

/*
128. 最长连续序列
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
题目来源： https://leetcode-cn.com/problems/longest-consecutive-sequence/

解题思路：
我们可以枚举X连续的数，  ... X-3, X-2, X-1, X, X+1, X+2, ....
我们发现我们只需要从最小的枚举。
我们可以采用 hash 表来查找
*/

import "math"

// Hash 算法
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int64]int)
	for _, num := range nums {
		m[int64(num)] = 1
	}

	res := 1
	for key := range m {

		if key != math.MinInt64 {
			// 这是避免最小值
			if _, ok := m[key-1]; ok {
				continue
			}
		}

		cnt := 1
		for key != math.MaxInt64 {
			if _, ok := m[key+1]; ok {
				cnt++
			} else {
				break
			}
			key = key + 1
		}

		if res < cnt {
			res = cnt
		}
	}
	return res
}
