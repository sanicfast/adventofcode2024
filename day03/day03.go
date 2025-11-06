package main

import (
	"fmt"
	"os"
	"strconv"
)

var intCharSet = map[byte]struct{}{
	'0': {},
	'1': {},
	'2': {},
	'3': {},
	'4': {},
	'5': {},
	'6': {},
	'7': {},
	'8': {},
	'9': {},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() string {
	filename := "day03.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	rawData, err := os.ReadFile(filename)
	check(err)

	inputString := string(rawData)

	return inputString
}

func parse(input string, part int) int {
	i := 0
	var numChar1, numChar2 []byte
	var numArgs int
	outSum := 0
	do := true
	for i < len(input) {
		if part == 2 {
			// fmt.Println(input[i:min(i+4, len(input)-1)])
			if input[i:min(i+4, len(input)-1)] == "do()" {
				do = true
				i += 4
				// fmt.Println(i, "do!")
			} else if input[i:min(i+7, len(input)-1)] == "don't()" {
				do = false
				// fmt.Println(i, "don't!")
				i += 7
			}

		}
		if input[i:min(i+4, len(input)-1)] == "mul(" && do {
			// fmt.Println(i)
			j := i + 4
			numChar1, numChar2, numArgs = nil, nil, 0
		jLoop:
			for j < len(input) {
				// if we see a number, add it to the appropriate argument slice
				if _, isNum := intCharSet[input[j]]; isNum {
					switch numArgs {
					case 0:
						numChar1 = append(numChar1, input[j])
					case 1:
						numChar2 = append(numChar2, input[j])
					case 2:
						panic("3 arguments detected")
					}
				} else if input[j] == ',' {
					// if we see a comma, increment argument counter.
					if numArgs == 0 && len(numChar1) == 0 {
						i = j
						break jLoop
					}
					numArgs++
					if numArgs > 1 {
						i = j
						break jLoop
					}
				} else if input[j] == ')' {
					// if we see a close paren, check if we have data in both args
					if len(numChar1) > 0 && len(numChar2) > 0 {
						arg1, e1 := strconv.Atoi(string(numChar1))
						arg2, e2 := strconv.Atoi(string(numChar2))
						check(e1)
						check(e2)
						// fmt.Println("mult", arg1, arg2, arg1*arg2)
						outSum += arg1 * arg2
						i = j
						break jLoop
					} else {
						i = j
						break jLoop
					}
				} else {
					i = j
					break jLoop
				}
				j += 1
			}
		}
		i += 1
	}
	return outSum
}

func main() {
	input := readInput()
	// fmt.Println(input)
	part1 := parse(input, 1)
	fmt.Println("Day 3, Part 1:", part1)
	part2 := parse(input, 2)
	fmt.Println("Day 3, Part 2:", part2)
}
