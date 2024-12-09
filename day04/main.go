package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile = "input"
	pattern   = "MAS"
)

func isValidEdge(edges []string) bool {
	for _, value := range edges {
		if value != string(pattern[0]) && value != string(pattern[len(pattern)-1]) {
			return false
		}
	}

	return true
}

func isXmas(data *[][]string, x int, y int) int {
	ne := (*data)[y+1][x+1]
	se := (*data)[y-1][x+1]
	sw := (*data)[y-1][x-1]
	nw := (*data)[y+1][x-1]

	isValidEdge := isValidEdge([]string{ne, se, sw, nw})
	if !isValidEdge {
		return 0
	}

	if (ne == string(pattern[0]) && sw != string(pattern[len(pattern)-1])) || (ne == string(pattern[len(pattern)-1]) && sw != string(pattern[0])) {
		return 0
	}

	if (nw == string(pattern[0]) && se != string(pattern[len(pattern)-1])) || (nw == string(pattern[len(pattern)-1]) && se != string(pattern[0])) {
		return 0
	}

	return 1
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
	for y := 1; y < len(data)-(len(pattern)/2); y++ {
		row := data[y]
		if len(row) <= 0 {
			continue
		}

		for x := 1; x < len(row)-(len(pattern)/2); x++ {
			if data[y][x] != string(pattern[len(pattern)/2]) {
				continue
			}

			appearances += isXmas(&data, x, y)
		}
	}

	fmt.Printf("XMAS appeared %v times\n", appearances)
}
