package day12

import (
	"AdventOfCode2022/util"
	"fmt"
	"gopkg.in/karalabe/cookiejar.v2/collections/queue"
	"log"
	"strconv"
)

type Map struct {
	mapVal    [][]int32
	distances map[string]int
	smallest  int
}

func getMap(lines []string) *Map {
	result := &Map{
		mapVal:    make([][]int32, 0),
		distances: make(map[string]int),
		smallest:  -1,
	}
	for _, line := range lines {
		temp := make([]int32, 0)
		for _, char := range line {
			if char == 'S' {
				temp = append(temp, 0)
				continue
			}
			if char == 'E' {
				temp = append(temp, 27)
				continue
			}
			temp = append(temp, char-96)
		}
		result.mapVal = append(result.mapVal, temp)
	}
	return result
}

func traverse(myQueue *queue.Queue, myMap *Map, history []string, r, c int) {
	rowLen := len(myMap.mapVal)
	colLen := len(myMap.mapVal[0])
	key := makeKey(r, c)
	keyOld := ""
	if len(history) > 0 {
		keyOld = history[len(history)-1]
	}

	historyNew := append(history, key)

	valueCurrent := myMap.mapVal[r][c]

	if valueCurrent == 27 {
		if myMap.smallest == -1 || len(historyNew) < myMap.smallest {
			myMap.smallest = len(historyNew) - 1
			log.Println("found one", myMap.smallest)
		}
		return
	}
	if myMap.smallest != -1 && myMap.smallest < len(historyNew)-1 {
		return
	}

	traverseInternal := func(rNew, cNew int) {
		keyNew := makeKey(rNew, cNew)
		_, exists := myMap.distances[keyNew]
		if exists {
			return
		}
		if keyNew != keyOld {
			value := myMap.mapVal[rNew][cNew]
			if (value) == valueCurrent || value-1 == valueCurrent || value < valueCurrent {
				myMap.distances[keyNew] = len(historyNew)
				myQueue.Push(func() {
					traverse(myQueue, myMap, historyNew, rNew, cNew)
				})
			}
		}
	}

	if (r + 1) < rowLen {
		traverseInternal(r+1, c)
	}
	if (c + 1) < colLen {
		traverseInternal(r, c+1)
	}
	if (r - 1) >= 0 {
		traverseInternal(r-1, c)
	}
	if (c - 1) >= 0 {
		traverseInternal(r, c-1)
	}
}

func makeKey(r, c int) string {
	return strconv.Itoa(r) + ":" + strconv.Itoa(c)
}

func printMap(myMap *Map) {
	rowLen := len(myMap.mapVal)
	colLen := len(myMap.mapVal[0])
	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			if myMap.mapVal[r][c] == 27 {
				fmt.Print("E  ")
			}
			myKey := makeKey(r, c)
			if distance, exists := myMap.distances[myKey]; exists {
				padded := "00" + strconv.Itoa(distance)
				if distance > 9 {
					padded = "0" + strconv.Itoa(distance)
				}
				if distance > 99 {
					padded = strconv.Itoa(distance)
				}
				fmt.Print(padded)
				fmt.Print(" ")
			} else {
				fmt.Print(".   ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func calcFromMap(myMap *Map, startR, startC int) int {
	history := make([]string, 0)
	myQueue := queue.New()
	keyNew := makeKey(startR, startC)
	myMap.distances[keyNew] = 0
	traverse(myQueue, myMap, history, startR, startC)
	for !myQueue.Empty() {
		myQueue.Pop().(func())()
	}
	return myMap.smallest
}

func part1Fn(myMap *Map) {
	rowLen := len(myMap.mapVal)
	colLen := len(myMap.mapVal[0])
	startR := 0
	startC := 0

	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			if myMap.mapVal[r][c] == 0 {
				myMap.mapVal[r][c] = 1
				startR = r
				startC = c
				break
			}
		}
	}

	calcFromMap(myMap, startR, startC)

	log.Println(len(myMap.distances), rowLen*colLen)
	log.Println("part1", myMap.smallest)
}

func part2Fn(myMap *Map) {
	rowLen := len(myMap.mapVal)
	colLen := len(myMap.mapVal[0])

	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			if myMap.mapVal[r][c] == 0 {
				myMap.mapVal[r][c] = 1
			}
		}
	}

	smallest := -1
	for r := 0; r < rowLen; r++ {
		for c := 0; c < colLen; c++ {
			if myMap.mapVal[r][c] == 1 {
				newMap := &Map{
					mapVal:    myMap.mapVal[:],
					distances: make(map[string]int),
					smallest:  -1,
				}
				calcFromMap(newMap, r, c)
				if newMap.smallest >= 0 {
					if smallest == -1 || smallest > newMap.smallest {
						smallest = newMap.smallest
					}
				}
			}
		}
	}
	log.Println("part2", smallest)
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	mapValue := getMap(results)
	if part2 {
		part2Fn(mapValue)
	} else {
		part1Fn(mapValue)
	}
}
