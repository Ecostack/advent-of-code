package day5

import (
	"AdventOfCode2022/util"
	"log"
)

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	total := 0
	for _, line := range results {

		log.Println(line)
	}
	log.Println(total)
}
