package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []string {
	filename := "day04_ex.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	rawData, err := os.ReadFile(filename)
	check(err)

	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")

	return lines
}

func checkCoordsDir(crossword []string, row int, col int, dr int, dc int) int {

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
func countXMAXes(crossword []string) int {
	sum := 0
	for row := range crossword {
		for col := range crossword[row] {
			sum += coordCheckDirections(crossword, row, col)
		}
	}
	return sum
}

func main() {
	crossword := readInput()
	// for _, line := range crossword {
	// 	fmt.Println(line)
	// }
	part1 := countXMAXes(crossword)
	fmt.Println("Part 1", part1)
}
