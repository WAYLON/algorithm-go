package main

import (
	"fmt"
	"time"
)

/**
快速排序 时间复杂的为nlogn
public class Solution {

   public static void quickSort(int[] arr, int low, int high) {
        int i, j, temp, t;
        if (low > high) {
            return;
        }
        i = low;
        j = high;
        //temp就是基准位
        temp = arr[low];

        while (i < j) {
            //先看右边，依次往左递减
            while (temp <= arr[j] && i < j) {
                j--;
            }
            //再看左边，依次往右递增
            while (temp >= arr[i] && i < j) {
                i++;
            }
            //如果满足条件则交换
            if (i < j) {
                t = arr[j];
                arr[j] = arr[i];
                arr[i] = t;
            }

        }
        //最后将基准为与i和j相等位置的数字交换
        arr[low] = arr[i];
        arr[i] = temp;
        //递归调用左半数组
        quickSort(arr, low, j - 1);
        //递归调用右半数组
        quickSort(arr, j + 1, high);
    }

    public static void main(String[] args) {
        int[] arr = {10, 7, 2, 4, 7, 62, 3, 4, 2, 1, 8, 9, 19};
        quickSort(arr, 0, arr.length - 1);
        for (int i = 0; i < arr.length; i++) {
            System.out.println(arr[i]);
        }
    }
}

*/

func quickSort(low int, high int, array []int) {
	if low > high {
		return
	}
	i := low
	j := high
	target := array[low]

	for i < j {
		for j > i && array[j] >= target {
			j--
		}

		for j > i && array[i] <= target {
			i++
		}
		if j > i {
			array[i], array[j] = array[j], array[i]
		}

	}
	array[low], array[i] = array[i], target
	quickSort(low, i-1, array)
	quickSort(i+1, high, array)

}
func main() {
	var arr []int = []int{1, 23, 4, 5, 6, 7, 56, 46, 456, 45, 654}
	/*for i := 0; i < 8; i++ {
		arr = append(arr, rand.Intn(900000))
	}*/

	//fmt.Println(arr)
	start := time.Now().Unix()
	//调用快速排序
	quickSort(0, len(arr)-1, arr)
	end := time.Now().Unix()
	for _, V := range arr {
		fmt.Printf("%d \t", V)
	}

	fmt.Printf("快速排序法耗时%d秒", end-start)
}
