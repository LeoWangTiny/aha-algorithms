package main

import "fmt"

func main() {
	arr := []int{10, 4, 5, 3, 9, 17}
	//EasyBucketSort(arr)
	//BubbleSort(arr)
	QuickSort(arr)
	fmt.Println(arr)
}

// EasyBucketSort
// 1.1 简易桶排序
// 特点: 当数据密度大，范围小时，比较合适。如果数据稀疏，会造成大量的数据浪费
func EasyBucketSort(arr []int) {
	l := len(arr)
	if len(arr) == 1 {
		return
	}
	max := arr[0]
	for i := 1; i < l; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	bucket := make([]int, max+1)
	for i := 0; i < l; i++ {
		bucket[arr[i]]++
	}
	j := 0
	for k, v := range bucket {
		for i := 0; i < v; i++ {
			arr[j] = k
			j++
		}
	}
	return
}

// BubbleSort
// 1.2冒泡排序
// 特点: 主要是阐述了算法思想，无论从算法效率还是执行次数上都可以说是最糙的了
func BubbleSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	for i := 0; i < l-1; i++ {
		for j := 1; j < l-1; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return
}

func QuickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) {
	if left > right {
		return
	}
	i, j, tmp := left, right, arr[left]
	for i != j {
		for arr[j] >= tmp && i < j {
			j--
		}
		for arr[i] <= tmp && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left], arr[i] = arr[i], tmp
	_quickSort(arr, left, i-1)
	_quickSort(arr, i+1, right)
}
