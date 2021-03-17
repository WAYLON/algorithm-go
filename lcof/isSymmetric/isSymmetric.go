package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
例如，二叉树[1,2,2,3,4,4,3] 是对称的。

  1
 / \
 2  2
/ \ / \
3 4 4 3
但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:
  1
 / \
 2  2
 \  \
 3  3

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false

限制：
0 <= 节点个数 <= 1000
*/

/**
方法1 递归
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

/**
方法2 bfs
*/
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	queue = append(queue, root.Left, root.Right) // 第一次先将root的左右append进去（保证循环有效性）
	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]
		if left == nil && right == nil { // 左右为空（对称），则继续
			continue
		} else if left == nil || right == nil { // 左右其中一个为空，则返回false
			return false
		} else if left.Val != right.Val { // 左右值不相等返回false
			return false
		}
		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}
	return true
}

func main() {

}
