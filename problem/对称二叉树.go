/*
	101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

解题思路：
对称二叉树的特点是： 可以看成两个二叉树， 左边二叉树左子树 和 右边的二叉树的右子树相同。
这样直接遍历就行了。


*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return solve(root.Left, root.Right)
}

func solve(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil && r != nil {
		return false
	}

	if l != nil && r == nil {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	if solve(l.Left, r.Right) != true {
		return false
	}

	if solve(l.Right, r.Left) != true {
		return false
	}
	return true
}

//使用队列解决。
func solve(cur *TreeNode) bool {
	if cur == nil {
		return false
	}
	q := []*TreeNode{}

	q = append(q, cur.Left, cur.Right)
	for len(q) > 0 {
		l, r := q[0], q[1]
		q = q[2:]

		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		q = append(q, l.Left, r.Right)
		q = append(q, l.Right, r.Left)
	}
	return true
}