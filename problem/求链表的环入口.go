package preoblem

/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
解题思路： 快慢指针， 快指针的速度为2， 慢指针的速度为1。
* 设链表到环的入口距离为 L， 其快慢指针的相遇点距离链表的入口地址S。
* 在相遇时慢指针走的路程： L + S(为什么慢指针走的路程为L+S 而不是其他的呢)
* 快指针的走的路程为 2(L+S) = L + S + nR
* 那么我们可以的到 L + S = nR，   L = nR-S
* 入口在相遇点再走： nR-S 步， 那么一定会回到环的入口点。
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	p1, p2 := head, head

	for {
		p2 = p2.Next
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
		if p2 == nil {
			return nil
		}
		p1 = p1.Next
		if p1 == p2 {
			break
		}
	}

	p1 = head
	for p1 == p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
