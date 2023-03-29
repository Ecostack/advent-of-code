package day18

import (
	"AdventOfCode2022/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Blueprint struct {
	id  int
	ore struct {
		ore int
	}
	clay struct {
		ore int
	}
	obsidian struct {
		ore  int
		clay int
	}
	geode struct {
		ore      int
		obsidian int
	}
}

func readOre(line string) int {
	return readValue(line, "([\\d]+) ore")
}

func readClay(line string) int {
	return readValue(line, "([\\d]+) clay")
}

func readObsidian(line string) int {
	return readValue(line, "([\\d]+) obsidian")
}

func readValue(line string, pattern string) int {
	// compile the pattern into a regular expression object
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regular expression:", err)
		panic(err)
	}

	// use the regular expression to find matches in a string
	matches := regex.FindStringSubmatch(line)
	if len(matches) > 0 {
		//fmt.Println("Found a match for the regular expression:", matches[0])
		res, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		return res
	} else {
		fmt.Println("No matches found for the regular expression")
		panic("No matches found for the regular expression")
	}
	panic("No matches found")
	return -1
}

func parseLineToBlueprint(line string) *Blueprint {
	//blueprint 1: Each Ore robot costs 4 Ore. Each Clay robot costs 2 Ore. Each Obsidian robot costs 3 Ore and 14 Clay. Each Geode robot costs 2 Ore and 7 Obsidian.
	split := strings.Split(line, ":")
	id := strings.Split(split[0], " ")
	each := strings.Split(split[1], ".")

	ore := readOre(each[0])

	clay := readOre(each[1])

	obsOre := readOre(each[2])
	obsClay := readClay(each[2])

	geoOre := readOre(each[3])
	geoObs := readObsidian(each[3])

	return &Blueprint{
		id: util.ParseInt(id[1]),
		ore: struct {
			ore int
		}{
			ore: ore,
		},
		clay: struct {
			ore int
		}{
			ore: clay,
		},
		obsidian: struct {
			ore  int
			clay int
		}{
			ore:  obsOre,
			clay: obsClay,
		},
		geode: struct {
			ore      int
			obsidian int
		}{
			ore:      geoOre,
			obsidian: geoObs,
		},
	}
}

type State struct {
	masterState      *MasterState `json:"master_state,omitempty"`
	Minute           int          `json:"Minute,omitempty"`
	RobotOre         int          `json:"robot_ore,omitempty"`
	RobotClay        int          `json:"robot_clay,omitempty"`
	RobotObsidian    int          `json:"robot_obsidian,omitempty"`
	RobotGeode       int          `json:"robot_geode,omitempty"`
	BuildingOre      int          `json:"building_ore,omitempty"`
	BuildingClay     int          `json:"building_clay,omitempty"`
	BuildingObsidian int          `json:"building_obsidian,omitempty"`
	BuildingGeode    int          `json:"building_geode,omitempty"`
	Ore              int          `json:"Ore,omitempty"`
	Clay             int          `json:"Clay,omitempty"`
	Obsidian         int          `json:"Obsidian,omitempty"`
	Geode            int          `json:"Geode,omitempty"`
	log              []string
	blueprint        *Blueprint `json:"blueprint,omitempty"`
}

type MasterState struct {
	maxGeode           map[int]*State
	maxGeodeWithString map[string]bool
	robotGeodeMinute   map[string]int
}

func (s *State) toString() string {
	//return fmt.Sprintf("blue-%v min-%v rO-%v rC-%v rObs-%v rG-%v,%v,%v,%v,%v", s.blueprint.id, s.Minute, s.RobotOre, s.RobotClay, s.RobotObsidian, s.RobotGeode, s.Ore, s.Clay, s.Obsidian, s.Geode)
	return fmt.Sprintf("blue-%v min-%v rO-%v rC-%v rObs-%v rG-%v", s.blueprint.id, s.Minute, s.RobotOre, s.RobotClay, s.RobotObsidian, s.RobotGeode)
}

func (s *State) clone() *State {
	newLog := make([]string, len(s.log))
	copy(newLog, s.log)
	return &State{
		Minute:           s.Minute,
		RobotOre:         s.RobotOre,
		RobotClay:        s.RobotClay,
		RobotObsidian:    s.RobotObsidian,
		RobotGeode:       s.RobotGeode,
		BuildingOre:      s.BuildingOre,
		BuildingClay:     s.BuildingClay,
		BuildingObsidian: s.BuildingObsidian,
		BuildingGeode:    s.BuildingGeode,
		Ore:              s.Ore,
		Clay:             s.Clay,
		Obsidian:         s.Obsidian,
		Geode:            s.Geode,
		blueprint:        s.blueprint,
		masterState:      s.masterState,
		log:              newLog,
	}
}

