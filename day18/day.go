package day18

import (
	"AdventOfCode2022/util"
	"log"
	"strconv"
	"strings"
)

func getLine(x, y, z string) string {
	//return strconv.Itoa(x)+","+strconv.Itoa(y)+","+strconv.Itoa(z)
	return x + "," + y + "," + z
}

func part1Fn(lines []string, example bool) {
	hashmap := make(map[string]bool)
	for _, line := range lines {
		hashmap[line] = true
	}
	count := 0
	for _, line := range lines {
		hasNeighbor := false
		split := strings.Split(line, ",")

		if _, exists := hashmap[getLine(strconv.Itoa(util.ParseInt(split[0])-1), split[1], split[2])]; !exists {
			count++
		} else {
			hasNeighbor = true
		}
		if _, exists := hashmap[getLine(strconv.Itoa(util.ParseInt(split[0])+1), split[1], split[2])]; !exists {
			count++
		} else {
			hasNeighbor = true
		}
		if _, exists := hashmap[getLine(split[0], strconv.Itoa(util.ParseInt(split[1])-1), split[2])]; !exists {
			count++
		} else {
			hasNeighbor = true
		}
		if _, exists := hashmap[getLine(split[0], strconv.Itoa(util.ParseInt(split[1])+1), split[2])]; !exists {
			count++
		} else {
			hasNeighbor = true
		}
		if _, exists := hashmap[getLine(split[0], split[1], strconv.Itoa(util.ParseInt(split[2])-1))]; !exists {
			count++
		} else {
			hasNeighbor = true
		}
		if _, exists := hashmap[getLine(split[0], split[1], strconv.Itoa(util.ParseInt(split[2])+1))]; !exists {
			count++
		} else {
			hasNeighbor = true
		}
		if !hasNeighbor {
			log.Print("no neighbor: ", line)
		}
	}
	log.Println("part1 ", count)
}

func part2Fn(lines []string, example bool) {
	hashmap := make(map[string]bool)
	for _, line := range lines {
		hashmap[line] = true
	}
	waterMax := 22
	for x := 0; x < waterMax; x++ {
		for y := 0; x < waterMax; y++ {
			for z := 0; x < waterMax; z++ {

			}
		}
	}
}

func getValue(file string, example bool, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	part1Fn(results, example)
}
