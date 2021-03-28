package main

/**
给定一个二叉树的根节点 root ，返回它的 中序遍历。

示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

示例 4：
输入：root = [1,2]
输出：[2,1]

示例 5：
输入：root = [1,null,2]
输出：[1,2]

提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

进阶:递归算法很简单，你可以通过迭代算法完成吗？

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
方法1 递归
*/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	//闭包函数必须先定义
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

/**
方法2 栈迭代
*/
func inorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

/**
方法3 莫里斯法
*/
func inorderTraversal3(root *TreeNode) (res []int) {
	var pre *TreeNode
	for root != nil {
		//如果左节点不为空，就将当前节点连带右子树全部挂到
		//左节点的最右子树下面
		if root.Left != nil {
			pre = root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = root
			//将root指向root的left
			tmp := root
			root = root.Left
			tmp.Left = nil
			//左子树为空，则打印这个节点，并向右边遍历
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return
}

func main() {

}