func build(state *State) {
	//fmt.Println("\n== Minute " + strconv.Itoa(state.Minute+1) + " ==")
	state.log = append(state.log, fmt.Sprintf("\n== Minute "+strconv.Itoa(state.Minute+1)+" =="))

	blueprint := state.blueprint
	// Geode
	if blueprint.geode.ore <= state.Ore && blueprint.geode.obsidian <= state.Obsidian {

		//robotGeodeString := fmt.Sprintf("b%v robot-%v", state.blueprint.id, state.RobotGeode+1)
		//oldVal, exists := state.MasterState.robotGeodeMinute[robotGeodeString]
		//if !exists || (exists && oldVal > state.Minute+1) {
		//state.MasterState.robotGeodeMinute[robotGeodeString] = state.Minute + 1
		clone2 := state.clone()
		clone2.Minute++
		clone2.RobotGeode++
		if _, exists := state.masterState.maxGeodeWithString[clone2.toString()]; !exists {

			clone := state.clone()
			clone.BuildingGeode++
			clone.Ore = clone.Ore - blueprint.geode.ore
			clone.Obsidian = clone.Obsidian - blueprint.geode.obsidian
			clone.log = append(clone.log, fmt.Sprintf("Spend X to start building a geode-collecting robot."))
			runInternalProceed(clone)
		}

		//}
	}
	if blueprint.obsidian.ore <= state.Ore && blueprint.obsidian.clay <= state.Clay {
		if state.RobotObsidian < state.blueprint.geode.obsidian {
			clone2 := state.clone()
			clone2.Minute++
			clone2.RobotObsidian++
			if _, exists := state.masterState.maxGeodeWithString[clone2.toString()]; !exists {
				clone := state.clone()
				clone.BuildingObsidian++
				clone.Ore = clone.Ore - blueprint.obsidian.ore
				clone.Clay = clone.Clay - blueprint.obsidian.clay
				clone.log = append(clone.log, fmt.Sprintf("Spend X to start building a obsidian-collecting robot."))
				runInternalProceed(clone)
			}
		}
	}
	if blueprint.clay.ore <= state.Ore {
		if state.RobotClay < state.blueprint.obsidian.clay {
			clone2 := state.clone()
			clone2.Minute++
			clone2.RobotClay++
			if _, exists := state.masterState.maxGeodeWithString[clone2.toString()]; !exists {
				clone := state.clone()
				clone.BuildingClay++
				clone.Ore = clone.Ore - blueprint.clay.ore
				clone.log = append(clone.log, fmt.Sprintf("Spend X to start building a clay-collecting robot."))
				runInternalProceed(clone)
			}
		}
	}
	if blueprint.ore.ore <= state.Ore {
		if state.RobotOre < state.blueprint.obsidian.ore || state.RobotOre < state.blueprint.geode.ore || state.RobotOre < state.blueprint.clay.ore {
			clone2 := state.clone()
			clone2.Minute++
			clone2.RobotOre++
			if _, exists := state.masterState.maxGeodeWithString[clone2.toString()]; !exists {
				clone := state.clone()
				clone.BuildingOre++
				clone.Ore = clone.Ore - blueprint.ore.ore
				clone.log = append(clone.log, fmt.Sprintf("Spend X to start building a ore-collecting robot."))
				runInternalProceed(clone)
			}
		}
	}
	runInternalProceed(state.clone())
}

