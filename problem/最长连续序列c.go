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

解题思路：并查集

我们发现只需要 把 X 与 X-1 进行合并就行。

这儿有一个初始化的亮点： mp[X]=X
*/

import "math"

var mp map[int64]int64
var sz map[int64]int

func longestConsecutive2(nums []int) int {
	mp = make(map[int64]int64)
	sz = make(map[int64]int)

	res := 0
	if len(nums) == 0 {
		res = 0
	} else {
		res = 1
	}

	for _, item := range nums {
		mp[int64(item)] = int64(item)
		sz[int64(item)] = 1
	}

	for _, key := range mp {
		if key != math.MinInt64 {
			if _, ok := mp[key-1]; ok {
				res = max(res, merge(key-1, key))
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func find(x int64) int64 {
	if mp[x] == x {
		return x
	}
	mp[x] = find(mp[x])
	return mp[x]
}

func merge(x, y int64) int {
	p1 := find(x)
	p2 := find(y)
	if p1 == p2 {
		return sz[p1]
	}

	mp[p2] = p1
	sz[p1] += sz[p2]
	return sz[p1]
}
