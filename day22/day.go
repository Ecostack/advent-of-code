package day5

import (
	"AdventOfCode2022/util"
)

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

}
