package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	r int
	c int
}

type Stack []Coord

func (stack *Stack) Push(c Coord) {
	*stack = append(*stack, c)
}

func (stack *Stack) Pop() Coord {
	// // queue Dequeue logic
	// out := (*queue)[0]
	// *queue = (*queue)[1:]
	// return out

	// stack Pop logic
	n := len(*stack)
	out := (*stack)[n-1]
	*stack = (*stack)[:n-1]
	return out
}

func readInput() ([][]int, Coord) {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	volcanoMap := make([][]int, len(lines))

	for r, line := range lines {
		for _, char := range line {
			if char == '.' {
				volcanoMap[r] = append(volcanoMap[r], -2)
				continue
			}
			num, e1 := strconv.Atoi(string(char))
			check(e1)
			volcanoMap[r] = append(volcanoMap[r], num)
		}
	}
	dim := Coord{len(lines), len(lines[0])}
	return volcanoMap, dim
}

func printVolcanoMap(volcanoMap [][]int) {
	for _, line := range volcanoMap {
		for _, num := range line {
			if num == -2 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", num)
			}
		}
		fmt.Printf(("\n"))
	}
}
func inBounds(c Coord, dim Coord) bool {
	if c.r >= 0 && c.r < dim.r && c.c >= 0 && c.c < dim.c {
		return true
	}
	return false
}

func getNumLocs(volcanoMap [][]int, findNum int) []Coord {
	out := []Coord{}
	for r, line := range volcanoMap {
		for c, num := range line {
			if num == findNum {
				out = append(out, Coord{r, c})
			}
		}
	}
	return out
}

func Traverse(volcanoMap [][]int, dim Coord) (int, int) {
	theStack := make(Stack, 0, 100)
	zeros := getNumLocs(volcanoMap, 0)
	sumReachableNines := 0
	validPathsCnt := 0

	for _, c := range zeros {
		reachableNinesSet := make(map[Coord]struct{})
		theStack.Push(c)
		for len(theStack) > 0 {
			loc := theStack.Pop()
			height := volcanoMap[loc.r][loc.c]
			cardinal_dirs := []Coord{
				{loc.r + 1, loc.c},
				{loc.r - 1, loc.c},
				{loc.r, loc.c + 1},
				{loc.r, loc.c - 1}}

			for _, newCoord := range cardinal_dirs {
				if inBounds(newCoord, dim) {
					if volcanoMap[newCoord.r][newCoord.c] == height+1 {
						if volcanoMap[newCoord.r][newCoord.c] == 9 {
							validPathsCnt += 1
							reachableNinesSet[newCoord] = struct{}{}
						} else {
							theStack.Push(newCoord)
						}
					}
				}
			}
		}
		sumReachableNines += len(reachableNinesSet)
	}
	return sumReachableNines, validPathsCnt
}

func main() {
	startTime := time.Now()
	volcanoMap, dim := readInput()
	// printVolcanoMap(volcanoMap)
	part1, part2 := Traverse(volcanoMap, dim)
	fmt.Println("Day 10, Part 1:", part1)
	fmt.Println("Day 10, Part 2:", part2)
	fmt.Println(time.Since(startTime))
}
