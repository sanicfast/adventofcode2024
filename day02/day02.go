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
	filename := "day02_ex.txt"
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
func main() {
	data := readInput()
	fmt.Println(len(data))
	fmt.Println(data)
}
