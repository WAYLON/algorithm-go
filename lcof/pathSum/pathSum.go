package main

/**
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

示例:
给定如下二叉树，以及目标和target = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

提示：
节点总数 <= 10000
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	res := [][]int{}
	temp := []int{}
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		temp = append(temp, root.Val)
		target -= root.Val

		if target == 0 && root.Right == nil && root.Left == nil {
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
		}

		if root.Left != nil {
			dfs(root.Left, target)
		}
		if root.Right != nil {
			dfs(root.Right, target)
		}
		temp = temp[:len(temp)-1]
	}
	dfs(root, target)
	return res
}

/**
Java版
class Solution {
    public List<List<Integer>> pathSum(TreeNode root, int target) {
        LinkedList<List<Integer>> res = new LinkedList<>();
        dfs(res, root, new LinkedList<>(), target);
        return res;
    }

    private void dfs(List<List<Integer>> res, TreeNode root, LinkedList<Integer> temp, int target) {
        if (root == null) {
            return;
        }
        temp.add(root.val);
        target -= root.val;

        if (target == 0 && root.right == null && root.left == null) {
            res.add(new ArrayList<>(temp));
        }

        if (root.left != null) {
            dfs(res, root.left, temp, target);
        }
        if (root.right != null) {
            dfs(res, root.right, temp, target);
        }

        temp.removeLast();
    }
}
*/
func main() {

}
