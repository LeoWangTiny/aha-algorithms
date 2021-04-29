package main

import "fmt"

func main() {
	arr := []int{10, 4, 5, 3, 9, 17}
	EasyBucketSort(arr)
	fmt.Println("bucket sort result:", arr)

	arr = []int{10, 4, 5, 3, 9, 17}
	BubbleSort(arr)
	fmt.Println("bubble sort result:", arr)

	arr = []int{10, 4, 5, 3, 9, 17}
	QuickSort(arr)
	fmt.Println("quick sort result:", arr)
}

// EasyBucketSort
// 1.1 简易桶排序
// 特点: 当数据密度大，范围小时，比较合适。如果数据稀疏，会造成大量的数据浪费
// 时间复杂度O(N+M) n代表排序的元素个数，m代表桶的个数
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
// 时间复杂度O(N^2)
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

// QuickSort
// 快速排序
// 特点: 通过分治的方式，降低了时间复杂度。但是在最坏的情况，当数组的顺序为倒序时，算法会进行退化
// 平均时间复杂度为O(N*logN) 最坏时间复杂度为O(N^2)
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
