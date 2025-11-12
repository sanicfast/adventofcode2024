package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() string {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	return inputString
}
func parseInput(inputString string) ([]int, []int) {
	data := make([]int, len(inputString))
	sum := 0
	for i := range inputString {
		n, e := strconv.Atoi(string(inputString[i]))
		check(e)
		data[i] = n
		sum += n

	}
	p := 0
	id := 0
	outSlice := make([]int, sum)
	for i, n := range data {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				outSlice[p] = id
				p += 1
			}
			id += 1
		} else {
			for j := 0; j < n; j++ {
				outSlice[p] = -1
				p += 1
			}

		}
	}
	return outSlice, data
}

func defrag1(memory []int) int {
	p1 := 0
	p2 := len(memory) - 1
	for {
		for memory[p1] >= 0 {
			p1 += 1
		}
		for memory[p2] == -1 {
			p2 -= 1
		}
		if p2 <= p1 {
			break
		}
		memory[p1] = memory[p2]
		memory[p2] = -1
	}
	// fmt.Println(p1, p2, ":", memory)
	// fmt.Println(memory[p1])
	return p2

}

func checksum1(memory []int, length int) int {
	sum := 0
	// fmt.Println(length, memory)
	for i := 0; i < length+1; i++ {
		sum += memory[i] * i
		// fmt.Println(i, memory[i], sum)
	}
	return sum
}
func copyIntSlice(myInts []int) []int {
	myCopy := make([]int, len(myInts))
	copy(myCopy, myInts)
	return myCopy
}

func defrag2(memory []int, dmap []int) []int {
	filled := make([]int, 0, len(dmap)/2+1)
	empty := make([]int, 0, len(dmap)/2+1)
	filledIndex := make([]int, 0, len(dmap)/2+1)
	emptyIndex := make([]int, 0, len(dmap)/2+1)
	sum := 0
	for i, num := range dmap {
		if i%2 == 0 {
			filled = append(filled, num)
			filledIndex = append(filledIndex, sum)
		} else {
			empty = append(empty, num)
			emptyIndex = append(emptyIndex, sum)
		}
		sum += num
	}

	numBlocks := len(filled)
	p := numBlocks - 1
	for p > 0 {
		// fmt.Println(memory)
		lenCurrentBlock := filled[p]
		actionTaken := false
		for i := 0; i < p; i++ { //iterate over empty slots to put the block in
			if lenCurrentBlock > empty[i] { //it doesn't fit
				continue
			} else { //it fits in empty space i
				for j := 0; j < lenCurrentBlock; j++ {
					memory[emptyIndex[i]+j] = p   //write it there
					memory[filledIndex[p]+j] = -1 //overwrite the original with -1
				}
				emptyIndex[i] += lenCurrentBlock //update the empty space start index
				empty[i] -= lenCurrentBlock      //update the empty space length
				actionTaken = true               //we did something
			}
			if actionTaken { // if we just did something, go to the next block from the right
				break
			}
		}
		p -= 1
	}
	return memory
}
func checksum2(memory []int) int {
	sum := 0
	for i, num := range memory {
		if num != -1 {
			sum += num * i
		}
	}
	return sum
}

func main() {
	startTime := time.Now()
	memory1, dmap := parseInput(readInput())
	memory2 := copyIntSlice(memory1)
	fmt.Println("Data Parsed in", time.Since(startTime))

	p1Time := time.Now()
	length := defrag1(memory1)
	part1 := checksum1(memory1, length)
	fmt.Println("Day 9, Part 1:", part1, ",", time.Since(p1Time))

	p2Time := time.Now()
	defrag2(memory2, dmap)
	part2 := checksum2(memory2)
	fmt.Println("Day 9, Part 2:", part2, time.Since(p2Time))

	fmt.Println(time.Since(startTime))

}
