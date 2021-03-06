package main

import "math"

/**
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

示例:
输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

说明:
1是丑数。
n不超过1690。
*/

/**
我们知道，丑数的排列肯定是1,2,3,4,5,6,8,10.... 然后有一个特点是，任意一个丑数都是由小于它的某一个丑数*2，*3或者*5得到的，那么如何得到所有丑数呢？ 现在假设有3个数组，分别是： A：{1*2，2*2，3*2，4*2，5*2，6*2，8*2，10*2......}
B：{1*3，2*3，3*3，4*3，5*3，6*3，8*3，10*3......}
C：{1*5，2*5，3*5，4*5，5*5，6*5，8*5，10*5......}
那么所有丑数的排列，必定就是上面ABC3个数组的合并结果然后去重得到的，那么这不就转换成了三个有序数组的无重复元素合并的问题了吗？而这三个数组就刚好是{1,2,3,4,5,6,8,10....}乘以2,3,5得到的。
合并有序数组的一个比较好的方法，就是每个数组都对应一个指针，然后比较这些指针所指的数中哪个最小，就将这个数放到结果数组中，然后该指针向后挪一位。
回到本题，要求丑数ugly数组中的第n项，而目前只知道ugly[0]=1，所以此时三个有序链表分别就只有一个元素：
A ： {1*2......}
B ： {1*3......}
C ：{1*5......}
假设三个数组的指针分别是i,j,k，此时均是指向第一个元素，然后比较A[i]，B[j]和C[k]，得到的最小的数A[i]，就是ugly[1]，此时ugly就变成{1,2}了，对应的ABC数组就分别变成了：
A ： {1*2，2*2......}
B ： {1*3, 2*3......}
C ：{1*5,2*5......}
此时根据合并有序数组的原理，A数组指针i就指向了下一个元素，即'2*2'，而j和k依然分别指向B[0]和C[0]，然后进行下一轮合并，就是A[1]和B[0]和C[0]比较，最小值作为ugly[2].....如此循环n次，就可以得到ugly[n]了。
此外，注意到ABC三个数组实际上就是ugly[]*2，ugly[]*3和ugly[]*5的结果，所以每次只需要比较A[i]=ugly[i]*2，B[j]=ugly[j]*3和C[k]=ugly[k]*5的大小即可。然后谁最小，就把对应的指针往后移动一个，为了去重，如果多个元素都是最小，那么这多个指针都要往后移动一个。
三指针
*/
func nthUglyNumber(n int) int {
	res := make([]int, n)
	res[0] = 1
	i2 := 0
	i3 := 0
	i5 := 0
	for i := 1; i < n; i++ {
		n2 := res[i2] * 2
		n3 := res[i3] * 3
		n5 := res[i5] * 5
		ugly := int(math.Min(float64(n2), math.Min(float64(n3), float64(n5))))
		if ugly == n2 {
			i2++
		}
		if ugly == n3 {
			i3++
		}
		if ugly == n5 {
			i5++
		}
		res[i] = ugly
	}
	return res[n-1]
}

/**
java版
class Solution {
    public int nthUglyNumber(int n) {
        int[] nums = new int[n];
        nums[0] = 1;
        int i2 = 0, i3 = 0, i5 = 0;
        for (int i = 1; i < n; i++) {
            int a2 = nums[i2] * 2;
            int a3 = nums[i3] * 3;
            int a5 = nums[i5] * 5;
            int ugly = Math.min(a2, Math.min(a3, a5));
            if (ugly == a2) {
                i2++;
            }
            if (ugly == a3) {
                i3++;
            }
            if (ugly == a5) {
                i5++;
            }
            nums[i] = ugly;
        }

        return nums[n - 1];
    }
}
*/
func main() {
	nthUglyNumber(11)
}
