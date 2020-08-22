/*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
通过次数140,335提交次数210,993

来源：https://leetcode-cn.com/problems/swap-nodes-in-pairs/

思路： 每次前进两步， 然后交换相邻的节点就行。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	newhead := new(ListNode)
	newhead.Next = head
	p := newhead

	for p != nil && p.Next != nil && p.Next.Next != nil {
		p1 := p.Next
		p2 := p.Next.Next

		p.Next = p2
		p1.Next = p2.Next
		p.Next.Next = p1

		p = p.Next.Next
	}
	return newhead.Next
}