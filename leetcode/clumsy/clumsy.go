package main

/**
通常，正整数 n 的阶乘是所有小于或等于 n 的正整数的乘积。例如，factorial(10) = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1。
相反，我们设计了一个笨阶乘 clumsy：在整数的递减序列中，我们以一个固定顺序的操作符序列来依次替换原有的乘法操作符：乘法(*)，除法(/)，加法(+)和减法(-)。
例如，clumsy(10) = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1。然而，这些运算仍然使用通常的算术运算顺序：我们在任何加、减步骤之前执行所有的乘法和除法步骤，并且按从左到右处理乘法和除法步骤。
另外，我们使用的除法是地板除法（floor division），所以10 * 9 / 8等于11。这保证结果是一个整数。
实现上面定义的笨函数：给定一个整数 N，它返回 N 的笨阶乘。

示例 1：
输入：4
输出：7
解释：7 = 4 * 3 / 2 + 1

示例 2：
输入：10
输出：12
解释：12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1

提示：
1 <= N <= 10000
-2^31 <= answer <= 2^31 - 1 （答案保证符合 32 位整数。）
*/

/*
方法1 有乘除直接算 加减直接入栈
*/
func clumsy(N int) int {
	stack := []int{}
	op := 0
	sum := 0
	stack = append(stack, N)
	for i := N - 1; i > 0; i-- {
		if op == 0 {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, temp*i)
		} else if op == 1 {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, temp/i)
		} else if op == 2 {
			stack = append(stack, i)
		} else {
			stack = append(stack, -i)
		}
		op = (op + 1) % 4
	}

	for len(stack) != 0 {
		sum += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return sum
}

/*
方法2 通用模板
*/
func clumsy2(N int) int {
	nums := []int{}
	ops := []byte{}
	// 维护运算符优先级
	m := make(map[byte]int)
	m['*'] = 2
	m['/'] = 2
	m['+'] = 1
	m['-'] = 1

	var calc func()
	calc = func() {
		b := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		a := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		ans := 0
		if op == '+' {
			ans = a + b
		} else if op == '-' {
			ans = a - b
		} else if op == '*' {
			ans = a * b
		} else if op == '/' {
			ans = a / b
		}
		nums = append(nums, ans)
	}

	cs := []byte{'*', '/', '+', '-'}
	j := 0
	for i := N; i > 0; i-- {
		op := cs[j%4]
		nums = append(nums, i)
		// 如果「当前运算符优先级」不高于「栈顶运算符优先级」，说明栈内的可以算
		for len(ops) != 0 && m[ops[len(ops)-1]] >= m[op] {
			calc()
		}
		if i != 1 {
			ops = append(ops, op)
		}
		j++
	}

	// 如果栈内还有元素没有算完，继续算
	for len(ops) != 0 {
		calc()
	}

	return nums[len(nums)-1]
}

/*
方法3 归纳总结
*/
func clumsy3(N int) int {
	special := []int{1, 2, 6, 7}
	diff := []int{1, 2, 2, -1}
	if N <= 4 {
		return special[(N-1)%4]
	}
	return N + diff[N%4]
}

/**
java版
方法1 有乘除直接算 加减直接入栈
class Solution {
    public int clumsy(int N) {
        Deque<Integer> stack = new LinkedList<>();
        int op = 0;
        int sum = 0;
        stack.push(N);
        for (int i = N - 1; i > 0; i--) {
            if (op == 0) {
                stack.push(stack.pop() * i);
            } else if (op == 1) {
                stack.push(stack.pop() / i);
            } else if (op == 2) {
                stack.push(i);
            } else {
                stack.push(-i);
            }
            op = (op + 1) % 4;
        }

        while (!stack.isEmpty()) {
            sum += stack.pop();
        }
        return sum;
    }
}
/**
方法2 通用模板
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.HashMap;
import java.util.Map;

class Solution {
    public int clumsy(int n) {
        Deque<Integer> nums = new ArrayDeque<>();
        Deque<Character> ops = new ArrayDeque<>();
        // 维护运算符优先级
        Map<Character, Integer> map = new HashMap<Character, Integer>() {{
            put('*', 2);
            put('/', 2);
            put('+', 1);
            put('-', 1);
        }};
        char[] cs = new char[]{'*', '/', '+', '-'};
        for (int i = n, j = 0; i > 0; i--, j++) {
            char op = cs[j % 4];
            nums.addLast(i);
            // 如果「当前运算符优先级」不高于「栈顶运算符优先级」，说明栈内的可以算
            while (!ops.isEmpty() && map.get(ops.peekLast()) >= map.get(op)) {
                calc(nums, ops);
            }
            if (i != 1) ops.add(op);
        }
        // 如果栈内还有元素没有算完，继续算
        while (!ops.isEmpty()) calc(nums, ops);
        return nums.peekLast();
    }

    private void calc(Deque<Integer> nums, Deque<Character> ops) {
        int b = nums.pollLast(), a = nums.pollLast();
        int op = ops.pollLast();
        int ans = 0;
        if (op == '+') ans = a + b;
        else if (op == '-') ans = a - b;
        else if (op == '*') ans = a * b;
        else if (op == '/') ans = a / b;
        nums.addLast(ans);
    }
}
方法3 归纳总结
class Solution {
    public int clumsy(int n) {
        int[] special = new int[]{1,2,6,7};
        int[] diff = new int[]{1,2,2,-1};
        if (n <= 4) return special[(n - 1) % 4];
        return n + diff[n % 4];
    }
}
*/

func main() {
	clumsy2(4)
}
