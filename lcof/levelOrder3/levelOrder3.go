package main

/**
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

例如:
给定二叉树:[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]

提示：
节点总数 <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
方法1 dfs
*/
var res = make([][]int, 0)

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, i int) {
	if root == nil {
		return
	}
	if len(res) == i {
		res = append(res, make([]int, 0))
	}
	if (i & 1) == 1 {
		res[i] = append([]int{root.Val}, res[i]...)
	} else {
		res[i] = append(res[i], root.Val)
	}
	dfs(root.Left, i+1)
	dfs(root.Right, i+1)
}

/**
方法2 bfs
*/
func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) != 0 {
		temp := make([]int, 0)
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if (len(res) & 1) == 1 {
				temp = append([]int{node.Val}, temp...)
			} else {
				temp = append(temp, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}

	return res
}

func main() {

}
