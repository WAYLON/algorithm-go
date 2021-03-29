package main

/**
归并排序 时间复杂度O(nlogn)  空间复杂度 O(n)
*/
func mergeSort(a []int, low int, high int) {
	mid := (low + high) / 2
	if low < high {
		//左面
		mergeSort(a, low, mid)
		mergeSort(a, low, high)
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
		if a[i] < a[j] {
			temp[k] = a[i]
			k++
			i++
		} else {
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
public static void mergeSort(int[] a, int low, int high) {
       int mid = (low + high) / 2;
       if (low < high) {
           // 左边
           mergeSort(a, low, mid);
           // 右边
           mergeSort(a, mid + 1, high);
           // 左右归并
           merge(a, low, mid, high);
           System.out.println(Arrays.toString(a));
       }

}

public static void merge(int[] a, int low, int mid, int high) {
       int[] temp = new int[high - low + 1];
       int i = low;// 左指针
       int j = mid + 1;// 右指针
       int k = 0;
       // 把较小的数先移到新数组中
       while (i <= mid && j <= high) {
           if (a[i] < a[j]) {
               temp[k++] = a[i++];
           } else {
               temp[k++] = a[j++];
           }
       }
       // 把左边剩余的数移入数组
       while (i <= mid) {
           temp[k++] = a[i++];
       }
       // 把右边边剩余的数移入数组
       while (j <= high) {
           temp[k++] = a[j++];
       }
       // 把新数组中的数覆盖nums数组
       for (int k2 = 0; k2 < temp.length; k2++) {
           a[k2 + low] = temp[k2];
       }
   }
*/
