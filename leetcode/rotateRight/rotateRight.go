package main

/**
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动k个位置。

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

示例 2：
输入：head = [0,1,2], k = 4
输出：[2,0,1]

提示：
链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	//特判
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	cur := head
	n := 1

	//统计链表的长度 n
	for cur.Next != nil {
		n++
		cur = cur.Next
	}

	//计算真正需要移动的步数
	k = k % n

	//特判
	if k == 0 {
		return head
	}
	//1.先成环、尾结点链接头结点
	cur.Next = head

	//2.找到新节点的尾结点
	temp := head
	for i := 0; i < n-1-k; i++ {
		temp = temp.Next
	}

	//3.找到新节点的头结点
	newHead := temp.Next

	//4.断开
	temp.Next = nil

	return newHead
}

func main() {

}
