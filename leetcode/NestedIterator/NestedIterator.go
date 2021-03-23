package main

/**
给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。

示例 1:
输入: [[1,1],2,[1,1]]
输出: [1,1,2,1,1]
解释: 通过重复调用next 直到hasNext 返回 false，next返回的元素的顺序应该是: [1,1,2,1,1]。

示例 2:
输入: [1,[4,[6]]]
输出: [1,4,6]
解释: 通过重复调用next直到hasNext 返回 false，next返回的元素的顺序应该是: [1,4,6]。

*/

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return -1
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

/**
方法1 栈
*/
type NestedIterator struct {
	Stack []*NestedInteger // 栈是用来暂存元素的
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var stack []*NestedInteger
	for i := len(nestedList) - 1; i >= 0; i-- { // 将元素从后开始一个个入栈
		stack = append(stack, nestedList[i]) // 从后是因为后进先出，位于前面的要先出
	}
	return &NestedIterator{Stack: stack} // 初始化并返回这个迭代器
}

func (this *NestedIterator) Next() int {
	this.stackTop2Integer()                     // 将栈顶元素转成integer
	top := this.Stack[len(this.Stack)-1]        // 获取栈顶
	this.Stack = this.Stack[:len(this.Stack)-1] // pop掉栈顶
	return top.GetInteger()                     // 返回integer
}

func (this *NestedIterator) HasNext() bool {
	this.stackTop2Integer()    // 将栈顶元素转成integer
	return len(this.Stack) > 0 // 如果stack没有空。说明能获取到integer
}

// 将栈顶元素变成integer（通过将list栈顶pop出来，再将它里面的元素一个个再放进去）
func (this *NestedIterator) stackTop2Integer() {
	for len(this.Stack) > 0 { // 直到栈空了
		top := this.Stack[len(this.Stack)-1] // 获取栈顶
		if top.IsInteger() {                 // 栈顶已经是integer，什么都不做
			return
		}
		this.Stack = this.Stack[:len(this.Stack)-1] // 栈顶是list，弹出栈顶
		list := top.GetList()                       // 获取list
		for i := len(list) - 1; i >= 0; i-- {       // 从后开始遍历list
			this.Stack = append(this.Stack, list[i]) // 将list的元素推入栈
		}
	}
}

/**
方法2 迭代

type NextedIterator struct { // 扁平迭代器
	index int   // index是当前Next调用应该返回的第index个数
	nums  []int // 存放展平后的数字
}

func Constructor(nestedList []*NestedInteger) *NextedIterator {
	nums := []int{}
	for _, node := range nestedList { // 遍历nestedList
		flatten(node, &nums)
	}
	return &NextedIterator{index: 0, nums: nums} // 创建并返回扁平迭代器
}

func flatten(node *NestedInteger, nums *[]int) {
	if node.IsInteger() { // 如果是数字
		*nums = append(*nums, node.GetInteger()) // 将数字推入res
	} else { // 是nestedList
		for _, child := range node.GetList() { // 遍历nestedList，child是当前子元素
			flatten(child, nums)
		}
	}
}

func (this *NextedIterator) Next() int { // 一个个获取展平后的数组的元素
	val := this.nums[this.index] // 获取元素
	this.index++                 // 更新index
	return val
}

func (this *NextedIterator) HasNext() bool { // 是否还能获取到元素
	return this.index <= len(this.nums)-1 // 取决于index是否越了nums的界
}

*/

func main() {

}
