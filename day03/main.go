package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "input"

var (
	patternStart = []byte("mul(")
	patternEnd   = []byte(")")
)

func main() {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Panicf("Can't read file: %v\n", inputFile)
	}

	num1 := ""
	num2 := ""

	match := 0
	sum := 0
	for index := 0; index < len(content); index++ {
		// After num(x,
		if match >= 2 {
			value := string(content[index])
			if value == ")" {
				match = 0
				num1Value, err := strconv.Atoi(num1)
				if err != nil {
					log.Panicf("Incorrect logic in building num1: %v\n", num1)
				}

				num2Value, err := strconv.Atoi(num2)
				if err != nil {
					log.Panicf("Incorrect logic in building num2: %v\n", num2)
				}

				sum += num1Value * num2Value
				num1 = ""
				num2 = ""
				continue
			}

			_, err := strconv.Atoi(value)
			if err != nil {
				num1 = ""
				num2 = ""
				match = 0
				continue
			}

			num2 += value
			continue
		}

		// After num(
		if match >= 1 {
			value := string(content[index])
			if value == "," {
				if num1 == "" {
					num1 = ""
					num2 = ""
					match = 0
					continue
				}

				match = 2
				continue
			}

			_, err := strconv.Atoi(value)
			if err != nil {
				num1 = ""
				num2 = ""
				match = 0
				continue
			}

			num1 += value
			continue
		}


		value := string(content[index:index+4])
		if value == string(patternStart) {
			index += 3
			match = 1
			continue
		}

		if index >= len(content) - 4 {
			break
		}
	}

	fmt.Printf("The sum of products is: %v\n", sum)
}