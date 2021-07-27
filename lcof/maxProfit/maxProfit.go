package main

/**
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

示例 1:
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

示例 2:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

限制：
0 <= 数组长度 <= 10^5
*/
/**
动态规划 前i天的最大收益 = max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
*/
func maxProfit(prices []int) int {
	n := len(prices)
	res := 0
	if n < 2 {
		return res
	}
	min := prices[0]
	for i := 1; i < n; i++ {
		if prices[i] <= min {
			min = prices[i]
		} else {
			if res < prices[i]-min {
				res = prices[i] - min
			}
		}
	}
	return res
}

/**
java
class Solution {
    //动态规划 前i天的最大收益 = max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
    public int maxProfit(int[] prices) {
        int length = prices.length;
        if (length < 2) {
            return 0;
        }
        int[] dp = new int[prices.length];
        int min = prices[0];
        for (int i = 1; i < prices.length; i++) {
            dp[i] = Math.max(dp[i - 1], prices[i] - min);
            min = Math.min(min, prices[i]);
        }
        return dp[prices.length - 1];
    }
}

class Solution {
    public int maxProfit(int[] prices) {
        if(prices == null || prices.length <= 1) {
            return 0;
        }
        int res = 0, min = prices[0];
        for(int i = 1; i < prices.length; i++) {
            if(prices[i] <= min) {
                min = prices[i];
            }else {
                res = Math.max(res, prices[i] - min);
            }
        }
        return res;
    }
}
*/
func main() {

}
