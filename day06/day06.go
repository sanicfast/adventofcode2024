package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Coord struct {
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
func readInput() (Guard, map[Coord]struct{}, Coord) {
	filename := "day06.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	rawData, err := os.ReadFile(filename)
	check(err)

	inputString := string(rawData)

	lines := strings.Split(inputString, "\n")
	obstacles := make(map[Coord]struct{})
	var guard Guard
	for row, line := range lines {
		for col, char := range line {
			switch char {
			case '#':
				obstacles[Coord{row, col}] = struct{}{}
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

func timeTick(guard *Guard, obstacles map[Coord]struct{}, dim Coord) int {
	newRow := guard.heading.row + guard.location.row
	newCol := guard.heading.col + guard.location.col

	// test bounds. 1 means the guard left.
	if newRow < 0 || newRow >= dim.row {
		return 1
	} else if newCol < 0 || newCol >= dim.row {
		return 1
	}
	// see if we bumped into an obstacle. if so, turn right and exit.
	// for _, o := range obstacles {
	// 	if o.row == newRow && o.col == newCol {
	// 		guard.turnRight()
	// 		return 0
	// 	}
	_, bump := obstacles[Coord{newRow, newCol}]
	if bump {
		guard.turnRight()
		return 0
	}

	// we didn't bump into anything! update guard position
	guard.location.row = newRow
	guard.location.col = newCol
	return 0
}

func countGuardLocations(guard Guard, obstacles map[Coord]struct{}, dim Coord) int {
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
	}
	return visitedLocationCount
}

func detectLoop(guard Guard, obstacles map[Coord]struct{}, dim Coord,
	seenGuardStates map[Guard]struct{}) int {
	// if guard.location == newObstacle {
	// 	return 0
	// }

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
func countLoopsBruteForce(guard Guard, obstacles map[Coord]struct{}, dim Coord) int {
	initialGuardCoord := guard.location
	sum := 0
	for i := range dim.row {
		for j := range dim.col {
			obstacleLocation := Coord{i, j}
			if obstacleLocation == initialGuardCoord {
				continue
			}
			seenGuardStates := make(map[Guard]struct{})
			_, bump := obstacles[Coord{i, j}]
			if !bump {
				obstacles[Coord{i, j}] = struct{}{}

				loopDetected := detectLoop(guard, obstacles, dim, seenGuardStates)
				sum += loopDetected
				delete(obstacles, Coord{i, j})
			}
		}
	}
	return sum
}

func copyGuardMap(original map[Guard]struct{}) map[Guard]struct{} {
	newMap := make(map[Guard]struct{}, len(original))

	for key, val := range original {
		newMap[key] = val
	}
	return newMap
}
func copyCoordMap(original map[Coord]struct{}) map[Coord]struct{} {
	newMap := make(map[Coord]struct{}, len(original))

	for key, val := range original {
		newMap[key] = val
	}
	return newMap
}

func countLoopsSmarter(guard Guard, obstacles map[Coord]struct{}, dim Coord) int {
	initialGuardCoord := guard.location
	sum := 0
	finished := 0
	visitedLocations := make(map[Coord]struct{})
	seenGuardStates := make(map[Guard]struct{})

	// while haven't finished
	for finished == 0 {
		// drop an obstacle in front of the guard and see if it created a loop
		newObstacle := Coord{
			guard.location.row + guard.heading.row,
			guard.location.col + guard.heading.col}

		_, alreadyTried := visitedLocations[newObstacle]
		_, bump := obstacles[newObstacle]

		if !alreadyTried && newObstacle != initialGuardCoord && !bump {
			guardMapClone := copyGuardMap(seenGuardStates)
			obstacles[newObstacle] = struct{}{}
			loopDetected := detectLoop(guard, obstacles, dim,
				guardMapClone,
			)
			sum += loopDetected
			delete(obstacles, newObstacle)

		}
		// add it to our list of places we've been
		visitedLocations[guard.location] = struct{}{}
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

	// part2 := countLoopsBruteForce(guard, obstacles, dim) // 2m15s for brute force! 36s if we put the obstacles in a map
	// fmt.Println("Day 6, Part 2:", part2)

	part2_2 := countLoopsSmarter(guard, obstacles, dim)
	fmt.Println("Day 6, Part 2:", part2_2)

	fmt.Println(time.Since(startTime))

}
