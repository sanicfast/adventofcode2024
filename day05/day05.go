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
func readInput() (map[int][]int, [][]int) {
	// read in the input file and returns a rule dict
	// with keys as the gt num and vals of the lt nums and
	// a list of the updates to be checked.
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	rawData, err := os.ReadFile(filename)
	check(err)

	inputString := string(rawData)
	splitData := strings.Split(inputString, "\n\n")

	rulesChar := strings.Split(splitData[0], "\n")
	updatesChar := strings.Split(splitData[1], "\n")

	var rules [][]int
	for i := range rulesChar {
		splitRule := strings.Split(rulesChar[i], "|")
		if len(splitRule) != 2 {
			panic("Too many numbers on one rule!")
		}
		num1, e1 := strconv.Atoi(splitRule[0])
		num2, e2 := strconv.Atoi(splitRule[1])
		check(e1)
		check(e2)
		rules = append(rules, []int{num1, num2})

	}

	ruleMap := make(map[int][]int, len(rules))
	for _, rule := range rules {
		ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
	}

	var updates [][]int
	for i := range updatesChar {
		splitUpdate := strings.Split(updatesChar[i], ",")
		var updateLine []int
		for _, numChar := range splitUpdate {
			numInt, e := strconv.Atoi(numChar)
			check(e)
			updateLine = append(updateLine, numInt)
		}
		updates = append(updates, updateLine)
	}

	return ruleMap, updates
}
func updateIsOrdered(ruleMap map[int][]int, update []int) bool {
	for i := len(update) - 1; i > 0; i-- {
		checkNum := update[i]
		numsAbove := update[:i]
		forbiddenNums, found := ruleMap[checkNum]
		if !found {
			continue
		}
		for _, numAbove := range numsAbove {
			for _, forbiddenNum := range forbiddenNums {
				if numAbove == forbiddenNum {
					return false
				}
			}
		}
	}
	return true
}
func sumMiddleNumsOfCorrectlyOrderedUpdates(ruleMap map[int][]int, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		if updateIsOrdered(ruleMap, update) {
			sum += update[len(update)/2]
		}
	}
	return sum
}

func sortUpdate(ruleMap map[int][]int, update []int) []int {
	for {
		swapped := false
		for i := len(update) - 1; i > 0; i-- {
			checkNum := update[i]
			numsAbove := update[:i]
			forbiddenNums, found := ruleMap[checkNum]
			if !found {
				continue
			}
		jloop:
			for j, numAbove := range numsAbove {
				for _, forbiddenNum := range forbiddenNums {
					if numAbove == forbiddenNum {
						// if we make a swap, the list wasn't sorted right
						// and we need to double check if it's sorted so we
						// don't return it and let our infinite loop continue
						// this is apparently similar to a bubble sort.
						update[i] = numAbove
						update[j] = checkNum
						swapped = true
						// we don't wan't to fiddle with this checknum anymore,
						// having swapped it so we break out of the loop checking it.
						break jloop
					}
				}
			}
		}
		if !swapped {
			return update
		}
	}
}
func sumMiddleNumsOfInorrectlyOrderedUpdates(ruleMap map[int][]int, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		if !updateIsOrdered(ruleMap, update) {
			update = sortUpdate(ruleMap, update)
			middleIndex := len(update) / 2
			middleNum := update[middleIndex]
			sum += middleNum
		}
	}
	return sum
}

func main() {
	startTime := time.Now()

	ruleMap, updates := readInput()
	part1 := sumMiddleNumsOfCorrectlyOrderedUpdates(ruleMap, updates)
	fmt.Println("Day 5, Part 1:", part1)
	part2 := sumMiddleNumsOfInorrectlyOrderedUpdates(ruleMap, updates)
	fmt.Println("Day 5, Part 2:", part2)
	fmt.Println(time.Since(startTime))

}
