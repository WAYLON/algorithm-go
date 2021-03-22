package main

/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如，给出
前序遍历 preorder =[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
限制：
0 <= 节点个数 <= 5000
*/
/**
preOrder = [1,2,4,7,3,5,6,8]
inOrder = [4,7,2,1,5,3,8,6]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var preOrder []int
var inMap = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	preOrder = preorder
	for i, v := range inorder {
		inMap[v] = i
	}
	return recursive(0, 0, len(inorder)-1)
}

func recursive(rootIndex int, inLeftIndex int, inRightIndex int) *TreeNode {
	if inLeftIndex > inRightIndex {
		return nil
	}
	node := &TreeNode{}
	node.Val = preOrder[rootIndex]
	//根结点在中序 中的位置
	i := inMap[preOrder[rootIndex]]
	//前序中 左子树根结点应为 rootIndex + 1
	node.Left = recursive(rootIndex+1, inLeftIndex, i-1)

	//前序中 右子树根结点应为 rootIndex + (左子树的长度 = 左子树的左边-右边 (i-1 - inLeftIndex +1) 。最后+1就是右子树的根了)
	node.Right = recursive(rootIndex+i-inLeftIndex+1, i+1, inRightIndex)
	return node
}

func main() {

}
