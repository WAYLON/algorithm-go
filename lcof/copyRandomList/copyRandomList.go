package main

/**
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

示例 1：
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

示例 2：
输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]

示例 3：
输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]

示例 4：
输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。

提示：
-10000 <= Node.val <= 10000
Node.random为空（null）或指向链表中的节点。
节点数目不超过 1000 。

*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
方法1 map映射
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	cur := head
	//1.建立对应映射关系
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	//2.循环建立关系 next及Random
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}

	return m[head]
}

/**
方法2 原地算法
*/
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	//1.复制链表
	for cur != nil {
		temp := &Node{Val: cur.Val}
		temp.Next = cur.Next
		cur.Next = temp
		cur = temp.Next
	}
	cur = head

	//2.建立 Random关系
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	//3.拆分链表
	res := head.Next
	cur = head.Next
	pre := head
	for cur.Next != nil {
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}

	//4.单独处理链表的尾结点
	pre.Next = nil
	return res
}
func main() {

}
