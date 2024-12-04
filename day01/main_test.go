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

func TestMerge(t *testing.T) {
    arr1 := []int{0, 1, 3, 4, 7, 9}
    arr2 := []int{0, 2, 5, 6, 8}

    resultArr := merge(arr1, arr2)
    expectedArr := []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    if len(resultArr) != len(expectedArr) {
        t.Fatal("Lengths of the resulting array and expected array are different")
    }

    for i, value := range resultArr {
        if value != expectedArr[i] {
            t.Fatal("Resulting array and expected array do not match")
        }
    }

}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{0 * i, 4 * i, 9 * i, 3 * i, 7 * i, 1 * i, 5 * i, 2 * i, 8 * i, 6 * i, 0 * i, 7 * i, 1 * i, 5 * i, 2 * i, 8 * i, 6 * i, 0 * i}
		mergeSort(arr)
	}
}

func TestMargeSort(t *testing.T) {
    inputArr := []int{0, 4, 9, 3, 7, 1, 5, 2, 8, 6, 0, 7, 1, 5, 2, 8, 6, 0}
    resultArr := mergeSort(inputArr)
    expectedArr := []int{0, 0, 0, 1, 1, 2, 2, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9}

    for i, value := range resultArr {
        if value != expectedArr[i] {
            t.Fatal("Resulting array and expected array do not match")
        }
    }
}
