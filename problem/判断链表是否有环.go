package interview

/* 环形链表
给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-cycle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

解题思路： 快慢指针
一个指针跑的速度为1， 另一个指针跑的速度为2，
那么第一个指针的跑一圈， 那么第二个指针跑两圈，
最终会在出发点相遇。
*/

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1 := head
	p2 := head.Next
	for {
		if p2 != nil {
			p2 = p2.Next
		} else {
			return false
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			return false
		}

		p1 = p1.Next
		if p1 == p2 {
			return true
		}
	}
}