func collect(state *State) {
	state.Obsidian = state.Obsidian + (1 * state.RobotObsidian)
	state.Geode = state.Geode + (1 * state.RobotGeode)
	state.Clay = state.Clay + (1 * state.RobotClay)
	state.Ore = state.Ore + (1 * state.RobotOre)

	if state.RobotOre > 0 {
		state.log = append(state.log, fmt.Sprintf("%v ore-collecting robot collects %v ore; you now have %v ore.\n", state.RobotOre, state.RobotOre, state.Ore))
		//fmt.Printf("%v ore-collecting robot collects %v ore; you now have %v ore.\n", state.RobotOre, state.RobotOre, state.Ore)
	}
	if state.RobotObsidian > 0 {
		state.log = append(state.log, fmt.Sprintf("%v obsidian-collecting robot collects %v obsidian; you now have %v obsidian.\n", state.RobotObsidian, state.RobotObsidian, state.Obsidian))
		//fmt.Printf("%v obsidian-collecting robot collects %v obsidian; you now have %v obsidian.\n", state.RobotObsidian, state.RobotObsidian, state.Obsidian)
	}
	if state.RobotClay > 0 {
		state.log = append(state.log, fmt.Sprintf("%v clay-collecting robot collects %v clay; you now have %v clay.\n", state.RobotClay, state.RobotClay, state.Clay))
		//fmt.Printf("%v clay-collecting robot collects %v clay; you now have %v clay.\n", state.RobotClay, state.RobotClay, state.Clay)
	}
	if state.RobotGeode > 0 {
		state.log = append(state.log, fmt.Sprintf("%v geode-collecting robot collects %v geode; you now have %v geode.\n", state.RobotGeode, state.RobotGeode, state.Geode))
		//fmt.Printf("%v geode-collecting robot collects %v geode; you now have %v geode.\n", state.RobotGeode, state.RobotGeode, state.Geode)
	}

	//== Minute 3 ==
	//Spend 2 ore to start building a clay-collecting robot.
	//1 ore-collecting robot collects 1 ore; you now have 1 ore.
	//The new clay-collecting robot is ready; you now have 1 of them.
}

func collectBuild(state *State) {
	state.RobotGeode = state.RobotGeode + (state.BuildingGeode)
	state.RobotObsidian = state.RobotObsidian + (state.BuildingObsidian)
	state.RobotOre = state.RobotOre + (state.BuildingOre)
	state.RobotClay = state.RobotClay + (state.BuildingClay)

	if state.BuildingOre > 0 {
		state.log = append(state.log, fmt.Sprintf("The new ore-collecting robot is ready; you now have %v of them.\n", state.RobotOre))
		//fmt.Printf("The new ore-collecting robot is ready; you now have %v of them.\n", state.RobotOre)
	}
	if state.BuildingObsidian > 0 {
		state.log = append(state.log, fmt.Sprintf("The new obsidian-collecting robot is ready; you now have %v of them.\n", state.RobotObsidian))
		//fmt.Printf("The new obsidian-collecting robot is ready; you now have %v of them.\n", state.RobotObsidian)
	}
	if state.BuildingClay > 0 {
		state.log = append(state.log, fmt.Sprintf("The new clay-collecting robot is ready; you now have %v of them.\n", state.RobotClay))
		//fmt.Printf("The new clay-collecting robot is ready; you now have %v of them.\n", state.RobotClay)
	}
	if state.BuildingGeode > 0 {
		state.log = append(state.log, fmt.Sprintf("The new geode-collecting robot is ready; you now have %v of them.\n", state.RobotGeode))
		//fmt.Printf("The new geode-collecting robot is ready; you now have %v of them.\n", state.RobotGeode)
	}

	state.BuildingObsidian = 0
	state.BuildingClay = 0
	state.BuildingOre = 0
	state.BuildingGeode = 0
}

func runInternalProceed(state *State) {
	collect(state)
	collectBuild(state)
	state.Minute++
	//state.masterState.maxGeodeWithString[state.toString()] = true
	if state.Minute == 20 {
		if oldVal, exists := state.masterState.maxGeode[state.blueprint.id]; exists {
			if oldVal.Geode < state.Geode {
				state.masterState.maxGeode[state.blueprint.id] = state
			}
		} else {
			state.masterState.maxGeode[state.blueprint.id] = state
		}

		return
	}
	build(state)
}

func run(blueprint *Blueprint) {
	ms := &MasterState{
		maxGeode:           make(map[int]*State),
		maxGeodeWithString: make(map[string]bool),
		robotGeodeMinute:   make(map[string]int),
	}
	state := &State{
		RobotOre:    1,
		blueprint:   blueprint,
		log:         make([]string, 0),
		masterState: ms,
		Minute:      0,
	}

	build(state)
	fmt.Printf("ID %v, Geode %v\n", blueprint.id, ms.maxGeode[blueprint.id])
}

func part1Fn(lines []string, example bool) {

	blueprints := make([]*Blueprint, len(lines))
	for i, line := range lines {
		blueprints[i] = parseLineToBlueprint(line)
		fmt.Println("\n== blueprint: ", blueprints[i].id)
		run(blueprints[i])
	}
}

func part2Fn(lines []string, example bool) {

}

func getValue(file string, example bool, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	if part2 {
		part2Fn(results, example)
	} else {
		part1Fn(results, example)
	}
}
