package main

/**
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：
  4
 /  \
 2   7
/ \  / \
1  3 6  9
镜像输出：
  4
 /  \
 7   2
/ \  / \
9  6 3 1

示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

限制：
0 <= 节点个数 <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
方法1 递归
*/
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}

/**
方法2 栈/队列
*/
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		lens := len(queue)
		for i := 0; i < lens; i++ {
			var node *TreeNode = queue[i]
			var temp *TreeNode = queue[i].Left
			node.Left = node.Right
			node.Right = temp
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 此处表示只保存新一轮的节点列表
		queue = queue[lens:]
	}

	return root
}

func main() {

}
