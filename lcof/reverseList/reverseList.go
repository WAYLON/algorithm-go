package main

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

限制：
0 <= 节点个数 <= 5000
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
方法1递归
*/
func reverseList(head *ListNode) *ListNode {
	return recur(head, nil)
}

func recur(cur *ListNode, pre *ListNode) *ListNode {
	if cur == nil { // 终止条件
		return pre
	}
	res := recur(cur.Next, cur) // 递归后继节点
	cur.Next = pre              // 修改节点引用指向
	return res                  // 返回反转链表的头节点
}

/**
方法2双指针
*/
func reverseList2(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func main() {

}
