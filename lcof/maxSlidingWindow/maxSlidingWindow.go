package main

/**
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7


提示：
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤输入数组的大小。
注意：本题与主站 239 题相同：https://leetcode-cn.com/problems/sliding-window-maximum/
*/
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}
	rn := n - k + 1
	res := make([]int, rn)
	queue := []int{}
	for i, j := 0, 0; i < n; i++ {
		//如果队列中不为空并且当前下标-栈顶元素下标 >= k
		if len(queue) != 0 && i-queue[0] >= k {
			queue = queue[1:]
		}
		//如果队列不为空并且 要加入的下标值大于队尾值 则将队列最后一个值删掉 维护一个递减队列
		for len(queue) != 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		//将当前下标加入到队列中
		queue = append(queue, i)

		//当前窗口大于k-1
		if i >= k-1 {
			res[j] = nums[queue[0]]
			j++
		}
	}

	return res
}

/**
class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        int n = nums.length;
        if (k == 0 || n == 0) return new int[]{};
        int rn = n - k + 1;
        int[] res = new int[rn];
        Deque<Integer> queue = new ArrayDeque<>();
        for (int i = 0, j = 0; i < nums.length; i++) {
            //如果队列中不为空并且当前下标-栈顶元素下标 >= k
            if (!queue.isEmpty() && i - queue.peek() >= k) {
                queue.poll();
            }
            //如果队列不为空并且 要加入的下标值大于队尾值 则将队列最后一个值删掉 维护一个递减队列
            while (!queue.isEmpty() && nums[i] > nums[queue.peekLast()]) {
                queue.pollLast();
            }
            //将当前下标加入到队列中
            queue.offer(i);

            //当前窗口大于k-1
            if (i >= k - 1) {
                res[j++] = nums[queue.peek()];
            }
        }
        return res;
    }
}
*/
func main() {

}
