package main

import "sort"

/**
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

示例1:
输入: [1,2,3,4,5]
输出: True

示例2:
输入: [0,0,1,2,5]
输出: True

限制：
数组长度为 5
数组的数取值为 [0, 13] .
*/
/**
排序
*/
func isStraight(nums []int) bool {
	joker := 0
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}

/**
set去重
*/
func isStraight1(nums []int) bool {
	set := make(map[int]bool)
	max := 0
	min := 14
	for _, v := range nums {
		if v == 0 {
			continue
		}
		max = mmax(max, v)
		min = mmin(min, v)
		//如果有重复提前返回false
		_, ok := set[v]
		if ok {
			return false
		}
		set[v] = true
	}
	return max-min < 5
}

func mmax(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func mmin(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

/**
java

class Solution {
    public boolean isStraight(int[] nums) {
        Set<Integer> repeat = new HashSet<>();
        int max = 0, min = 14;
        for (int num : nums) {
            if (num == 0) continue; // 跳过大小王
            max = Math.max(max, num); // 最大牌
            min = Math.min(min, num); // 最小牌
            if (repeat.contains(num)) return false; // 若有重复，提前返回 false
            repeat.add(num); // 添加此牌至 Set
        }
        return max - min < 5; // 最大牌 - 最小牌 < 5 则可构成顺子
    }
}

class Solution {
    public boolean isStraight(int[] nums) {
        int joker = 0;
        Arrays.sort(nums); // 数组排序
        for (int i = 0; i < 4; i++) {
            if (nums[i] == 0) joker++; // 统计大小王数量
            else if (nums[i] == nums[i + 1]) return false; // 若有重复，提前返回 false
        }
        return nums[4] - nums[joker] < 5; // 最大牌 - 最小牌 < 5 则可构成顺子
    }
}

自创
class Solution {
    public boolean isStraight(int[] nums) {
        Arrays.sort(nums);
        int num0 = 0;
        int i0 = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != 0) {
                i0 = i;
                break;
            } else {
                num0++;
            }
        }

        for (int i = i0; i < nums.length - 1; i++) {
            int sub = nums[i + 1] - nums[i] - 1;
            if (sub > 0) {
                num0 -= sub;
                if (num0 < 0) {
                    return false;
                }
            } else if (sub < 0) {
                return false;
            }
        }
        return true;
    }
}
*/
func main() {

}
