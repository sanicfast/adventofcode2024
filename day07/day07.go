package main

import (
	"fmt"
	"math"
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

func readInput() ([]int, [][]int) {
	filename := "day07_ex.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")

	answers := make([]int, len(lines))
	operands := make([][]int, len(lines))
	for i, line := range lines {
		splitLine := strings.Split(line, ": ")
		answer, e1 := strconv.Atoi(splitLine[0])
		check(e1)
		operandsString := strings.Split(splitLine[1], " ")
		operandsList := make([]int, len(operandsString))
		for j, o := range operandsString {
			num, e2 := strconv.Atoi(o)
			check(e2)
			operandsList[j] = num
		}
		answers[i] = answer
		operands[i] = operandsList
	}
	return answers, operands
}

func crunch1(partialAns int, answer int, operands []int) int {
	nextNum := operands[0]
	newPartial1 := partialAns + nextNum
	newPartial2 := partialAns * nextNum
	matchFound1 := 0
	matchFound2 := 0
	// fmt.Println(partialAns, nextNum, ",", newPartial1, newPartial2, operands[1:])

	if len(operands) == 1 {
		// if we get a match we return 1
		if newPartial1 == answer || newPartial2 == answer {
			// fmt.Println(answer, "found!")
			return 1
		}
	} else {
		if newPartial1 <= answer {
			matchFound1 = crunch1(newPartial1, answer, operands[1:])
		}
		if newPartial2 <= answer {
			matchFound2 = crunch1(newPartial2, answer, operands[1:])
		}
	}

	return max(matchFound1, matchFound2)
}
func sumPossibleAnswers(answers []int, operands [][]int, part int) int {

	sum_successes := 0
	sum_answers := 0
	yes := 0
	for i := range answers {
		if part == 1 {
			yes = crunch1(operands[i][0], answers[i], operands[i][1:])
		} else {
			yes = crunch2(operands[i][0], answers[i], operands[i][1:])

		}
		sum_successes += yes
		sum_answers += yes * answers[i]
	}
	return sum_answers
}

func concatInts(num1 int, num2 int) int {
	digits := int(math.Log10(float64(num2))) + 1
	num1_2 := num1 * int(math.Pow10(digits))
	out := num1_2 + num2
	return out
}

func crunch2(partialAns int, answer int, operands []int) int {

	nextNum := operands[0]
	newPartial1 := partialAns + nextNum
	newPartial2 := partialAns * nextNum
	newPartial3 := concatInts(partialAns, nextNum)
	matchFound1, matchFound2, matchFound3 := 0, 0, 0
	// fmt.Println(partialAns, nextNum, ",", newPartial1, newPartial2, operands[1:])

	if len(operands) == 1 {
		// if we get a match we return 1
		if newPartial1 == answer || newPartial2 == answer || newPartial3 == answer {
			// fmt.Println(answer, "found!")
			return 1
		}
	} else {
		if newPartial1 <= answer {
			matchFound1 = crunch2(newPartial1, answer, operands[1:])
		}
		if newPartial2 <= answer {
			matchFound2 = crunch2(newPartial2, answer, operands[1:])
		}
		if newPartial3 <= answer {
			matchFound3 = crunch2(newPartial3, answer, operands[1:])
		}
	}

	return max(matchFound1, matchFound2, matchFound3)
}

func main() {
	startTime := time.Now()

	answers, operands := readInput()

	part1 := sumPossibleAnswers(answers, operands, 1)
	fmt.Println("Day 7, Part 1:", part1)
	part2 := sumPossibleAnswers(answers, operands, 2)
	fmt.Println("Day 7, Part 2:", part2)
	//155ms

	fmt.Println(time.Since(startTime))
}
