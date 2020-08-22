package interview

/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]

提示：

节点总数 <= 1000
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处

解题思路：

主要BS 加上记录当前的节点的层数
代码注意点是如何记录每一层的节点数
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		//下一层的节点
		tmp := []*TreeNode{}
		tres := make([]int, len(queue), len(queue))
		for i := 0; i < len(queue); i++ {
			if level%2 == 0 {
				tres[i] = queue[i].Val
			} else {
				tres[len(queue)-i-1] = queue[i].Val
			}

			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}

			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		level++
		queue = tmp
		res = append(res, tres)
	}
	return res
}
