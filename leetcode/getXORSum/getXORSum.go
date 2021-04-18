package main

/**
列表的 异或和（XOR sum）指对所有元素进行按位 XOR 运算的结果。如果列表中仅有一个元素，那么其 异或和 就等于该元素。
例如，[1,2,3,4] 的 异或和 等于 1 XOR 2 XOR 3 XOR 4 = 4 ，而 [3] 的 异或和 等于 3 。
给你两个下标 从 0 开始 计数的数组 arr1 和 arr2 ，两数组均由非负整数组成。
根据每个(i, j) 数对，构造一个由 arr1[i] AND arr2[j]（按位 AND 运算）结果组成的列表。其中 0 <= i < arr1.length 且 0 <= j < arr2.length 。
返回上述列表的 异或和 。

示例 1：
输入：arr1 = [1,2,3], arr2 = [6,5]
输出：0
解释：列表 = [1 AND 6, 1 AND 5, 2 AND 6, 2 AND 5, 3 AND 6, 3 AND 5] = [0,1,2,0,2,1] ，
异或和 = 0 XOR 1 XOR 2 XOR 0 XOR 2 XOR 1 = 0 。

示例 2：
输入：arr1 = [12], arr2 = [4]
输出：4
解释：列表 = [12 AND 4] = [4] ，异或和 = 4 。

提示：
1 <= arr1.length, arr2.length <= 105
0 <= arr1[i], arr2[j] <= 109
*/
//(b1 ^ b2 ^ ...... ^ bn) & (a1 ^ a2 ^ ...... ^ an)
func getXORSum(arr1 []int, arr2 []int) int {
	res1 := 0
	res2 := 0
	for _, v := range arr1 {
		res1 ^= v
	}
	for _, v := range arr2 {
		res2 ^= v
	}
	return res1 & res2
}

//{ b1 & (a1 ^ a2 ^ ...... ^ an) } ^ { b2 & (a1 ^ a2 ^ ...... ^ an) } ^ ...... ^ { bn & (a1 ^ a2 ^ ...... ^ an) }
func getXORSum2(arr1 []int, arr2 []int) int {
	sum2 := 0
	ret := 0
	for _, v := range arr2 {
		sum2 ^= v
	}
	for _, v := range arr1 {
		ret ^= v & sum2
	}
	return ret
}

/**
解析
1、题目的意思说人话就是：
给两个数组a1、a2......an和b1、b2......bn。让我们计算：
(a1 & b1) ^ (a1 & b2) ^ ...... ^ (a1 & bn) ^
(a2 & b1) ^ (a2 & b2) ^ ...... ^ (a2 & bn) ^
......
(an & b1) ^ (an & b2) ^ ...... ^ (an & bn) = ？
2、理解题目意思后根据结合律、交换律、分配率化简做数学题
原式 = (a1 & b1) ^ (a1 & b2) ^ ...... ^ (a1 & bn) ^
 (a2 & b1) ^ (a2 & b2) ^ ...... ^ (a2 & bn) ^
 ......
 (an & b1) ^ (an & b2) ^ ...... ^ (an & bn)
 = (a1 & b1) ^ (a2 & b1) ^ ...... ^ (an & b1) ^
 (a1 & b2) ^ (a2 & b2) ^ ...... ^ (a2 & b1) ^
 ......
 (a1 & bn) ^ (a2 & bn) ^ ...... ^ (an & bn)
 = { b1 & (a1 ^ a2 ^ ...... ^ an) } ^ { b2 & (a1 ^ a2 ^ ...... ^ an) } ^ ...... ^ { bn & (a1 ^ a2 ^ ...... ^ an) }
 = (b1 ^ b2 ^ ...... ^ bn) & (a1 ^ a2 ^ ...... ^ an)
如果你理解结合律、交换律、分配率，基本上看着就秒懂了。但还不懂可以看看下面一步步解释：
给两个数组a1、a2......an和b1、b2......bn。这个题让我们计算：
(a1 & b1) ^ (a1 & b2) ^ ...... ^ (a1 & bn) ^
(a2 & b1) ^ (a2 & b2) ^ ...... ^ (a2 & bn) ^
......
(an & b1) ^ (an & b2) ^ ...... ^ (an & bn) = ？
其实就是和我们以前做数学时遇到的那种找规律求和的题一样。光想可能不好找出规律，但我们一列出上面的式子我们很容易看出来同一列似乎有某种规律。运用异或运算的交换律与结合律我们按列来计算，来看第一列：
(a1 & b1) ^ (a2 & b1) ^ ...... ^ (an & b1)
我们利用分配率可以转化为：b1 & (a1 ^ a2 ^ ...... ^ an)。
这就是第一列也就是关于b1这一列，后面还有b2到bn，所以再算上后面的b2到bn就是：
{ b1 & (a1 ^ a2 ^ ...... ^ an) } ^ { b2 & (a1 ^ a2 ^ ...... ^ an) } ^ ...... ^ { bn & (a1 ^ a2 ^ ...... ^ an) }
再转化一下就是：(b1 ^ b2 ^ ...... ^ bn) & (a1 ^ a2 ^ ...... ^ an)
*/

/**
class Solution {
    public int getXORSum(int[] arr1, int[] arr2) {
        int res1 = 0, res2 = 0;
        for(int num: arr1){
            res1 ^= num;
        }
        for(int num: arr2){
            res2 ^= num;
        }
        return res1 & res2;
    }
}

class Solution {
    public int getXORSum(int[] arr1, int[] arr2) {
        int sum2 = 0;
        int ret = 0;
        for (int i : arr2) {
            sum2 ^= i;
        }
        for (int i : arr1) {
            ret ^= (i & sum2);
        }
        return ret;
    }
}
*/
func main() {

}
