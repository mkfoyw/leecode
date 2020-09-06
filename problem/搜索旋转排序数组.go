package preoblem

/* 搜索旋转排序数组
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

解题思路： 二分思想
旋转数组进行二分， 那么一定有一半是有序的， 另外一半是无序的(或许也是有序的)。 那么我们非常容易判断查找的值，
是否在有序数组里面， 如果没在有序数组里面， 那么一定是在另一半里面。 只需要这样查找就行。

还有一点： 注意二分的写法
*/

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[l] < nums[mid] {

			if nums[l] <= target && nums[mid] >= target {
				r = mid
			} else {
				l = mid + 1
			}

		} else {
			if nums[mid+1] <= target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}

	if nums[l] == target {
		return l
	} else {
		return -1
	}
}
