package day8

import (
	"AdventOfCode2022/util"
	"log"
	"strconv"
)

func isTreeVisible(trees [][]int, r int, c int) bool {
	rowLen := len(trees)
	colLen := len(trees[r])
	val := trees[r][c]

	topVis := true
	botVis := true
	leftVis := true
	rightVis := true

	for i := r - 1; i >= 0; i-- {
		temp := trees[i][c]
		if temp >= val {
			topVis = false
		}
	}

	if topVis {
		return true
	}

	// bottom
	for i := r + 1; i < rowLen; i++ {
		temp := trees[i][c]
		if temp >= val {
			botVis = false
		}
	}
	if botVis {
		return true
	}

	// left
	for i := c - 1; i >= 0; i-- {
		temp := trees[r][i]
		if temp >= val {
			leftVis = false
		}
	}
	if leftVis {
		return true
	}

	// right
	for i := c + 1; i < colLen; i++ {
		temp := trees[r][i]
		if temp >= val {
			rightVis = false
		}
	}
	if rightVis {
		return true
	}
	return false
}

func part1Fn(trees [][]int) {
	total := 0
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			vis := isTreeVisible(trees, i, j)
			if vis {
				total += 1
			}

			//log.Println("i,j", i, j, vis)
		}
	}
	total += len(trees) * 2
	total += len(trees[0])*2 - 4

	log.Println("part1 ", total)
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	trees := make([][]int, 0)
	for i, line := range results {
		trees = append(trees, make([]int, 0))
		for _, char := range line {
			val, _ := strconv.Atoi(string(char))
			trees[i] = append(trees[i], val)
		}
	}

	part1Fn(trees)
}
