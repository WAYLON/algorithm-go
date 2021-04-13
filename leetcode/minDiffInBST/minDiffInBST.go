package main

import "math"

/**
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/ 相同

示例 1：
输入：root = [4,2,6,1,3]
输出：1

示例 2：
输入：root = [1,0,48,null,null,12,49]
输出：1

提示：
树中节点数目在范围 [2, 100] 内
0 <= Node.val <= 105
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
方法1 栈迭代
*/
func minDiffInBST(root *TreeNode) int {
	//前一个节点值
	pre := -1
	//返回结果
	res := math.MaxInt32
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != -1 && res > root.Val-pre {
			res = root.Val - pre
		}
		pre = root.Val
		root = root.Right
	}
	return res
}

/**
方法2 递归
*/
func minDiffInBST2(root *TreeNode) int {
	//前一个节点值
	pre := -1
	//返回结果
	res := math.MaxInt32
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != -1 && root.Val-pre < res {
			res = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return res
}

/**
方法3 莫里斯法
*/
func minDiffInBST3(root *TreeNode) int {
	//前一个节点值
	preVal := -1
	//返回结果
	res := math.MaxInt32

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
			if preVal != -1 && root.Val-preVal < res {
				res = root.Val - preVal
			}
			preVal = root.Val
			root = root.Right
		}
	}
	return res
}

/**
Java版
//栈解法
class Solution {
    public class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode() {
        }
        TreeNode(int val) {
            this.val = val;
        }
        TreeNode(int val, TreeNode left, TreeNode right) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    public int minDiffInBST(TreeNode root) {
        //记录前一个节点
        TreeNode pre = null;
        //最小值结果
        int res = Integer.MAX_VALUE;
        //定义栈数据结构
        Deque<TreeNode> stack = new ArrayDeque<>();
        while (root != null || !stack.isEmpty()) {
            //一直压栈至最底部
            while (root != null) {
                stack.push(root);
                root = root.left;
            }
            root = stack.pop();
            if (pre != null) {
                //记录最小
                res = Math.min(root.val - pre.val, res);
            }
            //记录前一个节点
            pre = root;
            //将节点指向右节点
            root = root.right;

        }
        return res;
    }
}

//递归解法
class Solution {
    //记录前一个节点
    private TreeNode pre = null;
    //最小值结果
    private int res = Integer.MAX_VALUE;

    public int minDiffInBST(TreeNode root) {
        dfs(root);
        return res;
    }

    private void dfs(TreeNode root) {
        if (root == null) return;
        dfs(root.left);
        if (pre != null) {
            res = Math.min(root.val - pre.val, res);
        }
        pre = root;
        dfs(root.right);
    }
}

//莫里斯法
class Solution {
    public int minDiffInBST(TreeNode root) {
        //记录前一个节点
        TreeNode pre = null;
        //最小值结果
        int res = Integer.MAX_VALUE;

        //指针
        TreeNode cur;

        while (root != null) {
            //如果当前节点的左节点不为空 则将当前节点连带右子树所有节点一起挂在左节点最右子树上
            if (root.left != null) {
                cur = root.left;
                while (cur.right != null) {
                    cur = cur.right;
                }
                cur.right = root;
                //将root指向左节点并将左节点置为空
                TreeNode temp = root.left;
                root.left = null;
                root = temp;
            } else {
                if (pre != null) {
                    res = Math.min(res, root.val - pre.val);
                }
                pre = root;
                root = root.right;
            }
        }
        return res;
    }
}

*/
func main() {

}
