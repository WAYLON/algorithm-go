package main

/**
给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。

示例:
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
/**
解题思路：只有凹的地方能存水，存水量遵循短板原则，所以用每个位置左右两侧最大值中的较小者减当前位置的值即可得到当前位置储水量。
解题方法：先倒叙遍历，用数组记录每个位置其右侧最大值max右，再正序遍历，时刻记录并更新当前位置左侧的最大值max左，然后当前位置存水量c=Min(max左，max右)-当前值，
如果c<=0则表示没有水，抛弃即可，最后每个位置的c累加一起的和即为总储水量。
*/
/**
方法1 双指针
*/
func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	res := 0
	for i := 1; i < n-1; i++ {
		t := min(leftMax[i], rightMax[i]) - height[i]
		res += max(t, 0)
	}
	return res
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

/**
方法2 单调栈
*/
func trap2(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	res := 0
	//存下标 用的时候取对应的下标高度即可 单调栈  栈顶最小 栈底最大
	var stack []int

	for i := 0; i < n; i++ {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//如果栈内没有元素，说明当前位置左边没有比其高的柱子，跳过
			if len(stack) != 0 {
				//比较当前值和栈顶元素哪个最小 - 中间元素高度
				h := min(height[stack[len(stack)-1]], height[i]) - height[mid]
				// 注意减一，只求中间宽度
				w := i - stack[len(stack)-1] - 1
				res += h * w
			}
		}
		stack = append(stack, i)
	}
	return res
}

/**
方法3 面积差值
*/
func trap3(height []int) int {
	n := len(height)

	sum := 0
	m := 0
	for i := 0; i < n; i++ {
		cur := height[i]
		sum += cur
		m = max(m, cur)
	}
	full := m * n

	lSum := 0
	lMax := 0
	for i := 0; i < n; i++ {
		lMax = max(lMax, height[i])
		lSum += lMax
	}

	rSum := 0
	rMax := 0
	for i := n - 1; i >= 0; i-- {
		rMax = max(rMax, height[i])
		rSum += rMax
	}

	return lSum + rSum - full - sum
}

/**
java版
方法1 双指针
class Solution {
    public int trap(int[] height) {
        int len = height.length;
        if (len == 0) return 0;

        int[] leftMax = new int[len];
        leftMax[0] = height[0];
        for (int i = 1; i < len; i++) {
            leftMax[i] = Math.max(leftMax[i - 1], height[i]);
        }

        int[] rightMax = new int[len];
        rightMax[len - 1] = height[len - 1];
        for (int j = len - 2; j >= 0; j--) {
            rightMax[j] = Math.max(rightMax[j + 1], height[j]);
        }

        int res = 0;
        for (int i = 1; i < len - 1; i++) {
            int t = Math.min(leftMax[i], rightMax[i]) - height[i];
            res += Math.max(t, 0);
        }
        return res;

    }

    public static void main(String[] args) {
        new Solution().trap(new int[]{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1});
    }
}
方法2 单调栈
class Solution {
    public int trap(int[] height) {
        int n = height.length;
        if (n <= 2) return 0;
        int ans = 0;
        //存下标 用的时候取对应的下标高度即可 单调栈  栈顶最小 栈底最大
        Deque<Integer> stack = new ArrayDeque<>();

        for (int i = 0; i < n; i++) {
            while (!stack.isEmpty() && height[i] > height[stack.peek()]) {
                int mid = stack.pop();
                //如果栈内没有元素，说明当前位置左边没有比其高的柱子，跳过
                if (!stack.isEmpty()) {
                    //比较当前值和栈顶元素哪个最小 - 中间元素高度
                    int h = Math.min(height[stack.peek()], height[i]) - height[mid];
                    // 注意减一，只求中间宽度
                    int w = i - stack.peek() - 1;
                    ans += h * w;
                }
            }
            stack.push(i);
        }
        return ans;
    }
}
方法3 面积差值
class Solution {
    public int trap(int[] height) {
        int n = height.length;

        int sum = 0, max = 0;
        for (int i = 0; i < n; i++) {
            int cur = height[i];
            sum += cur;
            max = Math.max(max, cur);
        }
        int full = max * n;

        int lSum = 0, lMax = 0;
        for (int i = 0; i < n; i++) {
            lMax = Math.max(lMax, height[i]);
            lSum += lMax;
        }

        int rSum = 0, rMax = 0;
        for (int i = n - 1; i >= 0; i--) {
            rMax = Math.max(rMax, height[i]);
            rSum += rMax;
        }

        return lSum + rSum - full - sum;
    }
}

*/
func main() {

}
