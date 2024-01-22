package main

import "fmt"

/*time: O(nlogn)
space:O(n)*/

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	// Create temp arrays
	L := make([]int, n1)
	R := make([]int, n2)

	// Copy data to temp arrays L[] and R[]
	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[m+1+j]
	}

	// Merge the temp arrays back into arr[l..r]
	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of L[], if there are any
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	// Copy the remaining elements of R[], if there are any
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

// mergeSort sorts arr[l..r] using
// merge() to merge the subarrays
func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2 //safe  l + (r-l)/2

	// Sort first and second halves
	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)

	// Merge the sorted halves
	merge(arr, l, m, r)

}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7, -2, 0}
	fmt.Println("Given array is: ", arr)

	mergeSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array is:", arr)
}
