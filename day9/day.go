package day9

import (
	"AdventOfCode2022/util"
	"fmt"
	"gopkg.in/karalabe/cookiejar.v2/exts/mathext"
	"log"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	row  int
	col  int
	tail *Position
}

func printVisited(pos map[string]bool) {
	for i := -100; i < 100; i++ {
		for j := -50; j < 50; j++ {
			key := strconv.Itoa(i*-1) + ":" + strconv.Itoa(j)
			if _, exists := pos[key]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func printPositions(pos *Position) {

	for i := -20; i < 20; i++ {
		for j := -20; j < 20; j++ {
			counter := 0
			temp := pos
			var onPosition *Position = nil
			for temp != nil {
				if temp.row == i*-1 && temp.col == j {
					onPosition = temp
					break
				}

				temp = temp.tail
				counter++
				continue
			}
			if onPosition != nil {
				if counter == 0 {
					fmt.Print("H")
				} else {
					fmt.Print(counter)
				}
			} else {
				if i == 0 && j == 0 {
					fmt.Print("s")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func part1Fn(lines []string) {
	first := &Position{
		row: 0,
		col: 0,
	}

	temp := &Position{
		row:  0,
		col:  0,
		tail: nil,
	}
	first.tail = temp
	for i := 1; i < 9; i++ {
		new := &Position{
			row:  0,
			col:  0,
			tail: nil,
		}
		temp.tail = new
		temp = new
	}

	visited := make(map[string]bool)

	evaluatePositions := func() {
		head := first
		tail := head.tail
		for tail != nil {
			changed := false
			if head.row == tail.row {
				if math.Abs(float64(head.col-tail.col)) >= 2 {
					changed = true
					newVal := mathext.SignInt(head.col - tail.col)
					tail.col = head.col - int(newVal)
				}
			}
			if head.col == tail.col {
				if math.Abs(float64(head.row-tail.row)) >= 2 {
					changed = true
					newVal := mathext.SignInt(head.row - tail.row)
					tail.row = head.row - int(newVal)
				}
			}
			if head.col != tail.col && head.row != tail.row {
				if math.Abs(float64(head.col-tail.col)) >= 2 && math.Abs(float64(head.row-tail.row)) >= 2 {
					changed = true
					newVal := mathext.SignInt(head.col - tail.col)
					newVal2 := mathext.SignInt(head.row - tail.row)
					tail.col = head.col - int(newVal)
					tail.row = head.row - newVal2
				}
				if math.Abs(float64(head.col-tail.col)) >= 2 {
					changed = true
					newVal := mathext.SignInt(head.col - tail.col)
					tail.col = head.col - int(newVal)
					tail.row = head.row
				}
				if math.Abs(float64(head.row-tail.row)) >= 2 {
					changed = true
					newVal := mathext.SignInt(head.row - tail.row)
					tail.row = head.row - int(newVal)
					tail.col = head.col

				}
			}
			if !changed {
				return
			}
			head = tail
			tail = head.tail

		}
	}

	printPositions(first)

	visit := func() {
		temp := first
		for temp.tail != nil {
			temp = temp.tail
		}
		key := strconv.Itoa(temp.row) + ":" + strconv.Itoa(temp.col)
		visited[key] = true
	}

	visit()

	for _, line := range lines {
		split := strings.Split(line, " ")
		cmdCount, _ := strconv.Atoi(split[1])
		for i := 0; i < cmdCount; i++ {
			if split[0] == "U" {
				first.row += 1
			}
			if split[0] == "D" {
				first.row -= 1
			}
			if split[0] == "L" {
				first.col -= 1
			}
			if split[0] == "R" {
				first.col += 1
			}
			evaluatePositions()
			//printPositions(first)
			visit()
		}

	}

	printVisited(visited)

	log.Println("part1 ", len(visited))
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	part1Fn(results)

}
