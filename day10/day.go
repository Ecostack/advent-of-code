package day10

import (
	"AdventOfCode2022/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const CRT_X_MAX = 40
const CRT_Y_MAX = 6

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	x := 1
	cycle := 0
	evalOutputs := []int{20, 60, 100, 140, 180, 220}
	result := 0

	crtX := 0
	crtY := 0

	printCRT := func() {
		if crtX == x || crtX == x-1 || crtX == x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}

	cycleFn := func() {
		cycle++
		if util.Contains(evalOutputs, cycle) {
			result += cycle * x
		}

		printCRT()
		crtX++
		if crtX == CRT_X_MAX {
			crtX = 0
			crtY++
			fmt.Print("\n")
		}
		if crtY == CRT_Y_MAX {
			crtY = 0
		}

	}

	for _, line := range results {
		split := strings.Split(line, " ")
		cmd := split[0]
		if cmd == "addx" {
			toAdd, _ := strconv.Atoi(split[1])
			for i := 0; i < 2; i++ {
				cycleFn()
			}
			x += toAdd
		}
		if cmd == "noop" {
			cycleFn()
		}
	}
	log.Println("part1", result)

}
