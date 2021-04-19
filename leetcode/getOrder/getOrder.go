package main

import (
	"container/heap"
	"sort"
)

/**
给你一个二维数组 tasks ，用于表示 n 项从 0 到 n - 1 编号的任务。其中 tasks[i] = [enqueueTimei, processingTimei]
意味着第 i 项任务将会于 enqueueTimei 时进入任务队列，需要 processingTimei 的时长完成执行。
现有一个单线程 CPU ，同一时间只能执行 最多一项 任务，该 CPU 将会按照下述方式运行：
如果 CPU 空闲，且任务队列中没有需要执行的任务，则 CPU 保持空闲状态。
如果 CPU 空闲，但任务队列中有需要执行的任务，则 CPU 将会选择 执行时间最短 的任务开始执行。如果多个任务具有同样的最短执行时间，则选择下标最小的任务开始执行。
一旦某项任务开始执行，CPU 在 执行完整个任务 前都不会停止。
CPU 可以在完成一项任务后，立即开始执行一项新任务。
返回 CPU 处理任务的顺序。

示例 1：
输入：tasks = [[1,2],[2,4],[3,2],[4,1]]
输出：[0,2,3,1]
解释：事件按下述流程运行：
- time = 1 ，任务 0 进入任务队列，可执行任务项 = {0}
- 同样在 time = 1 ，空闲状态的 CPU 开始执行任务 0 ，可执行任务项 = {}
- time = 2 ，任务 1 进入任务队列，可执行任务项 = {1}
- time = 3 ，任务 2 进入任务队列，可执行任务项 = {1, 2}
- 同样在 time = 3 ，CPU 完成任务 0 并开始执行队列中用时最短的任务 2 ，可执行任务项 = {1}
- time = 4 ，任务 3 进入任务队列，可执行任务项 = {1, 3}
- time = 5 ，CPU 完成任务 2 并开始执行队列中用时最短的任务 3 ，可执行任务项 = {1}
- time = 6 ，CPU 完成任务 3 并开始执行任务 1 ，可执行任务项 = {}
- time = 10 ，CPU 完成任务 1 并进入空闲状态

示例 2：
输入：tasks = [[7,10],[7,12],[7,5],[7,4],[7,2]]
输出：[4,3,2,0,1]
解释：事件按下述流程运行：
- time = 7 ，所有任务同时进入任务队列，可执行任务项  = {0,1,2,3,4}
- 同样在 time = 7 ，空闲状态的 CPU 开始执行任务 4 ，可执行任务项 = {0,1,2,3}
- time = 9 ，CPU 完成任务 4 并开始执行任务 3 ，可执行任务项 = {0,1,2}
- time = 13 ，CPU 完成任务 3 并开始执行任务 2 ，可执行任务项 = {0,1}
- time = 18 ，CPU 完成任务 2 并开始执行任务 0 ，可执行任务项 = {1}
- time = 28 ，CPU 完成任务 0 并开始执行任务 1 ，可执行任务项 = {}
- time = 40 ，CPU 完成任务 1 并进入空闲状态

提示：
tasks.length == n
1 <= n <= 105
1 <= enqueueTimei, processingTimei <= 109
*/
/**
创建一个新的二维数组，在保留题目提供的tasks数据的基础上，再把下标存进去；
把tasks按任务起始时间排序；
建小顶堆，堆顶元素是时长最小的任务，时长相同的按下标排序；
初始化当前时间now = 1，遍历tasks中的任务，当堆空或者now>=tasks[i][0](即任务开始时间)时，将任务加入堆中，之后再弹出堆顶元素，根据弹出的新任务需要执行的时间来更新当前时刻now
重复4直道所有任务都被遍历
依次弹出堆中所有元素
*/

type pair struct{ t, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.t < b.t || a.t == b.t && a.i < b.i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }

func getOrder(a [][]int) (ans []int) {
	for i := range a {
		a[i] = append(a[i], i)
	}
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	h := &hp{}
	for i, cur, n := 0, 0, len(a); i < n; {
		if h.Len() > 0 {
			p := h.pop()
			ans = append(ans, p.i)
			cur += p.t
		}
		if h.Len() == 0 && cur < a[i][0] {
			cur = a[i][0]
		}
		for ; i < n && a[i][0] <= cur; i++ {
			h.push(pair{a[i][1], a[i][2]})
		}
	}
	for h.Len() > 0 {
		ans = append(ans, h.pop().i)
	}
	return
}

/**
创建一个新的二维数组，在保留题目提供的tasks数据的基础上，再把下标存进去；
把tasks按任务起始时间排序；
建小顶堆，堆顶元素是时长最小的任务，时长相同的按下标排序；
初始化当前时间now = 1，遍历tasks中的任务，当堆空或者now>=tasks[i][0](即任务开始时间)时，将任务加入堆中，之后再弹出堆顶元素，根据弹出的新任务需要执行的时间来更新当前时刻now
重复4直道所有任务都被遍历
依次弹出堆中所有元素

class Solution {
    public int[] getOrder(int[][] oldTask) {
        int n = oldTask.length;
        int[][] tasks = new int[n][3];
        for (int i = 0; i < oldTask.length; i++) {
            tasks[i][0] = oldTask[i][0];
            tasks[i][1] = oldTask[i][1];
            tasks[i][2] = i;
        }

        Arrays.sort(tasks, (x, y) -> x[0] - y[0]);
        PriorityQueue<Integer> heap = new PriorityQueue<>((x, y) -> {
            if (tasks[x][1] == tasks[y][1]) {
                return tasks[x][2] - tasks[y][2];
            }
            return tasks[x][1] - tasks[y][1];
        });
        //时间
        int now = 1;
        //结果
        int[] ret = new int[n];
        //下标
        int i = 0;
        //结果下标
        int j = 0;
        //遍历任务
        while (i < n) {
            //将任务加入到优先队里中
            while (i < n && (heap.isEmpty() || now >= tasks[i][0])) {
                now = Math.max(now, tasks[i][0]);
                heap.add(i++);
            }
            //弹出任务
            int[] task = tasks[heap.poll()];
            ret[j++] = task[2];
            now += task[1];
        }
        //如果任务	队列中还有元素 继续遍历
        while (!heap.isEmpty()) {
            ret[j++] = tasks[heap.poll()][2];
        }
        return ret;
    }

    public static void main(String[] args) {
        new Solution().getOrder(new int[][]{{1, 2}, {2, 4}, {3, 2}, {4, 1}});
    }
}
*/
func main() {
	//getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}})
	getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}})
}
