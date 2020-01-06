package utils

import "sort"

// []int {9,8,7,6,5}
// []int {5,6,7,8,9}

// BubbleSort - Implementation of Bubble sort
func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				keepRunning = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

// Sort - Check the length of array and then decide which algorithm to use for sorting
func Sort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)
}
