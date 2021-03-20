package main

/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。
参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3

示例 1：
输入: [1,6,3,2,5]
输出: false

示例 2：
输入: [1,3,2,6,5]
输出: true

提示：
数组长度 <= 1000
*/
func verifyPostorder(postorder []int) bool {
	return recur(postorder, 0, len(postorder)-1)
}

func recur(postorder []int, i int, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	q := p
	for postorder[p] > postorder[j] {
		p++
	}

	return p == j && recur(postorder, i, q-1) && recur(postorder, q, j-1)
}

func verifyPostorder2(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	//在二叉搜索树中，左子树的元素是都小于根元素，右子树都大于根元素
	//在后序遍历中，最后一个元素是根元素
	head := len(postorder) - 1
	for head != 0 {
		//popinter 统计符合二叉搜索树的后序遍历的节点数
		popinter := 0
		//从前面开始遍历，小于的当前根元素的值是左子树的，当找到第一个大于当前根元素的值，可以确定后半段的元素都应是在当前节点的右子树
		for postorder[popinter] < postorder[head] {
			popinter++
		}
		//如果后半段里面有小于根元素的值的元素，就说明这个不是二叉搜索树的后序遍历，跳出循环
		for postorder[popinter] > postorder[head] {
			popinter++
		}
		//popinter != head 或 popinter < head 说明该数组不是某二叉搜索树的后序遍历结果
		if popinter != head {
			return false
		}
		//进入下一个节点继续验证
		head--
	}
	return true
}

func main() {

}
