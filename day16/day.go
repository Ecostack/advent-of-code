package day16

import (
	"AdventOfCode2022/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type System struct {
	valves             map[string]*Valve
	valveCountWithFlow int
	tree               *Tree
	highest            int
	memoBestScore      map[string]int
}

type Valve struct {
	name string
	flow int
	ways []string
}

const MINUTES = 31

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
		name: "",
		flow: 0,
		ways: make([]string, 0),
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
	system.valves[valve.name] = valve
}

func printTree(tree *Tree, layer int) {
	open := tree.open
	openStr := "O"
	if open {
		openStr = "X"
	}
	result := tree.valve.name + "(" + openStr + ")"
	for i := 0; i < layer*2; i++ {
		result = " " + result
	}
	fmt.Println(result)
	//for _, t := range tree.subTree {
	//	printTree(t, layer+1)
	//}
}

func makeKey(state *Tree) string {
	return fmt.Sprintf("%d-%s-%s", state.minute, state.valve.name, state.opened)
}

func buildTree(system *System, tree *Tree, valve *Valve, shouldOpen bool) {
	valvesStrings := valve.ways
	tree.valve = valve

	key := makeKey(tree)
	if _, exists := system.memoBestScore[key]; exists {
		return
		//system.memoBestScore[key] = tree.totalFlow
		//}
	}

	//wg := sync.WaitGroup{}
	//if tree.checkMinutes(system) {
	//	return
	//}
	tree.minute++
	tree.totalFlow += tree.accumulated

	if tree.checkMinutes(system) {
		return
	}

	isOpen := util.Contains(tree.opened, valve.name)

	if valve.flow > 0 && !isOpen && shouldOpen {
		tree.actionLog = append(tree.actionLog, valve.name+"_OPEN")
		tree.totalFlow += tree.accumulated
		tree.minute++
		if tree.checkMinutes(system) {
			return
		}
		//fmt.Println(valve.name + " open valve")
		tree.open = true
		tree.accumulated += valve.flow
		tree.opened = append(tree.opened, valve.name)
	}

	didVisit := false
	for _, valvesString := range valvesStrings {
		visitString := valve.name + " -> " + valvesString

		visitedTree := make(map[string]int)
		if _, exists := tree.visited[visitString]; !exists {
			visitedTree[visitString] = 0
		}

		if system.valveCountWithFlow == len(tree.opened) {
			continue
		}
		if visited, exists := visitedTree[visitString]; exists {
			if visited >= 2 {
				continue
			}
			if visited >= len(valve.ways)-1 {
				continue
			}
		}

		didVisit = true

		valveNew := system.valves[valvesString]
		newTree := &Tree{
			parent:      tree,
			minute:      tree.minute,
			valve:       valveNew,
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
		newTree.actionLog = append(newTree.actionLog, visitString)
		for s, b := range visitedTree {
			newTree.visited[s] = b
		}

		for _, b := range tree.opened {
			newTree.opened = append(newTree.opened, b)
			//newTree.opened[s] = b
		}
		if _, exists := newTree.visited[visitString]; !exists {
			newTree.visited[visitString] = 0
		}
		newTree.visited[visitString]++
		if valveNew.flow > 0 {
			buildTree(system, newTree, valveNew, true)
		}

		newTree = &Tree{
			parent:      tree,
			minute:      tree.minute,
			valve:       valveNew,
			open:        false,
			totalFlow:   tree.totalFlow,
			accumulated: tree.accumulated,
			opened:      make([]string, 0),
			//subTree:     make([]*Tree, 0),
			visited: make(map[string]int),
		}
		for s, b := range visitedTree {
			newTree.visited[s] = b
		}
		//for s, b := range tree.opened {
		//	newTree.opened[s] = b
		//}

		for _, b := range tree.opened {
			newTree.opened = append(newTree.opened, b)
			//newTree.opened[s] = b
		}
		for _, s := range tree.actionLog {
			newTree.actionLog = append(newTree.actionLog, s)
		}
		newTree.actionLog = append(newTree.actionLog, visitString)
		if _, exists := newTree.visited[visitString]; !exists {
			newTree.visited[visitString] = 0
		}
		newTree.visited[visitString]++

		//tree.subTree = append(tree.subTree, newTree)

		//fmt.Println("Moving from " + valve.name + " to  " + valveNew.name + " should NOT open")
		//wg.Add(1)
		//go func() {
		//	defer wg.Done()
		buildTree(system, newTree, valveNew, false)
		//}()

	}
	if !didVisit {
		visitString := valve.name + " -> " + valve.name
		tree.actionLog = append(tree.actionLog, visitString)
		//go func() {
		buildTree(system, tree, valve, true)
		//}

		//buildTree(system, tree, valve, false)
	}

}

func part1Fn(system *System) {
	currentValve := system.valves["AA"]
	tree := &Tree{
		minute:      0,
		valve:       currentValve,
		open:        false,
		totalFlow:   0,
		accumulated: 0,
		opened:      make([]string, 0),
		visited:     make(map[string]int),
		actionLog:   make([]string, 0),
	}
	system.tree = tree
	buildTree(system, tree, currentValve, true)

	log.Println("part1", system.highest)
}

func getValue(file string, example bool, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	system := &System{valves: make(map[string]*Valve), memoBestScore: make(map[string]int)}
	for _, result := range results {
		parseLine(system, result)
	}
	part1Fn(system)
}
