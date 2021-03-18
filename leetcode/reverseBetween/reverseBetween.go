package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
给你单链表的头节点head 和两个整数left 和 right ，其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
示例 1：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：
输入：head = [5], left = 1, right = 1
输出：[5]

提示：
链表中节点数目为 n
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n

进阶： 你可以使用一趟扫描完成反转吗？
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//避免讨论头节点
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	//从假头遍历到左节点的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	//左节点
	leftNode := pre.Next
	rightNode := pre
	//再从左节点遍历到右节点
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	//剩余的节点
	cur := rightNode.Next

	//切割链表
	pre.Next = nil
	rightNode.Next = nil
	//反转链表
	reverseLinkedList(leftNode)

	pre.Next = rightNode
	leftNode.Next = cur

	return dummyNode.Next
}

func reverseLinkedList(node *ListNode) {
	pre := &ListNode{}
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func main() {

}
