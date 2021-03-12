package main

/**
序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如"1,,3" 。

示例 1:
输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
输出: true

示例2:
输入: "1,#"
输出: false

示例 3:
输入: "9,#,#,1"
输出: false
*/

/**
满二叉树叶子节点 = 非叶子节点 + 1
*/
func isValidSerialization(preorder string) bool {
	n := len(preorder)
	stack := 1
	i := 0
	for i < n {
		if stack == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			stack--
			i++
		} else {
			// 读一个数字 列如 92
			for i < n && preorder[i] != ',' {
				i++
			}
			stack++
		}

	}

	return stack == 0
}

func main() {
	isValidSerialization("9,#,92,#,#")
}
