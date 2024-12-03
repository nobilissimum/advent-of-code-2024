package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func merge(arr1 []int, arr2 []int) []int {
	pointer1 := 0
	pointer2 := 0

	totalLength := len(arr1) + len(arr2)
	result := make([]int, totalLength)

	for i := 0; i < totalLength; i++ {
		if pointer1 >= len(arr1) {
			result[i] = arr2[pointer2]
			pointer2 += 1
		} else if pointer2 >= len(arr2) {
			result[i] = arr1[pointer1]
			pointer1 += 1
		} else if arr1[pointer1] < arr2[pointer2] {
			result[i] = arr1[pointer1]
			pointer1 += 1
		} else {
			result[i] = arr2[pointer2]
			pointer2 += 1
		}
	}

	return result
}

func mergeSort(arr []int) []int {
	midpoint := len(arr) / 2

	arr1 := arr[:midpoint]
	if len(arr1) > 1 {
		arr1 = mergeSort(arr1)
	}

	arr2 := arr[midpoint:]
	if len(arr2) > 1 {
		arr2 = mergeSort(arr2)
	}

	return merge(arr1, arr2)
}

func main() {
	var locationIDs1 []int
	var locationIDs2 []int

	ids := make(map[int]int)
	reader := bufio.NewReader(os.Stdin)
	for true {
		input, _ := reader.ReadString('\n')
		if input == "\n" {
			break
		}
		inputArr := strings.Fields(input)

		col1, err := strconv.Atoi(inputArr[0])
		if err != nil {
			log.Panic("Invalid input %v", inputArr[0])
		}

		col2, err := strconv.Atoi(inputArr[1])
		if err != nil {
			log.Panic("Invalid input %v", inputArr[0])
		}

		locationIDs1 = append(locationIDs1, col1)
		locationIDs2 = append(locationIDs2, col2)

		ids[col2] += 1
	}

	locationIDs1 = mergeSort(locationIDs1)
	locationIDs2 = mergeSort(locationIDs2)

	totalDistance := 0
	totalSimilarity := 0
	for i := 0; i < len(locationIDs1); i++ {
		value := locationIDs1[i]
		totalDistance += int(math.Abs(float64(value) - float64(locationIDs2[i])))
		totalSimilarity += value * ids[value]
	}

	fmt.Printf("Total distance: %v\n", totalDistance)
	fmt.Printf("Total similarity: %v\n", totalSimilarity)
}
