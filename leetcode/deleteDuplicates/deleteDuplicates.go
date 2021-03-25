package main

/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例1:
输入: 1->1->2
输出: 1->2

示例2:
输入: 1->1->2->3->3
输出: 1->2->3

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return dummyHead.Next
}

func main() {

}
