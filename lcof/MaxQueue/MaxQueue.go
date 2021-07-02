package main

/**
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。
若队列为空，pop_front 和 max_value需要返回 -1

示例 1：
输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出:[null,null,null,2,1,2]

示例 2：
输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出:[null,-1,-1]

限制：
1 <= push_back,pop_front,max_value的总操作数<= 10000
1 <= value <= 10^5
*/

/**
5 1 3
push : 移除deque末尾比当前元素小的元素
queue: 5, 5 1, 5 1 3
deque: 5, 5 1, 5 3

pop : 如果移除元素等于deque头元素则移除
queue:  1 3, 3,
deque:  3,   3,
*/
type MaxQueue struct {
	queue []int
	//维护一个递减队列
	deque []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue: make([]int, 0),
		deque: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)

	for len(this.deque) != 0 && this.deque[len(this.deque)-1] < value {
		this.deque = this.deque[:len(this.deque)-1]
	}

	this.deque = append(this.deque, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	v := this.queue[0]
	this.queue = this.queue[1:]

	if this.Max_value() == v {
		this.deque = this.deque[1:]
	}

	return v
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

/**
java版

class MaxQueue {
    Queue<Integer> queue;
    Deque<Integer> deque;

    public MaxQueue() {
        queue = new LinkedList<>();
        deque = new LinkedList<>();
    }

    public int max_value() {
        return deque.isEmpty() ? -1 : deque.peekFirst();
    }

    public void push_back(int value) {
        queue.offer(value);
        while (!deque.isEmpty() && deque.peekLast() < value)
            deque.pollLast();
        deque.offerLast(value);
    }

    public int pop_front() {
        if (queue.isEmpty()) return -1;
        if (queue.peek().equals(deque.peekFirst()))
            deque.pollFirst();
        return queue.poll();
    }
}
*/
func main() {

}
