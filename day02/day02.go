package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() [][]int {
	filename := "day02.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var outData [][]int
	for scanner.Scan() {
		newRow := []int{}

		line := scanner.Text()
		if line == "" {
			continue
		}

		charArray := strings.Fields(line)
		for _, val := range charArray {
			intVal, e := strconv.Atoi(val)
			check(e)
			newRow = append(newRow, intVal)
		}
		outData = append(outData, newRow)
	}
	return outData
}

func absInt(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func calcDiffs(inArray []int) int {
	var min, max int
	for i := range inArray {
		if i == len(inArray)-1 {
			continue
		}
		diff := inArray[i+1] - inArray[i]
		if i == 0 {
			min = diff
			max = diff
		} else {
			if diff < min {
				min = diff
			}
			if diff > max {
				max = diff
			}
		}
	}
	if min == 0 || max == 0 {
		return 0
	} else if min < 0 && max > 0 {
		return 0
	} else if absInt(min) > 3 || absInt(max) > 3 {
		return 0
	} else {
		return 1
	}
}

func part1Diffs(data [][]int) int {
	out := 0
	for i := range data {
		safe := calcDiffs(data[i])
		out += safe
	}
	return out
}

func copySlice(inSlice []int) []int {
	// append input slice onto an empty slice to ensure that a new slice copy is created
	return append([]int{}, inSlice...)
}

func problemDamper(inSlice []int) int {
	// lets see if we can remove any of the levels to make it safe
	for i := range inSlice {
		copy := copySlice(inSlice)
		attempt := append(copy[:i], copy[i+1:]...)
		safe := calcDiffs(attempt)
		if safe == 1 {
			return 1
		}
	}
	return 0
}

func part2Diffs(data [][]int) int {
	// count safe levels with problem damper. first we see if it's safe, if not we try out the damper.
	out := 0
	for i := range data {
		safe := calcDiffs(data[i])
		if safe == 1 {
			out += safe
		} else {
			safeDamp := problemDamper(data[i])
			out += safeDamp
		}
	}
	return out
}

func main() {
	// read in data
	data := readInput()

	// part 1
	part1 := part1Diffs(data)
	fmt.Println("Day 2, Part 1:", part1)

	// part 2
	part2 := part2Diffs(data)
	fmt.Println("Day 2, Part 2:", part2)
}
