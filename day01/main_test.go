package main

import (
	"testing"
)

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr1 := []int{0 * i, 1 * i, 3 * i, 4 * i, 7 * i, 9 * i}
		arr2 := []int{0 * i, 2 * i, 5 * i, 6 * i, 8 * i}

		merge(arr1, arr2)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{0 * i, 4 * i, 9 * i, 3 * i, 7 * i, 1 * i, 5 * i, 2 * i, 8 * i, 6 * i, 0 * i, 7 * i, 1 * i, 5 * i, 2 * i, 8 * i, 6 * i, 0 * i}
		mergeSort(arr)
	}
}
