package day16

import (
	"AdventOfCode2022/util"
	"log"
)

const FALL_ORDER = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

type Part int

const EmptyPart Part = 0
const StonePart Part = 1

type Rock [][]Part

func getRocks(lines []string) []Rock {
	result := make([]Rock, 0)
	rock := make(Rock, 0)
	for _, line := range lines {
		if len(line) == 0 {
			result = append(result, rock)
			rock = make(Rock, 0)
			continue
		}
		rockParts := make([]Part, 0)
		for _, i3 := range line {
			if i3 == '.' {
				rockParts = append(rockParts, EmptyPart)
			} else {
				rockParts = append(rockParts, StonePart)
			}
		}
		rock = append(rock, rockParts)
	}
	return result
}

func part1Fn(lines []string) {
	rocks := getRocks(lines)
	log.Println(rocks)
}

func getValue(file string, example bool, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	part1Fn(results)
}
