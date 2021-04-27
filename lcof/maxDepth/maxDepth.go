package main

/**
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

例如：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3 。

提示：
节点总数 <= 10000
注意：本题与主站 104题相同：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
迭代
*/
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		qlen := len(queue)
		for i := 0; i < qlen; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res++
		queue = queue[qlen:]
	}
	return res
}

/**
class Solution {
    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }

    public int maxDepth(TreeNode root) {
 		if (root == null) return 0;
        int res = 0;
        List<TreeNode> list = new ArrayList<>();
        list.add(root);
        List<TreeNode> temp;
        while (list.size() > 0) {
            temp = new ArrayList<>();
            for (TreeNode treeNode : list) {
                if (treeNode.left != null) {
                    temp.add(treeNode.left);
                }
                if (treeNode.right != null) {
                    temp.add(treeNode.right);
                }
            }
            res++;
            list = temp;
        }
        return res;
    }
}


class Solution {
    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;

        TreeNode(int x) {
            val = x;
        }
    }

    public int maxDepth(TreeNode root) {
        return root == null ? 0 : Math.max(maxDepth(root.left), maxDepth(root.right)) + 1;
    }
}
*/
func main() {

}
