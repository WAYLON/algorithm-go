package main

/**
请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个叶值序列 。
举个例子，如上图所示，给定一棵叶值序列为(6, 7, 4, 9, 8)的树。
如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是叶相似的。
如果给定的两个根结点分别为root1 和root2的树是叶相似的，则返回true；否则返回 false 。

示例 1：
输入：root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
输出：true

示例 2：
输入：root1 = [1], root2 = [1]
输出：true

示例 3：
输入：root1 = [1], root2 = [2]
输出：false

示例 4：
输入：root1 = [1,2], root2 = [2,2]
输出：true

示例 5：
输入：root1 = [1,2,3], root2 = [1,3,2]
输出：false


提示：
给定的两棵树可能会有1到 200个结点。
给定的两棵树上的值介于 0 到 200 之间。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
方法1：深度优先搜索dfs
*/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var l1 []int
	var l2 []int
	var list []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			list = append(list, root.Val)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root1)
	l1 = list
	list = []int{}
	dfs(root2)
	l2 = list
	if len(l1) == len(l2) {
		for i := 0; i < len(l1); i++ {
			if l1[i] != l2[i] {
				return false
			}
		}
		return true
	}
	return false
}

/**
方法2：栈迭代
*/
func leafSimilar2(root1 *TreeNode, root2 *TreeNode) bool {
	var l1 []int
	var l2 []int
	var list []int
	var process func(root *TreeNode)
	process = func(root *TreeNode) {
		var stack []*TreeNode
		for len(stack) != 0 || root != nil {
			for root != nil {
				stack = append(stack, root)
				root = root.Left
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Left == nil && root.Right == nil {
				list = append(list, root.Val)
			}
			root = root.Right
		}
	}

	process(root1)
	l1 = list
	list = []int{}
	process(root2)
	l2 = list
	if len(l1) == len(l2) {
		for i := 0; i < len(l1); i++ {
			if l1[i] != l2[i] {
				return false
			}
		}
		return true
	}
	return false
}

/**
class Solution {
    public boolean leafSimilar(TreeNode t1, TreeNode t2) {
        List<Integer> l1 = new ArrayList<>(), l2 = new ArrayList<>();
        dfs(t1, l1);
        dfs(t2, l2);
        if (l1.size() == l2.size()) {
            for (int i = 0; i < l1.size(); i++) {
                if (!l1.get(i).equals(l2.get(i))) return false;
            }
            return true;
        }
        return false;
    }
    void dfs(TreeNode root, List<Integer> list) {
        if (root == null) return;
        if (root.left == null && root.right == null) {
            list.add(root.val);
            return;
        }
        dfs(root.left, list);
        dfs(root.right, list);
    }
}

class Solution {
    public boolean leafSimilar(TreeNode t1, TreeNode t2) {
        List<Integer> l1 = new ArrayList<>(), l2 = new ArrayList<>();
        process(t1, l1);
        process(t2, l2);
        if (l1.size() == l2.size()) {
            for (int i = 0; i < l1.size(); i++) {
                if (!l1.get(i).equals(l2.get(i))) return false;
            }
            return true;
        }
        return false;
    }
    void process(TreeNode root, List<Integer> list) {
        Deque<TreeNode> d = new ArrayDeque<>();
        while (root != null || !d.isEmpty()) {
            while (root != null) {
                d.addLast(root);
                root = root.left;
            }
            root = d.pollLast();
            if (root.left == null && root.right == null) list.add(root.val);
            root = root.right;
        }
    }
}
*/

func main() {
	leafSimilar(&TreeNode{Val: 1}, &TreeNode{Val: 2})
}
