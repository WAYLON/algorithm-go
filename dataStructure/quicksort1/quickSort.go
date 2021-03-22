package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
快速排序 时间复杂的为nlogn
public class Solution {

    public static void quickSort(int[] arr, int low, int high) {
        if (low >= high) return;

        //基准
        int target = arr[low];
        //变量
        int i = low;
        int j = high;
        //临时变量
        int temp;
        while (i < j) {
            //右侧j寻找小于基准数字的节点
            while (j > i && arr[j] >= target) {
                j--;
            }

            //左侧i寻找大于基准数字的节点
            while (i < j && arr[i] <= target) {
                i++;
            }

            if (i < j) {
                temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }

            //如果i == j找到了 交换基准数字
            arr[low] = arr[i];
            arr[i] = target;

        }

        quickSort(arr, low, i - 1);
        quickSort(arr, i + 1, high);

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
		array[low], array[i] = array[low], array[j]
	}

	quickSort(low, i-1, array)
	quickSort(i+1, high, array)

}
func main() {
	var arr []int
	for i := 0; i < 800000; i++ {
		arr = append(arr, rand.Intn(900000))
	}

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
