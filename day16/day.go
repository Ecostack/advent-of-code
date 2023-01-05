package day16

import (
	"AdventOfCode2022/util"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type System struct {
	valves             []*Valve
	valveMap           map[string]*Valve
	valveWithFlow      []*Valve
	valveCountWithFlow int
	highest            int
	memoBestScore      map[string]int
	distancePath       [][]int
}

type Valve struct {
	name  string
	index int
	flow  int
	ways  []string
}

const MINUTES = 30

type Tree struct {
	parent      *Tree
	minute      int
	valve       *Valve
	open        bool
	totalFlow   int
	accumulated int
	actionLog   []string
	visited     map[string]int
	opened      []string
}

func (tree *Tree) checkMinutes(system *System) bool {
	if tree.minute == MINUTES {
		if system.highest < tree.totalFlow {
			log.Println("totalFlow  ", tree.totalFlow)
			log.Println("open-> ", tree.opened)
			system.highest = tree.totalFlow
		}
		key := makeKey(tree)
		if best, exists := system.memoBestScore[key]; exists {
			if tree.totalFlow > best {
				system.memoBestScore[key] = tree.totalFlow
			}
		} else {
			system.memoBestScore[key] = tree.totalFlow
		}

		return true
	}
	return false
}

func parseLine(system *System, line string) {
	reg, err := regexp.Compile("Valve ([A-Z]+) has flow rate=([\\d]+); tunnels? leads? to valves? (.*)")
	util.PanicOnError(err)

	valve := &Valve{
		index: 0,
		name:  "",
		flow:  0,
		ways:  make([]string, 0),
	}

	subMatches := reg.FindAllStringSubmatch(line, -1)
	log.Println(subMatches)

	valve.name = subMatches[0][1]
	res, err := strconv.Atoi(subMatches[0][2])
	util.PanicOnError(err)
	valve.flow = res
	split := strings.Split(subMatches[0][3], ", ")
	for _, s := range split {
		valve.ways = append(valve.ways, s)
	}
	if valve.flow > 0 {
		system.valveCountWithFlow++
	}
	valve.index = len(system.valves)
	system.valves = append(system.valves, valve)
	if valve.flow > 0 {
		system.valveWithFlow = append(system.valveWithFlow, valve)
	}

	system.valveMap[valve.name] = valve
}

func makeKey(state *Tree) string {
	return fmt.Sprintf("%d-%s-%s", state.minute, state.valve.name, state.opened)
}

func printMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	for i := 0; i < len(matrix[0]); i++ {
		if i == 0 {
			fmt.Print("  ")
		}
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Print("\n")
	for i, ints := range matrix {
		fmt.Print(i)
		fmt.Print(" ")
		for _, i3 := range ints {
			if i3 == math.MaxInt {
				fmt.Print("-")
			} else {
				fmt.Print(i3)
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func buildMatrix(valves []*Valve) [][]int {
	n := len(valves)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distance := math.MaxInt
			valveI := valves[i]
			valveJ := valves[j]
			if i == j {
				distance = 0
			}
			if util.Contains(valveI.ways, valveJ.name) {
				distance = 1
			}
			matrix[i][j] = distance
		}
	}
	return matrix
}

func nextOptimalValve(system *System, valve *Valve, minutesLeft int, contesters []*Valve) (*Valve, int) {
	var optimalValve *Valve = nil

	value := 0
	for _, v := range contesters {
		if v != valve {
			distance := system.distancePath[valve.index][v.index]
			newTime := minutesLeft - distance - 1
			if newTime <= 0 {
				continue
			}
			score := newTime * v.flow

			newConstesters := make([]*Valve, 0)
			for _, contester := range contesters {
				if contester != v && contester != valve {
					newConstesters = append(newConstesters, contester)
				}
			}

			_, temp := nextOptimalValve(system, v, newTime, newConstesters)
			score += temp

			if score > value {
				optimalValve = v
				value = score
			}
		}
	}
	return optimalValve, value
}

func part1Fn(system *System) {
	matrix := buildMatrix(system.valves)
	dp := util.FloydWarshall(matrix)
	printMatrix(dp)

	system.distancePath = dp

	currentValve := system.valveMap["AA"]

	valve, score := nextOptimalValve(system, currentValve, 30, system.valveWithFlow)
	log.Println("v, s", valve, score)
}

func combinationHelper(combinations [][]int, data []int, start, end, index int) [][]int {
	if index == len(data) {
		clone := make([]int, len(data))
		for i, datum := range data {
			clone[i] = datum
		}
		return append(combinations, clone)
	} else if start <= end {
		data[index] = start
		temp := combinationHelper(combinations, data, start+1, end, index+1)
		return combinationHelper(temp, data, start+1, end, index)
	}
	return combinations
}

func combinations(n, r int) [][]int {
	return combinationHelper(make([][]int, 0), make([]int, r), 0, n-1, 0)
}

func getPermutations[T any](val []T) [][]int {
	result := make([][]int, 0)
	for i := 1; i < len(val); i++ {
		temp := combinations(len(val), i)
		for _, ints := range temp {
			result = append(result, ints)
		}
	}
	return result
}

func part2Fn(system *System) {
	matrix := buildMatrix(system.valves)
	dp := util.FloydWarshall(matrix)
	printMatrix(dp)
	system.distancePath = dp

	currentValve := system.valveMap["AA"]

	combinations := getPermutations(system.valveWithFlow)
	log.Println("comb", combinations)
	highest := 0
	for _, comb := range combinations {
		temp := make([]*Valve, 0)
		temp2 := make([]*Valve, 0)
		for i, valve := range system.valveWithFlow {
			if util.Contains(comb, i) {
				temp = append(temp, valve)
			} else {
				temp2 = append(temp2, valve)
			}
		}
		_, score1 := nextOptimalValve(system, currentValve, 26, temp)
		_, score2 := nextOptimalValve(system, currentValve, 26, temp2)
		scoreComb := score1 + score2
		if highest < scoreComb {
			highest = scoreComb
		}
	}

	log.Println("part2", highest)
}

func getValue(file string, example bool, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	system := &System{
		valves:        make([]*Valve, 0),
		valveWithFlow: make([]*Valve, 0),
		valveMap:      make(map[string]*Valve),
		memoBestScore: make(map[string]int)}
	for _, result := range results {
		parseLine(system, result)
	}

	if part2 {
		part2Fn(system)
	} else {
		part1Fn(system)
	}
}
