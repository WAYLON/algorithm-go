package main

import "fmt"

//编写函数进行排序selectSort 完成排序 选择排序
func SelectSort(arr []int) {
	//标准的访问方式
	//(*arr)[0] = 100
	//arr[1] = 100
	for j := 0; j < len(arr)-1; j++ {
		//1.先完成第一个最大值和 arr[0]交换 ==>先易后难
		//假设 arr[0] 最大值
		max := arr[j]
		maxIndex := j
		//2.遍历后面的1--len-1比较
		for i := j + 1; i < len(arr); i++ {
			//找到真正的最大值
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
		fmt.Printf("第%d次循环 %v \n", j+1, arr)
	}
}

func main() {
	//定义一个数组 从大到小排
	arr := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}
	fmt.Println("排序前", arr)
	SelectSort(arr)
	fmt.Println("主函数", arr)

}
