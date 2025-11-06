package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readInput() ([]int, []int) {
	//reads in the data
	// Choose filename: first command-line arg overrides default
	filename := "day01.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	rawData, err := os.ReadFile(filename)
	check(err)

	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")

	var intSlice1, intSlice2 []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		numsTmp := strings.Split(line, "   ")
		num1, e1 := strconv.Atoi(numsTmp[0])
		num2, e2 := strconv.Atoi(numsTmp[1])
		check(e1)
		check(e2)
		intSlice1 = append(intSlice1, num1)
		intSlice2 = append(intSlice2, num2)
	}
	return intSlice1, intSlice2
}

func absInt(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func comparify(array1 []int, array2 []int) int {
	var totalAbsDiff int
	for i := range array1 {
		diff := array1[i] - array2[i]
		absDiff := absInt(diff)
		totalAbsDiff += absDiff
	}
	return totalAbsDiff
}

func countUnique(input []int) map[int]int {
	countDict := make(map[int]int)
	for _, num := range input {
		countDict[num] += 1
	}
	return countDict
}

func getTotalSimilarity(myArray []int, myMap map[int]int) int {
	totalSimilarity := 0
	for _, val := range myArray {
		count, _ := myMap[val]
		similarity := val * count
		totalSimilarity += similarity
	}
	return totalSimilarity
}

func main() {
	intSlice1, intSlice2 := readInput()
	sort.Ints(intSlice1)
	sort.Ints(intSlice2)
	part1 := comparify(intSlice1, intSlice2)
	fmt.Println("Day 1, Part 1:", part1)
	uniqueDict := countUnique(intSlice2)
	part2 := getTotalSimilarity(intSlice1, uniqueDict)
	fmt.Println("Day 1, Part 2:", part2)
}
