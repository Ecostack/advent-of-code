package day13

import (
	"AdventOfCode2022/util"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"reflect"
	"sort"
)

type Line any

type Lines []Line

// implement the functions from the sort.Interface
func (d Lines) Len() int {
	return len(d)
}

func (d Lines) Less(i, j int) bool {
	return compare(d[i], d[j])
}

func (d Lines) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func compareValue(val1, val2 Line) int {
	t1 := reflect.TypeOf(val1).String()
	t2 := reflect.TypeOf(val2).String()

	if t1 == "[]interface {}" && t2 == "[]interface {}" {
		for i := 0; i < int(math.Max(float64(len(val1.([]any))), float64(len(val2.([]any))))); i++ {
			if i >= len(val1.([]any)) {
				return 1
			}
			if i >= len(val2.([]any)) {
				return -1
			}

			temp1 := val1.([]any)[i]
			temp2 := val2.([]any)[i]
			res := compareValue(temp1, temp2)
			if res > 0 {
				return 1
			} else if res < 0 {
				return -1
			}
		}
	}

	if t1 == "[]interface {}" && t2 == "float64" {
		temp := []any{val2.(float64)}

		return compareValue(val1, temp)
	}

	if t1 == "float64" && t2 == "[]interface {}" {
		temp := []any{val1.(float64)}

		return compareValue(temp, val2)
	}

	if t1 == "float64" && t2 == "float64" {
		f1 := val1.(float64)
		f2 := val2.(float64)
		if f1 < f2 {
			return 1
		}
		if f1 > f2 {
			return -1
		}
	}
	return 0
}

func compare(arr1, arr2 Line) bool {
	res := compareValue(arr1, arr2)
	if res == 1 {
		return true
	}
	return false
}

func part1Fn(lines []string) {
	array1 := make([]any, 0)
	array2 := make([]any, 0)
	indexes := make([]int, 0)
	pairIndex := 0
	indexSum := 0
	compInternal := func() {
		pairIndex++
		res := compare(array1, array2)
		if res {
			indexSum += pairIndex
			indexes = append(indexes, pairIndex)
		}
		log.Println(res, array1, array2)

		array1 = make([]any, 0)
		array2 = make([]any, 0)
	}

	for i, line := range lines {
		mod3 := (i + 1) % 3
		if mod3 == 1 {
			json.Unmarshal([]byte(line), &array1)
		}
		if mod3 == 2 {
			json.Unmarshal([]byte(line), &array2)
			compInternal()
		}
	}

	log.Println("part1", indexSum, indexes)
}

func part2GetEntries(lines []string) Lines {
	allEntries := make(Lines, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		arr := make([]any, 0)
		json.Unmarshal([]byte(line), &arr)
		allEntries = append(allEntries, arr)
	}

	arr := make([]any, 0)
	json.Unmarshal([]byte("[[2]]"), &arr)
	allEntries = append(allEntries, arr)

	arr = make([]any, 0)
	json.Unmarshal([]byte("[[6]]"), &arr)

	allEntries = append(allEntries, arr)

	return allEntries
}

func printLines(lines Lines) {
	for _, line := range lines {
		log.Println(line)
	}
}

func part2Fn(lines []string) {
	entries := part2GetEntries(lines)
	log.Println(entries)
	sort.Sort(entries)
	printLines(entries)

	sum := 1
	for i, line := range entries {
		newVal := fmt.Sprintf("%v", line)
		if newVal == "[[2]]" || newVal == "[[6]]" {
			sum *= i + 1
		}
	}
	log.Println("part2", sum)
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	if part2 {
		part2Fn(results)
	} else {
		part1Fn(results)
	}

}
