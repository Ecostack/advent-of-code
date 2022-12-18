package day4

import (
	"AdventOfCode2022/util"
	"log"
	"strconv"
	"strings"
)

func getRange(ranges string) []int {
	numbers := strings.Split(ranges, "-")
	n1 := numbers[0]
	n2 := numbers[1]
	val1, _ := strconv.Atoi(n1)
	val2, _ := strconv.Atoi(n2)
	result := make([]int, 0)
	for i := val1; i <= val2; i++ {
		result = append(result, i)
	}
	return result
}

func arrContains[T comparable](super []T, sub []T) bool {
	if len(sub) == 0 || len(super) == 0 {
		return false
	}
	for _, value := range sub {
		result := util.Contains(super, value)
		if !result {
			return false
		}
	}
	return true
}

func arrContainsAtLeastOne[T comparable](super []T, sub []T) bool {
	if len(sub) == 0 || len(super) == 0 {
		return false
	}
	for _, value := range sub {
		result := util.Contains(super, value)
		if result {
			return true
		}
	}
	return false
}

func isPairs(line string) bool {
	splits := strings.Split(line, ",")
	range1 := getRange(splits[0])
	range2 := getRange(splits[1])

	if arrContains(range1, range2) {
		return true
	}

	if arrContains(range2, range1) {
		return true
	}
	return false
}

func isPairsAtLeastOne(line string) bool {
	splits := strings.Split(line, ",")
	range1 := getRange(splits[0])
	range2 := getRange(splits[1])

	if arrContainsAtLeastOne(range1, range2) {
		return true
	}

	if arrContainsAtLeastOne(range2, range1) {
		return true
	}
	return false
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	total := 0
	for _, line := range results {
		val := false
		if part2 {
			val = isPairsAtLeastOne(line)
		} else {
			val = isPairs(line)
		}
		if val {
			total++
		}
		//log.Println(val)
	}
	log.Println(total)
}
