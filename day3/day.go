package day3

import (
	"AdventOfCode2022/util"
	"log"
)

func getPointByInt(val int32) int {
	if val >= 97 {
		return int(val - 96)
	}
	if val >= 65 {
		return int(val - 38)
	}
	return 0
}

func findSimilarityASCII(lines []string) int32 {
	temp := make(map[int32]map[int]bool)
	for i, line := range lines {
		for _, ascii := range line {
			if _, exists := temp[ascii]; !exists {
				temp[ascii] = make(map[int]bool)
			}
			temp[ascii][i] = false
		}
	}
	for ascii, val := range temp {
		if len(val) == len(lines) {
			return ascii
		}
	}
	return 0
}

func splitHalf(val string) int {
	half1 := val[:len(val)/2]
	half2 := val[len(val)/2:]
	similarityAscii := findSimilarityASCII([]string{
		half1, half2,
	})
	return getPointByInt(similarityAscii)
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	total := 0
	lines := make([]string, 0)

	for _, result := range results {
		if part2 {
			lines = append(lines, result)
			if len(lines) == 3 {
				similarityAscii := findSimilarityASCII(lines)
				points := getPointByInt(similarityAscii)
				total += points
				lines = make([]string, 0)
			}
		} else {
			val := splitHalf(result)
			log.Println(result, val)
			total += val
		}
	}

	log.Println("total ", total)
}
