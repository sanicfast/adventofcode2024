package main

import (
	"fmt"
	"os"
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

func readInput() (map[rune][]Coord, Coord) {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	lines := strings.Split(inputString, "\n")
	antennas := make(map[rune][]Coord, 10)
	for r, line := range lines {
		for c, char := range line {
			if char == '.' {
				continue
			} else {
				antennas[char] = append(antennas[char], Coord{r, c})
			}
		}
	}
	dim := Coord{len(lines), len(lines[0])}
	return antennas, dim
}
func viewAntennas(antennas map[rune][]Coord) {
	for k, v := range antennas {
		fmt.Println(string(k), v)
	}
}
func getAntinodesPair(a1 Coord, a2 Coord) [2]Coord {
	// vector pointing from a1 to a2
	v := Coord{a2.r - a1.r, a2.c - a1.c}
	output := [2]Coord{}
	// subtract v from a1 and add v to a2
	output[0] = Coord{a1.r - v.r, a1.c - v.c}
	output[1] = Coord{a2.r + v.r, a2.c + v.c}
	return output
}
func processAntennasFreq1(inputSlice []Coord) []Coord {
	countAntinodes := 0
	for i := range inputSlice {
		countAntinodes += i
	}
	antinodes := make([]Coord, 0, countAntinodes*2)
	for i := range inputSlice {
		for j := range inputSlice {
			if i > j {
				pair := getAntinodesPair(inputSlice[i], inputSlice[j])
				// fmt.Println(inputSlice[i], inputSlice[j], pair)
				antinodes = append(antinodes, pair[0], pair[1])
			}
		}
	}
	return antinodes
}

func processFreqs1(antennas map[rune][]Coord, dim Coord, plotEm bool) int {
	antinodes := []Coord{}
	for _, v := range antennas {
		newAntinodes := processAntennasFreq1(v)
		antinodes = append(antinodes, newAntinodes...)
	}
	if plotEm {
		plotAntinodes(antinodes, dim)
	}
	// deduplicate
	antinodeSet := make(map[Coord]struct{})
	for _, a := range antinodes {
		// bounds check
		if inBounds(a, dim) {
			antinodeSet[a] = struct{}{}
		}
	}
	return len(antinodeSet)
}

func inBounds(c Coord, dim Coord) bool {
	if c.r >= 0 && c.r < dim.r && c.c >= 0 && c.c < dim.c {
		return true
	}
	return false
}

func getAntinodesResonant(a1 Coord, a2 Coord, dim Coord) []Coord {
	// vector pointing from a1 to a2
	v := Coord{a2.r - a1.r, a2.c - a1.c}
	antinodes := []Coord{}
	// subtract v from a1 and add v to a2 repeatedly until out of bounds
	for inBounds(a1, dim) {
		antinodes = append(antinodes, a1)
		a1 = Coord{a1.r - v.r, a1.c - v.c}
	}
	for inBounds(a2, dim) {
		antinodes = append(antinodes, a2)
		a2 = Coord{a2.r + v.r, a2.c + v.c}
	}
	return antinodes
}

func processAntennasFreq2(inputSlice []Coord, dim Coord) []Coord {

	antinodes := []Coord{}
	for i := range inputSlice {
		for j := range inputSlice {
			if i > j {
				a := getAntinodesResonant(inputSlice[i], inputSlice[j], dim)
				antinodes = append(antinodes, a...)
			}
		}
	}
	return antinodes
}

func processFreqs2(antennas map[rune][]Coord, dim Coord, plotEm bool) int {
	antinodes := []Coord{}
	for _, v := range antennas {
		newAntinodes := processAntennasFreq2(v, dim)
		antinodes = append(antinodes, newAntinodes...)
	}
	if plotEm {
		plotAntinodes(antinodes, dim)
	}
	// deduplicate
	antinodeSet := make(map[Coord]struct{})
	for _, a := range antinodes {
		antinodeSet[a] = struct{}{}
	}
	return len(antinodeSet)
}
func plotAntinodes(a []Coord, dim Coord) {
	plot := make([][]int, dim.r)
	for i := range plot {
		plot[i] = make([]int, dim.c)
	}
	for _, c := range a {
		if inBounds(c, dim) {
			plot[c.r][c.c] += 1
		}
	}
	for i := range plot {
		fmt.Println(plot[i])
	}
}

func main() {
	startTime := time.Now()

	antennas, dim := readInput()
	// viewAntennas(antennas)

	part1 := processFreqs1(antennas, dim, false)
	fmt.Println("Day 8, Part 1:", part1)
	part2 := processFreqs2(antennas, dim, false)
	fmt.Println("Day 8, Part 2:", part2)

	fmt.Println(time.Since(startTime))

}

// could get rid of some hashmaps to make this more efficient but it's 1ms so eh
