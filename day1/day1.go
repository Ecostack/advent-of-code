package day1

import (
	"AdventOfCode2022/util"
	"log"
	"sort"
	"strconv"
)

func getValue(file string) (int, int) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	calories := make([]int, 0)

	temp := 0
	highestCalories := 0
	topThree := 0
	for _, result := range results {
		if len(result) == 0 {
			calories = append(calories, temp)
			if highestCalories < temp {
				highestCalories = temp
			}
			temp = 0
			continue
		}
		parsed, err := strconv.Atoi(result)
		util.PanicOnError(err)
		temp = temp + parsed
	}

	caloriesTemp := calories[:]
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesTemp)))
	log.Println("highest calories: ", highestCalories)
	log.Println("top three calories: ", caloriesTemp[0]+caloriesTemp[1]+caloriesTemp[2])
	return highestCalories, topThree
}
