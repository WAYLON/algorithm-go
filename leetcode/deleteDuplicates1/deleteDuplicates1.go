package main

/**
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中没有重复出现的数字。
返回同样按升序排列的结果链表。

示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]

提示：
链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	for head != nil {
		//head 已经没有下一个节点 || head 有一下个节点，但是值与 head 不相同
		if head.Next == nil || head.Val != head.Next.Val {
			cur.Next = head
			cur = cur.Next
		}
		// 如果 head 与下一节点相同，跳过相同节点
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		head = head.Next
	}
	//防止出现[1,2,2]这种情况
	cur.Next = nil
	return dummyHead.Next
}

func main() {

}
