package main

/**
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

示例 1:
输入: [7,5,6,4]
输出: 5

限制：
0 <= 数组长度 <= 50000
*/
var n int

func reversePairs(nums []int) int {
	n = 0
	mergeSort(nums, 0, len(nums)-1)
	return n
}

func mergeSort(a []int, low int, high int) {
	mid := (low + high) / 2
	if low < high {
		//左面
		mergeSort(a, low, mid)
		mergeSort(a, mid+1, high)
		//左右归并
		merge(a, low, mid, high)
	}
}

func merge(a []int, low int, mid int, high int) {
	temp := make([]int, high-low+1)
	i := low     //左指针
	j := mid + 1 //右指针
	k := 0
	// 把较小的数先移到新数组中
	for i <= mid && j <= high {
		if a[i] <= a[j] {
			temp[k] = a[i]
			k++
			i++
		} else {
			n += mid - i + 1
			temp[k] = a[j]
			k++
			j++
		}
	}
	// 把左边剩余的数移入数组
	for i <= mid {
		temp[k] = a[i]
		k++
		i++
	}
	// 把右边边剩余的数移入数组
	for j <= high {
		temp[k] = a[j]
		k++
		j++
	}

	// 把新数组中的数覆盖nums数组
	for i := 0; i < len(temp); i++ {
		a[low+i] = temp[i]
	}
}

func main() {

}

/**
class Solution {
    int n;

    public int reversePairs(int[] nums) {
        n = 0;
        mergeSort(nums, 0, nums.length - 1);
        return n;
    }

    public void mergeSort(int[] a, int low, int high) {
        int mid = low + (high - low) / 2;
        if (low < high) {
            mergeSort(a, low, mid);
            mergeSort(a, mid + 1, high);
            merge(a, low, mid, high);
        }
    }

    private void merge(int[] a, int low, int mid, int high) {
        int[] temp = new int[high - low + 1];
        int i = low;
        int j = mid + 1;
        int k = 0;
        while (i <= mid && j <= high) {
			//注意此处应该是<= 因为 等于时不算逆序对
            if (a[i] <= a[j]) {
                temp[k++] = a[i++];
            } else {
                n += (mid - i+1);
                temp[k++] = a[j++];
            }
        }

        while (i <= mid) {
            temp[k++] = a[i];
            i++;
        }

        while (j <= high) {
            temp[k++] = a[j];
            j++;
        }

        for (int v = 0; v < temp.length; v++) {
            a[low + v] = temp[v];
        }
    }
}
*/
