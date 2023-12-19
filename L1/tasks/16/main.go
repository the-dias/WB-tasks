package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import (
	"fmt"
)

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}

	QuickSort(arr)

	fmt.Println(arr)
}

func QuickSort(arr []int) {
	low := 0
	high := len(arr) - 1
	quickSort(arr, low, high)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
