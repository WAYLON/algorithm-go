package main

/**
给定二叉搜索树的根结点root，返回值位于范围 [low, high] 之间的所有结点的值的和。

示例 1：
输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32

示例 2：
输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
输出：23

提示：
树中节点数目在范围 [1, 2 * 104] 内
1 <= Node.val <= 105
1 <= low <= high <= 105
所有 Node.val 互不相同
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归
*/
func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if low <= root.Val && root.Val <= high {
			res += root.Val
		} else if root.Val > high {
			return
		}

		dfs(root.Right)
	}
	dfs(root)
	return res
}

/**
迭代
*/
func rangeSumBST2(root *TreeNode, low int, high int) int {
	res := 0
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if low <= root.Val && root.Val <= high {
			res += root.Val
		} else if root.Val > high {
			break
		}
		root = root.Right
	}
	return res
}

/**
class Solution {
    public int rangeSumBST(TreeNode root, int L, int R) {
        if (root == null) {
            return 0;
        }
        if (root.val < L) {
            return rangeSumBST(root.right, L, R);
        }
        if (root.val > R) {
            return rangeSumBST(root.left, L, R);
        }
        return root.val + rangeSumBST(root.left, L, R) + rangeSumBST(root.right, L, R);
    }
}

class Solution {
    int low;
    int high;
    int res;

    public int rangeSumBST(TreeNode root, int low, int high) {
        this.high = high;
        this.low = low;
        dfs(root);
        return res;
    }

    private void dfs(TreeNode root) {
        if (root != null) {
            dfs(root.left);
            if (low <= root.val && root.val <= high) {
                res += root.val;
            }
            if (root.val > high) {
                return;
            }
            dfs(root.right);
        }
    }
}

*/
func main() {

}
