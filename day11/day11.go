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

func readInput() (LinkedList, map[int]int) {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	rawData, err := os.ReadFile(filename)
	check(err)
	inputString := string(rawData)
	nums := strings.Fields(inputString)
	mySlice := make([]int, len(nums))
	stonesMap := make(map[int]int, 0)
	for i, numC := range nums {
		n, e := strconv.Atoi(numC)
		check(e)
		stonesMap[n] += 1
		mySlice[i] = n
	}
	myLL := makeLinkedList(mySlice)
	return myLL, stonesMap
}

type Node struct {
	n    int
	Next *Node
	Prev *Node
}
type LinkedList struct {
	Head *Node
	Tail *Node
}

func makeLinkedList(slice []int) LinkedList {
	myLL := LinkedList{}
	for i, n := range slice {
		newNode := &Node{n: n}
		if i == 0 {
			myLL.Head = newNode
		} else {
			myLL.Tail.Next = newNode
			newNode.Prev = myLL.Tail
		}
		myLL.Tail = newNode
	}
	return myLL
}

func displayLL(ll LinkedList) {
	node := ll.Head
	for node != nil {
		fmt.Println(node.n)
		node = node.Next
	}
}

func (ll *LinkedList) Len() int {
	node := ll.Head
	length := 0
	for node != nil {
		length += 1
		node = node.Next
	}
	return length
}

func blinkLL(ll LinkedList) LinkedList {
	node := ll.Head
	for node != nil {
		if node.n == 0 {
			node.n = 1
			node = node.Next
			continue
		}
		strNum := strconv.Itoa(node.n)
		numDigits := len(strNum)
		if numDigits%2 == 0 {
			firstHalf, e1 := strconv.Atoi(strNum[0 : numDigits/2])
			secondHalf, e2 := strconv.Atoi(strNum[numDigits/2:])
			check(e1)
			check(e2)
			node.n = firstHalf
			newNode := &Node{n: secondHalf, Prev: node, Next: node.Next}
			if node.Next != nil {
				node.Next.Prev = newNode
			} else {
				ll.Tail = newNode
			}
			node.Next = newNode
			node = node.Next
		} else {
			node.n = node.n * 2024
		}
		node = node.Next
	}
	return ll
}

func blinkNTimesLL(stonesLL LinkedList, n int) int {
	for range n {
		stonesLL = blinkLL(stonesLL)
	}
	return stonesLL.Len()
}

func blinkMap(stonesMap map[int]int) map[int]int {
	newMap := make(map[int]int, len(stonesMap))
	for k, v := range stonesMap {
		if k == 0 {
			newMap[1] += v
			continue
		}

		strNum := strconv.Itoa(k)
		numDigits := len(strNum)
		if numDigits%2 == 0 {
			firstHalf, e1 := strconv.Atoi(strNum[0 : numDigits/2])
			secondHalf, e2 := strconv.Atoi(strNum[numDigits/2:])
			check(e1)
			check(e2)
			newMap[firstHalf] += v
			newMap[secondHalf] += v
			continue
		}
		newMap[k*2024] += v
	}
	return newMap
}
func sumVals(stonesMap map[int]int) int {
	sum := 0
	for _, v := range stonesMap {
		sum += v
	}
	return sum
}

func blinkNTimesMap(stonesMap map[int]int, n int) int {
	for range n {
		stonesMap = blinkMap(stonesMap)
	}
	sum := sumVals(stonesMap)
	return sum
}
func main() {
	startTime := time.Now()
	_, stonesMap := readInput()

	// stonesLL, stonesMap := readInput()
	// p1LLTime := time.Now()
	// part1LL := blinkNTimesLL(stonesLL, 25)
	// fmt.Println("Day 11, Part 1:", part1LL, time.Since(p1LLTime))

	p1MapTime := time.Now()
	part1Map := blinkNTimesMap(stonesMap, 25)
	fmt.Println("Day 11, Part 1:", part1Map, time.Since(p1MapTime))

	p2MapTime := time.Now()
	part2Map := blinkNTimesMap(stonesMap, 75)
	fmt.Println("Day 11, Part 2:", part2Map, time.Since(p2MapTime))
	fmt.Println(time.Since(startTime))
}

// I did this in nvim!
