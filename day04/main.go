package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "input"
	pattern   = "XMAS"
)

func isXmas(data *[][]string, x int, y int, xDirection int, yDirection int) int {
	add := 1
	for i := 1; i < len(pattern); i++ {
		if (*data)[y+(i*yDirection)][x+(i*xDirection)] != string(pattern[i]) {
			add = 0
			break
		}
	}
	return add
}

func main() {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Panicf("Can't read file: %v\n", inputFile)
	}

	contentRows := strings.Split(string(content), "\n")
	data := make([][]string, len(contentRows)-1)
	for index := 0; index < len(contentRows)-1; index++ {
		data[index] = strings.Split(contentRows[index], "")
	}

	appearances := 0
	for y := 0; y < len(data); y++ {
		row := data[y]
		if len(row) <= 0 {
			continue
		}

		for x := 0; x < len(row); x++ {
			if data[y][x] != string(pattern[0]) {
				continue
			}

			n := y >= len(pattern) - 1
			e := x <= len(row)-len(pattern)
			s := y <= len(data)-len(pattern)
			w := x >= len(pattern) - 1

			if n {
				appearances += isXmas(&data, x, y, 0, -1)
			}

			if n && e {
				appearances += isXmas(&data, x, y, 1, -1)
			}

			if e {
				appearances += isXmas(&data, x, y, 1, 0)
			}

			if s && e {
				appearances += isXmas(&data, x, y, 1, 1)
			}

			if s {
				appearances += isXmas(&data, x, y, 0, 1)
			}

			if s && w {
				appearances += isXmas(&data, x, y, -1, 1)
			}

			if w {
				appearances += isXmas(&data, x, y, -1, 0)
			}

			if n && w {
				appearances += isXmas(&data, x, y, -1, -1)
			}
		}
	}

	fmt.Printf("XMAS appeared %v times\n", appearances)
}
