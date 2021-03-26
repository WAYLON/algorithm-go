package main

/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]

示例 2：
输入：arr = [0,1,2,1], k = 1
输出：[0]

限制：
0 <= k <= arr.length <= 10000
0 <= arr[i]arr<= 10000
*/
/**
方法1 快速选择
*/
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	// 最后一个参数表示我们要找的是下标为k-1的数
	return quickSearch(arr, 0, len(arr)-1, k-1)
}

func quickSearch(arr []int, low int, high int, k int) []int {
	j := partition(arr, low, high)
	// 每快排切分1次，找到排序后下标为j的元素，如果j恰好等于k就返回j以及j左边所有的数；
	// 否则根据下标j与k的大小关系来决定继续切分左段还是右段。
	if j == k {
		return arr[:j+1]
	} else if j > k {
		return quickSearch(arr, low, j-1, k)
	} else {
		return quickSearch(arr, j+1, high, k)
	}
}

// 快排切分，返回下标j，使得比nums[j]小的数都在j的左边，比nums[j]大的数都在j的右边。
func partition(arr []int, low int, high int) int {
	i := low
	j := high
	target := arr[low]

	for i < j {
		for j > i && arr[j] >= target {
			j--
		}

		for j > i && arr[i] <= target {
			i++
		}
		if j > i {
			arr[i], arr[j] = arr[j], arr[i]
		}

	}
	arr[low], arr[i] = arr[i], target

	return i
}

/**
方法2 最大堆
*/
/**
class Solution {
    public int[] getLeastNumbers(int[] arr, int k) {
        if (k == 0) {
            return new int[0];
        }
        // 使用一个最大堆（大顶堆）
        // Java 的 PriorityQueue 默认是小顶堆，添加 comparator 参数使其变成最大堆
        Queue<Integer> heap = new PriorityQueue<>(k, (i1, i2) -> Integer.compare(i2, i1));

        for (int e : arr) {
            // 当前数字小于堆顶元素才会入堆
            if (heap.isEmpty() || heap.size() < k || e < heap.peek()) {
                heap.offer(e);
            }
            if (heap.size() > k) {
                heap.poll(); // 删除堆顶最大元素
            }
        }

        // 将堆中的元素存入数组
        int[] res = new int[heap.size()];
        int j = 0;
        for (int e : heap) {
            res[j++] = e;
        }
        return res;
    }
}
*/

func main() {

}
