package main

/**
传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。
传送带上的第 i个包裹的重量为weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。

示例 1：
输入：weights = [1,2,3,4,5,6,7,8,9,10], D = 5
输出：15
解释：
船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
第 1 天：1, 2, 3, 4, 5
第 2 天：6, 7
第 3 天：8
第 4 天：9
第 5 天：10

请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。

示例 2：
输入：weights = [3,2,2,4,1,4], D = 3
输出：6
解释：
船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
第 1 天：3, 2
第 2 天：2, 4
第 3 天：1, 4

示例 3：
输入：weights = [1,2,3,1,1], D = 4
输出：3
解释：
第 1 天：1
第 2 天：2
第 3 天：3
第 4 天：1, 1


提示：
1 <= D <= weights.length <= 50000
1 <= weights[i] <= 500
*/

/**
理论最低运力:只确保所有包裹能够被运送，自然也包括重量最大的包裹，此时理论最低运力为 max ,max为数组 weights 中的最大值
理论最高运力：使得所有包裹在最短时间（一天）内运算完成，此时理论最低运力为 sum，sum 为数组 weights 的总和
*/
func shipWithinDays(weights []int, D int) int {
	low := 0
	high := 0
	for _, v := range weights {
		high += v
		low = max(low, v)
	}
	for low < high {
		mid := (low + high) >> 1
		if check(weights, mid, D) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}

func check(weights []int, mid int, d int) bool {
	n := len(weights)
	cnt := 1
	sum := weights[0]
	for i := 1; i < n; {
		for i < n && sum+weights[i] <= mid {
			sum += weights[i]
			i++
		}
		cnt++
		sum = 0
	}

	return cnt-1 <= d
}

/**
我自己写的
*/
func shipWithinDays2(weights []int, D int) int {
	low := 0
	high := 0
	for _, v := range weights {
		high += v
		low = max(low, v)
	}
	for low < high {
		mid := (low + high) >> 1
		if simulate(weights, mid) <= D {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}

func simulate(weights []int, mid int) int {
	res := 1
	sum := 0
	for i := 0; i < len(weights); i++ {
		sum += weights[i]
		if sum > mid {
			sum = 0
			i--
			res++
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
class Solution {
    public int shipWithinDays(int[] ws, int d) {
        int max = 0, sum = 0;
        for (int w : ws) {
            max = Math.max(max, w);
            sum += w;
        }
        int l = max, r = sum;
        while (l < r) {
            int mid = l + r >> 1;
            if (check(ws, mid, d)) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        return r;
    }

    boolean check(int[] ws, int t, int d) {
        int n = ws.length;
        int cnt = 1;
        for (int i = 1, sum = ws[0]; i < n; sum = 0, cnt++) {
            while (i < n && sum + ws[i] <= t) {
                sum += ws[i];
                i++;
            }
        }
        return cnt - 1 <= d;
    }
}

class Solution {
    public int shipWithinDays(int[] weights, int D) {
        int low = 0;
        int high = 0;
        for (int weight : weights) {
            high += weight;
            low = Math.max(low, weight);
        }
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (simulate(mid, weights) <= D) {
                high = mid;
            } else {
                low = mid + 1;
            }

        }
        return high;
    }

    private int simulate(int mid, int[] weights) {
        int res = 1;
        int sum = 0;
        for (int i = 0; i < weights.length; i++) {
            sum += weights[i];
            if (sum > mid) {
                sum = 0;
                i--;
                res++;
            }
        }
        return res;
    }
}
*/
func main() {
	shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
}
