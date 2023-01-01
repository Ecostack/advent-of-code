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
	valves []*Valve
}

type Valve struct {
	name string
	flow int
	ways []string
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
	system.valves = append(system.valves, valve)
}

func getValue(file string, example bool, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	system := &System{valves: make([]*Valve, 0)}
	for _, result := range results {
		parseLine(system, result)
	}
	fmt.Println(system)
}
