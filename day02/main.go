package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "input"

func main() {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Panicf("Can't read file: %v\n", inputFile)
	}

	lines := strings.Split(string(content), "\n")
	safeReportQuantity := 0
	for y, _line := range lines {
		if _line == "" {
			continue
		}

		line := strings.Split(_line, " ")
		isSafe := true
		direction := 1
		for x := 1; x < len(line); x++ {
			num, err := strconv.Atoi(line[x])
			if err != nil {
				log.Panicf("Line %v (%v) contains non-numeric data %v\n", y, line, line[x])
			}

			numPrev, err := strconv.Atoi(line[x-1])
			if err != nil {
				log.Panicf("Line %v (%v) contains non-numeric data %v\n", y, line, line[x-1])
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

		if !isSafe {
			continue
		}

		safeReportQuantity += 1
	}

	fmt.Printf("The number of reports that are safe is: %v\n", safeReportQuantity)
}
