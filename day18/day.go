package day18

import (
	"AdventOfCode2022/util"
	"log"
	"strconv"
	"strings"
)

func getLine(x, y, z string) string {
	return x + "," + y + "," + z
}
func getLineFromInt(x, y, z int) string {
	return getLine(strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
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

const waterMax = 22

func exploreWater(hashmap map[string]bool, x, y, z int) {
	if x < -1 || y < -1 || z < -1 || x > waterMax || y > waterMax || z > waterMax {
		return
	}
	newLine := getLineFromInt(x, y, z)
	if _, exists := hashmap[newLine]; exists {
		return
	}
	hashmap[newLine] = false
	exploreWater(hashmap, x+1, y, z)
	exploreWater(hashmap, x-1, y, z)

	exploreWater(hashmap, x, y+1, z)
	exploreWater(hashmap, x, y-1, z)

	exploreWater(hashmap, x, y, z+1)
	exploreWater(hashmap, x, y, z-1)
}

type SomeVal struct {
	counter int
}

func part2Fn(lines []string, example bool) {
	hashmap := make(map[string]bool)
	for _, line := range lines {
		hashmap[line] = true
	}
	bla := &SomeVal{
		counter: 0,
	}

	exploreWater(hashmap, 0, 0, 0)

	count := 0
	for _, line := range lines {
		split := strings.Split(line, ",")
		x := util.ParseInt(split[0])
		y := util.ParseInt(split[1])
		z := util.ParseInt(split[2])

		if value, exists := hashmap[getLineFromInt(x-1, y, z)]; exists && !value {
			count++
		}
		if value, exists := hashmap[getLineFromInt(x+1, y, z)]; exists && !value {
			count++
		}

		if value, exists := hashmap[getLineFromInt(x, y+1, z)]; exists && !value {
			count++
		}
		if value, exists := hashmap[getLineFromInt(x, y-1, z)]; exists && !value {
			count++
		}

		if value, exists := hashmap[getLineFromInt(x, y, z+1)]; exists && !value {
			count++
		}
		if value, exists := hashmap[getLineFromInt(x, y, z-1)]; exists && !value {
			count++
		}

	}
	log.Println("part2 ", count, bla.counter)
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
