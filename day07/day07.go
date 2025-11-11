package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() ([]int, [][]int) {
	filename := "input.txt"
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

const MaxOperands = 15

type CrunchKey struct {
	PartialAns  int
	Answer      int
	OperandsArr [MaxOperands]int
	length      int
}

func memoizeCrunch1() func(int, int, []int) int {
	cache := make(map[CrunchKey]int)

	var memo func(int, int, []int) int

	memo = func(partialAns int, answer int, operands []int) int {
		var keyArr [MaxOperands]int
		copy(keyArr[:], operands)
		key := CrunchKey{
			PartialAns:  partialAns,
			Answer:      answer,
			OperandsArr: keyArr,
			length:      len(operands),
		}
		if result, found := cache[key]; found {
			return result
		}

		// paste in crunch1 logic lol
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
				matchFound1 = memo(newPartial1, answer, operands[1:])
			}
			if newPartial2 <= answer {
				matchFound2 = memo(newPartial2, answer, operands[1:])
			}
		}

		result := max(matchFound1, matchFound2)

		// end crunch1 logic
		cache[key] = result
		return result
	}
	return memo
}

func memoizeCrunch2() func(int, int, []int) int {
	cache := make(map[CrunchKey]int)

	var memo func(int, int, []int) int

	memo = func(partialAns int, answer int, operands []int) int {
		var keyArr [MaxOperands]int
		copy(keyArr[:], operands)
		key := CrunchKey{
			PartialAns:  partialAns,
			Answer:      answer,
			OperandsArr: keyArr,
			length:      len(operands),
		}
		if result, found := cache[key]; found {
			return result
		}

		// paste in crunch2 logic lol
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

		result := max(matchFound1, matchFound2, matchFound3)

		// end crunch2 logic
		cache[key] = result
		return result
	}
	return memo
}

func sumPossibleAnswersMemo(answers []int, operands [][]int, part int) int {
	crunch1Memo := memoizeCrunch1()
	crunch2Memo := memoizeCrunch2()
	sum_successes := 0
	sum_answers := 0
	yes := 0
	for i := range answers {
		if part == 1 {
			yes = crunch1Memo(operands[i][0], answers[i], operands[i][1:])
		} else {
			yes = crunch2Memo(operands[i][0], answers[i], operands[i][1:])

		}
		sum_successes += yes
		sum_answers += yes * answers[i]
	}
	return sum_answers
}

func sumPossibleAnswersParallel(answers []int, operands [][]int, part int) int {
	var wg sync.WaitGroup

	resultsChannel := make(chan int, len(answers))

	for i := range answers {
		wg.Add(1)
		answerVal := answers[i]
		operandSet := operands[i]

		go func() {
			defer wg.Done()
			yes := 0
			if part == 1 {
				yes = crunch1(operandSet[0], answerVal, operandSet[1:])

			} else {
				yes = crunch2(operandSet[0], answerVal, operandSet[1:])
			}
			resultsChannel <- yes * answerVal
		}()
	}
	go func() {
		wg.Wait()
		close(resultsChannel)
	}()
	finalSumAnswers := 0

	for r := range resultsChannel {
		finalSumAnswers += r
	}
	return finalSumAnswers
}

func main() {
	startTime := time.Now()

	answers, operands := readInput()

	// basic method
	// part1 := sumPossibleAnswers(answers, operands, 1)
	// fmt.Println("Day 7, Part 1:", part1)
	// part2 := sumPossibleAnswers(answers, operands, 2)
	// fmt.Println("Day 7, Part 2:", part2)
	// 155 ms

	// memoization
	// part1 := sumPossibleAnswersMemo(answers, operands, 1)
	// fmt.Println("Day 7, Part 1:", part1)
	// part2 := sumPossibleAnswersMemo(answers, operands, 2)
	// fmt.Println("Day 7, Part 2:", part2)
	//424ms... memoization housekeeping is too expensive.

	// basic + parallelization
	part1 := sumPossibleAnswersParallel(answers, operands, 1)
	fmt.Println("Day 7, Part 1:", part1)
	part2 := sumPossibleAnswersParallel(answers, operands, 2)
	fmt.Println("Day 7, Part 2:", part2)
	// 81 ms - twice as fast!

	fmt.Println(time.Since(startTime))
}
