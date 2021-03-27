package main

import "container/heap"

/**
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

例如，
[2,3,4]的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。

示例 1：
输入：
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
输出：[null,null,null,1.50000,null,2.00000]

示例 2：
输入：
["MedianFinder","addNum","findMedian","addNum","findMedian"]
[[],[2],[],[3],[]]
输出：[null,null,2.00000,null,2.50000]

限制：
最多会对addNum、findMedian 进行50000次调用。
*/

type Heap []int

func (m *Heap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *Heap) Len() int {
	return len(*m)
}

func (m *Heap) Pop() (v interface{}) {
	*m, v = (*m)[:len(*m)-1], (*m)[len(*m)-1]
	return
}

func (m *Heap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

type minHeap struct {
	Heap
}

func (m *minHeap) Less(i, j int) bool {
	return m.Heap[i] < m.Heap[j]
}

type maxHeap struct {
	Heap
}

func (m *maxHeap) Less(i, j int) bool {
	return m.Heap[i] > m.Heap[j]
}

type MedianFinder struct {
	RightMin *minHeap
	LeftMax  *maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := new(MedianFinder)
	m.RightMin = new(minHeap)
	m.LeftMax = new(maxHeap)
	heap.Init(m.LeftMax)
	heap.Init(m.RightMin)
	return *m
}

func (this *MedianFinder) AddNum(num int) {
	// 压入数据时有两种情况；
	// （1）原有数据为偶数个，则有 leftheap.len() == rightheap.len()；
	//     此时将数据放入leftheap， 高明之处：需先将数据放在rightheap，再从rightheap中pop出一个元素，将其放入leftheap；
	//     上面做法的高明之处 在于 省去了判断 num 与 rightheap最小值 及 leftheap最大值 的比较，通过比较判断num应该放在哪个heap，
	//     然后再根据左右heap的长度，对左右heap进行调整
	//  (2) 原有数据为奇数个，则有 leftheap.len() == rightheap.len() + 1  (通过step 1 来保证当为奇数时，left总比right 大 1，这样可以简化判断)
	//     类似step 1 需先将数据放在leftheap，再从leftheap中pop出一个元素，将其放入rightheap中
	if this.LeftMax.Len() == this.RightMin.Len() {
		heap.Push(this.RightMin, num)
		heap.Push(this.LeftMax, heap.Pop(this.RightMin))
	} else {
		heap.Push(this.LeftMax, num)
		heap.Push(this.RightMin, heap.Pop(this.LeftMax))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.RightMin.Len()+this.LeftMax.Len())%2 == 0 {
		//此处根据heap的实现原理，直接获取heap的第一个元素，即为大顶堆的最大值
		return float64(this.LeftMax.Heap[0]+this.RightMin.Heap[0]) / 2
	} else {
		return float64(this.LeftMax.Heap[0])
	}
}

/**
		JAVA
class MedianFinder {
    Queue<Integer> A, B;
    public MedianFinder() {
        A = new PriorityQueue<>(); // 小顶堆，保存较大的一半
        B = new PriorityQueue<>((x, y) -> (y - x)); // 大顶堆，保存较小的一半
    }
    public void addNum(int num) {
        if(A.size() != B.size()) {
            A.add(num);
            B.add(A.poll());
        } else {
            B.add(num);
            A.add(B.poll());
        }
    }
    public double findMedian() {
        return A.size() != B.size() ? A.peek() : (A.peek() + B.peek()) / 2.0;
    }
}
*/

func main() {

}
