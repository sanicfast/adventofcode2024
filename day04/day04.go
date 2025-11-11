package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readInput() []string {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	rawData, err := os.ReadFile(filename)
	check(err)

	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")

	// make sure we don't have an empty line at the end
	var filteredLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			filteredLines = append(filteredLines, line)
		}
	}
	return filteredLines
}
func checkCoordsDir(crossword []string, row int, col int, dr int, dc int) int {
	// check one direction of a certain coord
	for _, letter := range "XMAS" {
		if row >= len(crossword) || row < 0 {
			return 0
		}
		if col >= len(crossword[row]) || col < 0 {
			return 0
		}

		if crossword[row][col] != byte(letter) {
			return 0
		}
		row += dr
		col += dc
	}
	return 1
}
func coordCheckDirections(crossword []string, row int, col int) int {
	// check one coord
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				sum += checkCoordsDir(crossword, row, col, i, j)
			}
		}
	}
	return sum
}
func countXMASes(crossword []string) int {
	// check the whole crossword
	sum := 0
	for row := range crossword {
		for col := range crossword[row] {
			sum += coordCheckDirections(crossword, row, col)
		}
	}
	return sum
}
func countMASXes(crossword []string) int {
	sum := 0
	for row := range crossword {
		for col := range crossword[row] {
			sum += coordCheckMASX(crossword, row, col)
		}
	}
	return sum
}

func coordCheckMASX(crossword []string, row int, col int) int {
	// check if there's a MASX for a certain coord
	if crossword[row][col] != byte('A') {
		return 0
	}
	if row <= 0 || row >= len(crossword)-1 {
		return 0
	}
	if col <= 0 || col >= len(crossword[row])-1 {
		return 0
	}
	topLeft := crossword[row-1][col-1]
	topRight := crossword[row-1][col+1]
	botLeft := crossword[row+1][col-1]
	botRight := crossword[row+1][col+1]

	MASCount := 0
	if (topLeft == byte('M') && botRight == byte('S')) ||
		(topLeft == byte('S') && botRight == byte('M')) {
		MASCount += 1
	}
	if (botLeft == byte('M') && topRight == byte('S')) ||
		(botLeft == byte('S') && topRight == byte('M')) {
		MASCount += 1
	}
	if MASCount == 2 {
		return 1
	} else {
		return 0
	}
}

func main() {
	startTime := time.Now()

	crossword := readInput()

	part1 := countXMASes(crossword)
	fmt.Println("Day 4, Part 1:", part1)

	part2 := countMASXes(crossword)
	fmt.Println("Day 4, Part 2:", part2)
	fmt.Println(time.Since(startTime))
}
