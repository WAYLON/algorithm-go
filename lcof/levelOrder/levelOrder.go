package main

/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：
[3,9,20,15,7]

提示：
节点总数 <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
bfs
*/
func levelOrder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		//出队列
		r := queue[0]
		queue = queue[1:]

		res = append(res, r.Val)

		if r.Left != nil {
			queue = append(queue, r.Left)
		}
		if r.Right != nil {
			queue = append(queue, r.Right)
		}
	}
	return res
}

func main() {
	node := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	levelOrder(node)
}
