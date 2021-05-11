package main

/**
给你一个整数数组perm，它是前n个正整数的排列，且n是个 奇数。
它被加密成另一个长度为 n - 1的整数数组encoded，满足encoded[i] = perm[i] XOR perm[i + 1]。比方说，如果perm = [1,3,2]，那么encoded = [2,1]。
给你encoded数组，请你返回原始数组perm。题目保证答案存在且唯一。

示例 1：
输入：encoded = [3,1]
输出：[1,2,3]
解释：如果 perm = [1,2,3] ，那么 encoded = [1 XOR 2,2 XOR 3] = [3,1]

示例 2：
输入：encoded = [6,5,4,6]
输出：[2,4,1,5,3]

提示：
3 <= n <105
n是奇数。
encoded.length == n - 1
*/

/**
因为是小于等于n的数，且不会重复，即所有数字在[1,n]中且不重复
* 设n=5
  则perm=[a,b,c,d,e]
      enco=[f,g,h,i]
  因为n确定，所以可以知道所有数的异或结果
  即a^b^c^d^e的结果是知道的
* 我们知道的是enco的值，只需要找到perm的随意一个位置的值就可以构造答案
* 可以发现
  f=a^b
  g=b^c
  h=c^d
  i=d^e
* 现在知道a^b^c^d^e，可以用g=b^c和i=d^e去消除，只剩下一个a，这样perm就构造出来一个元素，即答案
*/
func decode(encoded []int) []int {
	n := len(encoded) + 1
	res := make([]int, n)
	a := 0
	for i := 1; i <= n; i++ {
		a ^= i
	}
	b := 0
	for i := 1; i < n-1; i++ {
		b ^= encoded[i]
	}
	res[0] = a ^ b
	for i := 1; i < n; i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}
	return res
}

/**
class Solution {
    public int[] decode(int[] encoded) {
        int n = encoded.length + 1;
        int[] res = new int[n];
        int a = 0;
        for (int i = 1; i <= n; i++) {
            a ^= i;
        }
        int b = 0;
        for (int i = 1; i < n - 1; i += 2) {
            b ^= encoded[i];
        }
        res[0] = a ^ b;
        for (int i = 1; i < n; i++) {
            res[i] = res[i - 1] ^ encoded[i - 1];
        }
        return res;
    }
}
*/
func main() {

}
