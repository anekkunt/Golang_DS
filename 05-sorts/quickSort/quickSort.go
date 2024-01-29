package main

import "fmt"

/*
Divide and Conquer algorithm

The key process in quickSort is a partition(). The target of partitions is to place the pivot (any element can be chosen to be a pivot) at its correct position in the sorted array and put all smaller elements to the left of the pivot, and all greater elements to the right of the pivot.

Partition is done recursively on each side of the pivot after the pivot is placed in its correct position and this finally sorts the array.

	https://www.geeksforgeeks.org/quick-sort/

	best and average : O(nlogn)
	Worst Case: O(n²)
*/

// partition rearranges the elements based on the pivot
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	start := low
	end := high - 1

	for start < end {
		for start < end && arr[start] <= pivot {
			start++
		}
		for start < end && arr[end] > pivot {
			end--
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}
	if arr[start] > pivot {
		arr[start], arr[high] = arr[high], arr[start]
	}

	return start

}

// quickSort sorts the elements using Quick Sort algorithm
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	// pi is partitioning index
	pi := partition(arr, low, high)

	// Recursively sort elements
	quickSort(arr, low, pi-1)
	quickSort(arr, pi+1, high)

}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5, -2}
	fmt.Println("Original array:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:  ", arr)
}
