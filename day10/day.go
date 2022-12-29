package day10

import (
	"AdventOfCode2022/util"
	"log"
	"strconv"
	"strings"
)

func part1Fn() {

}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	x := 1
	cycle := 0
	evalOutputs := []int{20, 60, 100, 140, 180, 220}
	result := 0
	for _, line := range results {
		split := strings.Split(line, " ")
		cmd := split[0]
		if cmd == "addx" {
			toAdd, _ := strconv.Atoi(split[1])
			for i := 0; i < 2; i++ {
				cycle++
				if util.Contains(evalOutputs, cycle) {
					result += cycle * x
				}
			}
			x += toAdd
		}
		if cmd == "noop" {
			cycle++
			if util.Contains(evalOutputs, cycle) {
				result += cycle * x
			}
		}
	}
	log.Println("part1", result)

}
