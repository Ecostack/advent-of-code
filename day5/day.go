package day5

import (
	"AdventOfCode2022/util"
	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
	"log"
	"strconv"
	"strings"
)

const STACK_LENGTH = 9

//const STACK_LENGTH = 3

func parseFirst(line string) []uint8 {
	results := make([]uint8, STACK_LENGTH)
	for i := 0; i < len(line); i = i + 4 {
		val := line[i]
		if val == ' ' {
			continue
		}
		results[i/4] = line[i+1]
	}
	return results
}

type Movement struct {
	from  uint
	to    uint
	count uint
}

func parseSecond(line string) *Movement {
	temp := strings.Split(line, " ")
	from, _ := strconv.Atoi(temp[3])
	to, _ := strconv.Atoi(temp[5])
	count, _ := strconv.Atoi(temp[1])
	return &Movement{
		from:  uint(from) - 1,
		to:    uint(to) - 1,
		count: uint(count),
	}
}

func printStacks(stacks []*stack.Stack) {
	result := ""
	for _, s := range stacks {
		if s.Empty() || s.Top() == nil {
			result = result + " "
			continue
		}
		result = result + string(s.Top().(uint8))
	}
	log.Println(result)
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	stacks := make([]*stack.Stack, STACK_LENGTH)
	for i, _ := range stacks {
		stacks[i] = stack.New()
	}

	for i := 7; i >= 0; i-- {
		//for i := 3; i >= 0; i-- {
		line := results[i]
		parsed := parseFirst(line)
		for i2, u := range parsed {
			if u > 0 {
				stacks[i2].Push(u)
			}

		}
		printStacks(stacks)
	}

	for _, s := range stacks {
		log.Println(s)
	}

	for i, line := range results {
		//if i < 5 {
		if i < 10 {
			continue
		}
		move := parseSecond(line)
		log.Println(move)
		if part2 {
			temp := stack.New()
			for i := uint(0); i < move.count; i++ {
				temp.Push(stacks[move.from].Pop())
			}
			for !temp.Empty() {
				stacks[move.to].Push(temp.Pop())
			}
		} else {
			for i := uint(0); i < move.count; i++ {
				stacks[move.to].Push(stacks[move.from].Pop())
			}
		}
	}

	printStacks(stacks)
	//result := ""
	//for _, s := range stacks {
	//	result = result + string(s.Top().(uint8))
	//}
	//log.Println(result)
}
