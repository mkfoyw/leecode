package interview 

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:

输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

解题思路：
	接上一个题， 搜索旋转数组I， 就是 nums[i] == nums[mid] 我们不能判断 nums[i, mid] 是否有序
*/


func search(nums []int, target int) bool {
    if len(nums) == 0{
        return false
    }
    i, j := 0, len(nums)-1 

    for i != j {
        mid := (i + j)/2 
        if nums[mid] == target{
            return true
        }
        // 这是避免 111111113  3 旋转后 111311111
        if nums[i] == nums[mid]{
            i ++
            continue
        }

        if nums[i] <= nums[mid]{
            if target >= nums[i] && target <= nums[mid]{
                j = mid 
            }else{
                i = mid + 1
            }
        }else{
            if target >= nums[mid] && target <= nums[j]{
                i = mid
            }else{
                j = mid -1
            }
        }
    }

    if nums[i] == target{
        return true
    }else{
        return false
    }
}