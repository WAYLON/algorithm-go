package main

/**
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

提示：
各函数的调用总次数不超过 20000 次
*/

type MinStack struct {
	stack1 []int
	stack2 []int
}

func Constructor() MinStack {
	return MinStack{stack1: make([]int, 0), stack2: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stack1 = append(this.stack1, x)
	if len(this.stack2) == 0 || this.stack2[len(this.stack2)-1] >= x {
		this.stack2 = append(this.stack2, x)
	}
}

func (this *MinStack) Pop() {
	temp := this.stack1[len(this.stack1)-1]
	this.stack1 = this.stack1[:len(this.stack1)-1]
	if temp == this.stack2[len(this.stack2)-1] {
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack1) == 0 {
		return -1
	}
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) Min() int {
	if len(this.stack2) == 0 {
		return -1
	}
	return this.stack2[len(this.stack2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

/**
Java版
class MinStack {
    private Deque<Integer> stack1;

    private Deque<Integer> stack2;
    public MinStack() {
        stack1 = new ArrayDeque<>();
        stack2 = new ArrayDeque<>();
    }

    public void push(int x) {
        stack1.push(x);
        if (stack2.isEmpty() || stack2.peekFirst() >= x) {
            stack2.push(x);
        }
    }

    public void pop() {
        if (stack1.pop().equals(stack2.peekFirst())) {
            stack2.pop();
        }
    }

    public int top() {
        return stack1.peekFirst();
    }

    public int min() {
        return stack2.peekFirst();
    }
}
*/
func main() {

}
