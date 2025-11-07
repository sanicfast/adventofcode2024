package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Coord struct {
	// Fields representing the X and Y dimensions
	row int
	col int
}
type Guard struct {
	location Coord
	heading  Coord
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readInput() (Guard, []Coord, Coord) {
	filename := "day06_ex.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	rawData, err := os.ReadFile(filename)
	check(err)

	inputString := string(rawData)

	lines := strings.Split(inputString, "\n")
	var obstacles []Coord
	var guard Guard
	for row, line := range lines {
		for col, char := range line {
			switch char {
			case '#':
				obstacles = append(obstacles, Coord{row, col})
			case '^':
				guard = Guard{
					location: Coord{row, col},
					heading:  Coord{-1, 0},
				}
			}
		}
	}
	dim := Coord{len(lines), len(lines[0])}
	return guard, obstacles, dim
}
func (g *Guard) turnRight() {
	oldHeading := g.heading
	var newHeading Coord
	switch oldHeading {
	case Coord{-1, 0}: //north
		newHeading = Coord{0, 1} //east
	case Coord{0, 1}: //east
		newHeading = Coord{1, 0} // south
	case Coord{1, 0}: //south
		newHeading = Coord{0, -1} //west
	case Coord{0, -1}: //west
		newHeading = Coord{-1, 0} //north
	}
	g.heading = newHeading
}

func timeTick(guard *Guard, obstacles []Coord, dim Coord) int {
	newRow := guard.heading.row + guard.location.row
	newCol := guard.heading.col + guard.location.col

	// test bounds. 1 means the guard left.
	if newRow < 0 || newRow >= dim.row {
		return 1
	} else if newCol < 0 || newCol >= dim.row {
		return 1
	}
	// see if we bumped into an obstacle. if so, turn right and exit.
	for _, o := range obstacles {
		if o.row == newRow && o.col == newCol {
			// fmt.Println("bump!", o)
			guard.turnRight()
			return 0
		}
	}
	// we didn't bump into anything! update guard position
	guard.location.row = newRow
	guard.location.col = newCol
	return 0
}

func countGuardLocations(guard Guard, obstacles []Coord, dim Coord) int {
	visitedLocations := make(map[Coord]struct{})
	visitedLocationCount := 0
	finished := 0
	for finished == 0 {
		_, visited := visitedLocations[guard.location]
		if !visited {
			visitedLocationCount += 1
			visitedLocations[guard.location] = struct{}{}
		}

		finished = timeTick(&guard, obstacles, dim)
		// fmt.Println(guard)
	}
	return visitedLocationCount
}

func detectLoop(guard Guard, obstacles []Coord, dim Coord,
	newObstacle Coord, seenGuardStates map[Guard]struct{}) int {
	if guard.location == newObstacle {
		return 0
	}
	obstacles = append(obstacles, newObstacle)

	loopDetected := 0
	finished := 0
	// if the guard ends up in the same exact location with the same exact
	// heading twice we know he's in a loop. if he's not in a loop, he'll
	// escape and we'll see that.
	for finished == 0 && loopDetected == 0 {
		_, visited := seenGuardStates[guard]
		if !visited {
			seenGuardStates[guard] = struct{}{}
		} else {
			loopDetected = 1
		}

		finished = timeTick(&guard, obstacles, dim)
	}
	if loopDetected == 1 {
		return 1
	}
	return 0
}
func countLoopsBruteForce(guard Guard, obstacles []Coord, dim Coord) int {
	sum := 0
	for i := range dim.row {
		for j := range dim.col {
			seenGuardStates := make(map[Guard]struct{})
			sum += detectLoop(guard, obstacles, dim, Coord{i, j}, seenGuardStates)
		}
	}
	return sum
}

func copyGuardMap(original map[Guard]struct{}) map[Guard]struct{} {
	newMap := make(map[Guard]struct{}, len(original))

	for key, _ := range original {
		newMap[key] = struct{}{}
	}
	return newMap
}

func countLoopsSmarter(guard Guard, obstacles []Coord, dim Coord) int {
	initialGuardCoord := guard.location
	sum := 0
	finished := 0
	visitedLocations := make(map[Coord]struct{})
	seenGuardStates := make(map[Guard]struct{})

	// while haven't finished
	for finished == 0 {
		// check if we've been here
		_, visited := visitedLocations[guard.location]
		// if not and it's not the initial location
		if !visited && guard.location != initialGuardCoord {
			// add it to our list of places we've been
			visitedLocations[guard.location] = struct{}{}
			guardClone := copyGuardMap(seenGuardStates)
			sum += detectLoop(guard, obstacles, dim,
				Coord{guard.location.row, guard.location.col}, guardClone,
			)
		}
		// add this guard state to our dictionary
		seenGuardStates[guard] = struct{}{}
		// check if we've finished and if not move forward a tick
		finished = timeTick(&guard, obstacles, dim)
	}
	return sum
}

func main() {
	startTime := time.Now()
	guard, obstacles, dim := readInput()

	part1 := countGuardLocations(guard, obstacles, dim)
	fmt.Println("Day 6, Part 1:", part1)
	part2 := countLoopsBruteForce(guard, obstacles, dim) // 2m15s for brute force!
	fmt.Println("Day 6, Part 2:", part2)
	part2_2 := countLoopsSmarter(guard, obstacles, dim)
	fmt.Println("Day 6, Part 2:", part2_2) // not right yet...

	fmt.Println(time.Since(startTime))

}
