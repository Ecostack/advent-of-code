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

func copyTree(tree *Tree) *Tree {
	newTree := &Tree{
		parent:      tree,
		minute:      tree.minute,
		valve:       tree.valve,
		open:        false,
		totalFlow:   tree.totalFlow,
		accumulated: tree.accumulated,
		actionLog:   make([]string, 0),
		opened:      make([]string, 0),
		visited:     make(map[string]int),
	}
	for _, s := range tree.actionLog {
		newTree.actionLog = append(newTree.actionLog, s)
	}
	for s, b := range tree.visited {
		newTree.visited[s] = b
	}
	for _, b := range tree.opened {
		newTree.opened = append(newTree.opened, b)
	}
	return newTree
}

func nextOptimalValve(system *System, valve *Valve, minute int, contesters []*Valve) (*Valve, int) {
	//timeLeft := MINUTES - minute
	var optimalValve *Valve = nil
	value := 0
	for _, v := range contesters {
		if v != valve {
			distance := system.distancePath[valve.index][v.index]
			newTime := minute + distance + 1
			if newTime >= MINUTES {
				continue
			}
			score := newTime * v.flow

			newConstesters := make([]*Valve, 0)
			for _, contester := range contesters {
				if contester != v {
					newConstesters = append(newConstesters, contester)
				}
			}

			_, value := nextOptimalValve(system, v, newTime, newConstesters)
			score += value

			if score > value {
				optimalValve = v
				value = score
			}
		}
	}
	return optimalValve, value
}

func buildTree(system *System, tree *Tree) {
	valve := tree.valve

	isOpen := util.Contains(tree.opened, valve.name)

	if !util.Contains(tree.opened, tree.valve.name) && !isOpen && util.Contains(system.valveWithFlow, valve) {
		newTree := copyTree(tree)
		newTree.minute++
		newTree.totalFlow += tree.accumulated
		if newTree.checkMinutes(system) {
			return
		}
		newTree.accumulated += valve.flow
		newTree.open = true
		newTree.opened = append(newTree.opened, valve.name)
		newTree.actionLog = append(newTree.actionLog, valve.name+"_OPEN")
		buildTree(system, newTree)
	}

	foundUnexplored := false
	for _, v := range system.valveWithFlow {
		if !util.Contains(tree.opened, v.name) && v != valve {
			foundUnexplored = true
			newTree := copyTree(tree)
			newTree.valve = v
			visitString := valve.name + " -> " + v.name
			newTree.actionLog = append(newTree.actionLog, visitString)
			distance := system.distancePath[valve.index][v.index]
			for i := 0; i < distance; i++ {
				newTree.minute++
				newTree.totalFlow += newTree.accumulated
				newTree.actionLog = append(newTree.actionLog, visitString)
				if newTree.checkMinutes(system) {
					return
				}
			}
			buildTree(system, newTree)
		}
	}

	if !foundUnexplored {
		newTree := copyTree(tree)
		visitString := valve.name + " -> " + valve.name

		for {
			newTree.actionLog = append(newTree.actionLog, visitString)
			newTree.minute++
			newTree.totalFlow += newTree.accumulated
			if newTree.checkMinutes(system) {
				return
			}
		}
	}
}

func part1Fn(system *System) {
	matrix := buildMatrix(system.valves)
	//printMatrix(matrix)
	dp := util.FloydWarshall(matrix)
	printMatrix(dp)

	system.distancePath = dp

	currentValve := system.valveMap["AA"]
	//tree := &Tree{
	//	minute:      0,
	//	valve:       currentValve,
	//	open:        false,
	//	totalFlow:   0,
	//	accumulated: 0,
	//	opened:      make([]string, 0),
	//	visited:     make(map[string]int),
	//	actionLog:   make([]string, 0),
	//}
	//buildTree(system, tree)

	valve, score := nextOptimalValve(system, currentValve, 0, system.valveWithFlow)
	log.Println("v, s", valve, score)
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
	part1Fn(system)
}
