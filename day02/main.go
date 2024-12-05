package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input"

func safeCheck(lineSlice []string) bool {
	direction := 1
	isSafe := true
	for x := 1; x < len(lineSlice); x++ {
		num, err := strconv.Atoi(lineSlice[x])
		if err != nil {
			log.Panicf("%v contains non-numeric data %v\n", lineSlice, lineSlice[x])
		}

		numPrev, err := strconv.Atoi(lineSlice[x-1])
		if err != nil {
			log.Panicf("%v contains non-numeric data %v\n", lineSlice, lineSlice[x-1])
		}

		difference := numPrev - num
		if x == 1 && difference < 0 {
			direction = -1
		}

		if difference*direction < 1 || difference*direction > 3 {
			isSafe = false
			break
		}
	}

	return isSafe
}

func main() {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Panicf("Can't read file: %v\n", inputFile)
	}

	lines := strings.Split(string(content), "\n")
	safeReportQuantity := 0
	for _, _line := range lines {
		if _line == "" {
			continue
		}

		line := strings.Split(_line, " ")
		isSafe := false
		for x := 0; x < len(line); x++ {
			lineSlice := make([]string, 0, len(line)-1)
			lineSlice = append(lineSlice, line[:x]...)
			lineSlice = append(lineSlice, line[x+1:]...)

			isSafe = safeCheck(lineSlice)
			if isSafe {
				safeReportQuantity += 1
				break
			}
		}
	}

	fmt.Printf("The number of reports that are safe is: %v\n", safeReportQuantity)
}
