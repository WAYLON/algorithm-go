package main

/**
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
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
  [9,20],
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
	if len(res)-1 < i {
		res = append(res, make([]int, 0))
	}
	res[i] = append(res[i], root.Val)
	dfs(root.Left, i+1)
	dfs(root.Right, i+1)
}

/**
方法2 bfs
*/
func levelOrder2(root *TreeNode) [][]int {
	var res1 = make([][]int, 0)
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		temp := make([]int, 0)
		//牛逼plus 倒序防止queue改变大小
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res1 = append(res1, temp)
	}
	return res1
}

/**
//错误写法 不知道为啥不行
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	dfs(res,root, 0)
	return res
}

func dfs(res [][]int, root *TreeNode, i int) {
	if root == nil {
		return
	}
	if len(res)-1 < i {
		res = append(res, make([]int, 0))
	}
	res[i] = append(res[i], root.Val)
	dfs(res,root.Left, i+1)
	dfs(res,root.Right, i+1)
}
*/

func main() {
	node := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	levelOrder2(node)
}
