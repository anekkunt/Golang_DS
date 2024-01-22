package main

import "fmt"

/*
 1. traverse from left and compare adjacent elements and the higher one is placed at right side.
 2. In this way, the largest element is moved to the rightmost end at first.
 3. This process is then continued to find the second largest and place it and so on until the data is sorted.
 time complexity
    Average Case: O(n²)
	Worst  Case:  O(n²)
 space complexity  O(1)
*/

func bubbleSort(a []int) {

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

}

func main() {

	var a = []int{4, 6, 7, 8, 9, 10, 0, -2, 1, 1}

	bubbleSort(a)

	fmt.Println(a)

}
