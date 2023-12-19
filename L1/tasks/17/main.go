package main

// Реализовать бинарный поиск встроенными методами языка.

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 4, 7, 10, 15, 20, 25, 30}

	fmt.Println(binarySearchWithIncludedMethods(arr, 30))
	fmt.Println(binarySearch(arr, 15))

}

func binarySearchWithIncludedMethods(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
