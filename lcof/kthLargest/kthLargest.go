package main

/**
给定一棵二叉搜索树，请找出其中第k大的节点。

示例 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
  2
输出: 4

示例 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4

限制：
1 ≤ k ≤ 二叉搜索树元素个数
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
迭代
*/
func kthLargest(root *TreeNode, k int) int {
	res := 0
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			res = root.Val
			break
		}
		root = root.Left
	}
	return res
}

/**
递归
*/
var num int
var val int

func kthLargest1(root *TreeNode, k int) int {
	num = k
	dfs(root)
	return val
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Right)
		num--
		if num == 0 {
			val = root.Val
			return
		}
		dfs(root.Left)
	}
}

/**
class Solution {
    public int kthLargest(TreeNode root, int k) {
        int res = 0;
        Deque<TreeNode> deque = new ArrayDeque<>();
        while (root != null || !deque.isEmpty()) {
            while (root != null) {
                deque.push(root);
                root = root.right;
            }
            root = deque.pop();

            if (--k == 0) {
                res = root.val;
                break;
            }

            root = root.left;
        }
        return res;
    }
}

class Solution {
    int k;
    int val;

    public int kthLargest(TreeNode root, int k) {
        this.k = k;
        dfs(root);
        return val;
    }

    private void dfs(TreeNode root) {
        if (root != null) {
            dfs(root.right);
            if (--k == 0) {
                val = root.val;
                return;
            }
            dfs(root.left);
        }
    }
}
*/
func main() {

}
