package main

import "container/list"

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：
输入：head = [1,3,2]
输出：[2,3,1]

限制：
0 <= 链表长度 <= 10000
*/

/**
方法1：栈
*/
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	res := list.New()
	for head != nil {
		res.PushFront(head.Val)
		head = head.Next
	}

	ret := []int{}
	for e := res.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(int))
	}

	return ret
}

/**
不定长切片
*/
func reversePrint1(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}

	count := 0
	newHead := head
	for head != nil {
		count++
		head = head.Next
	}

	res := make([]int, count)
	i := 0
	for newHead != nil {
		res[count-i-1] = newHead.Val
		i++
		newHead = newHead.Next
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
