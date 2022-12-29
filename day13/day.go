package day13

import (
	"AdventOfCode2022/util"
	"encoding/json"
	"log"
)

func compare(arr1, arr2 []any) int {
	result := -1
	for i := 0; i < len(arr1); i++ {

	}
	//for i, a := range arr1 {
	//	t := reflect.TypeOf(a)
	//
	//}
	return result
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	array1 := make([]any, 0)
	array2 := make([]any, 0)
	indexes := make([]int, 0)
	for i, line := range results {
		mod3 := (i + 1) % 3
		if len(line) == 0 {
			res := compare(array1, array2)
			if res >= 0 {
				indexes = append(indexes, res)
			}
			//compare
			log.Println("1", array1, array2)

			array1 = make([]any, 0)
			array2 = make([]any, 0)
			continue
		}
		if mod3 == 1 {
			json.Unmarshal([]byte(line), &array1)
		}
		if mod3 == 2 {
			json.Unmarshal([]byte(line), &array2)
		}
	}
}
