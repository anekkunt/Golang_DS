package main

import (
	"fmt"
)

// binarySearch function performs a binary search on a sorted array
func binarySearch(arr []int, l, r, x int) int {
	if l >= r {
		return -1 //element not present in array
	}

	mid := l + (r-l)/2

	if arr[mid] == x {
		return mid
	}

	if arr[mid] > x {
		return binarySearch(arr, l, mid-1, x)
	} else {
		return binarySearch(arr, mid+1, r, x)
	}

}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 10

	// Function call
	result := binarySearch(arr, 0, len(arr)-1, x)
	if result != -1 {
		fmt.Printf("Element is present at index %d\n", result)
	} else {
		fmt.Println("Element is not present in array")
	}
}
