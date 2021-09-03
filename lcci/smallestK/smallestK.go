package main

import "math/rand"

/**
设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。

示例：
输入： arr = [1,3,5,7,2,4,6,8], k = 4
输出： [1,2,3,4]

提示：
0 <= len(arr) <= 100000
0 <= k <= min(100000, len(arr))
*/
func smallestK(arr []int, k int) []int {
	n := len(arr)
	ans := make([]int, k)
	if k == 0 {
		return ans
	}

	var qsort func(arr []int, l int, r int)
	qsort = func(arr []int, l int, r int) {
		if l >= r {
			return
		}
		i := l
		j := r
		ridx := rand.Intn(r-l+1) + l
		arr[ridx], arr[l] = arr[l], arr[ridx]
		x := arr[l]
		for i < j {
			for i < j && arr[j] >= x {
				j--
			}

			for i < j && arr[i] <= x {
				i++
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[i], arr[l] = arr[l], arr[i]

		if i > k {
			qsort(arr, l, i-1)
		}

		if i < k {
			qsort(arr, i+1, r)
		}

	}

	qsort(arr, 0, n-1)
	copy(ans, arr[:k])
	return ans
}

/** java
class Solution {
    int k;

    public int[] smallestK(int[] arr, int _k) {
        k = _k;
        int n = arr.length;
        int[] ans = new int[k];
        if (k == 0) {
            return ans;
        }
        qsort(arr, 0, n - 1);
        if (k >= 0) {
            System.arraycopy(arr, 0, ans, 0, k);
        }
        return ans;
    }

    void qsort(int[] arr, int l, int r) {
        if (l >= r) {
            return;
        }
        int i = l, j = r;
        int ridx = new Random().nextInt(r - l + 1) + l;
        swap(arr, ridx, l);
        int x = arr[l];
        while (i < j) {
            while (i < j && arr[j] >= x) {
                j--;
            }
            while (i < j && arr[i] <= x) {
                i++;
            }
            swap(arr, i, j);
        }
        swap(arr, i, l);
        // 集中答疑：因为题解是使用「基准点左侧」来进行描述（不包含基准点的意思），所以这里用的 k（写成 k - 1 也可以滴
        if (i > k) {
            qsort(arr, l, i - 1);
        }
        if (i < k) {
            qsort(arr, i + 1, r);
        }
    }

    void swap(int[] arr, int l, int r) {
        int tmp = arr[l];
        arr[l] = arr[r];
        arr[r] = tmp;
    }
}

class Solution {
    public int[] smallestK(int[] arr, int k) {
        int[] res = new int[k];
        Arrays.sort(arr);
        for (int i = 0; i < k; i++) {
            res[i] = arr[i];
        }
        return res;
    }
}

class Solution {
    public int[] smallestK(int[] arr, int k) {
        int[] res = new int[k];
        //小根堆
        PriorityQueue<Integer> queue = new PriorityQueue<Integer>((x, y) -> x - y);
        for (int v : arr) {
            queue.add(v);
        }
        for (int i = 0; i < k; i++) {
            res[i] = queue.poll();
        }
        return res;
    }
}

class Solution {
    public int[] smallestK(int[] arr, int k) {
        int[] res = new int[k];
        if (k == 0) return res;
        //大根堆
        PriorityQueue<Integer> queue = new PriorityQueue<Integer>((x, y) -> y - x);
        for (int v : arr) {
            if (queue.size() == k && queue.peek() < v) {
                continue;
            }

            if (queue.size() == k) {
                queue.poll();
            }

            queue.add(v);
        }
        for (int i = k - 1; i >= 0; i--) {
            res[i] = queue.poll();
        }
        return res;
    }
}

*/
func main() {

}
