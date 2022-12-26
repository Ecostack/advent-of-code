package day6

import (
	"AdventOfCode2022/util"
	"log"
)

func getUniqueEntries(toCheck []rune) int {
	s := make(map[rune]bool)

	for _, r := range toCheck {
		s[r] = true
	}
	return len(s)
}

func findFirstUniqueCharacters(line string, marker int) int {
	last := make([]rune, 0)
	for i, char := range line {
		if len(last) == marker {
			last = last[1:]
		}
		last = append(last, char)

		if len(last) == marker {
			uniques := getUniqueEntries(last)
			if uniques == marker {
				return i + 1
			}
		}
	}
	return 0
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	result := findFirstUniqueCharacters(results[0], 14)

	log.Println(result)
}
