package main

/**
在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。

示例 1：
输入：root = [1,2,3,4], x = 4, y = 3
输出：false

示例 2：
输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
输出：true

示例 3：
输入：root = [1,2,3,null,4], x = 2, y = 3
输出：false

提示：
二叉树的节点数介于2 到100之间。
每个节点的值都是唯一的、范围为1 到100的整数。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	xpar := 0
	xdep := 0
	ypar := 0
	ydep := 0

	var dfs func(node *TreeNode, dep int, x int, y int, par int)
	dfs = func(node *TreeNode, dep int, x int, y int, par int) {
		if node == nil {
			return
		}
		if node.Val == x {
			xpar = par
			xdep = dep
		} else if node.Val == y {
			ypar = par
			ydep = dep
		} else {
			dfs(node.Left, dep+1, x, y, node.Val)
			dfs(node.Right, dep+1, x, y, node.Val)
		}
	}

	dfs(root.Left, 1, x, y, root.Val)
	dfs(root.Right, 1, x, y, root.Val)
	return (xpar != ypar) && (xdep == ydep)
}

/**
class Solution {
    int xpar, xdep, ypar, ydep;

    public boolean isCousins(TreeNode root, int x, int y) {
        dfs(root.left, 1, x, y, root.val);
        dfs(root.right, 1, x, y, root.val);
        return (xpar != ypar) && (xdep == ydep);
    }

    public void dfs(TreeNode node, int dep, int x, int y, int par) {
        if (node == null) {
            return;
        }
        if (node.val == x) {
            xpar = par;
            xdep = dep;
        } else if (node.val == y) {
            ypar = par;
            ydep = dep;
        } else {
            dfs(node.left, dep+1, x, y, node.val);
            dfs(node.right, dep+1, x, y, node.val);
        }
    }
}

*/
func main() {

}
