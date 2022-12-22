package day5

import (
	"AdventOfCode2022/util"
	"log"
	"strconv"
	"strings"
)

type LineParsed struct {
	name       string
	value      *int
	operator   *string
	leftValue  *string
	rightValue *string
}

func parseLine(line string) *LineParsed {
	parsed := strings.Split(line, ": ")

	result := &LineParsed{
		name:       parsed[0],
		value:      nil,
		operator:   nil,
		leftValue:  nil,
		rightValue: nil,
	}

	secondPartSplit := strings.Split(parsed[1], " ")
	if len(secondPartSplit) == 3 {
		result.leftValue = &secondPartSplit[0]
		result.operator = &secondPartSplit[1]
		result.rightValue = &secondPartSplit[2]
	} else {
		temp, err := strconv.Atoi(secondPartSplit[0])
		if err != nil {
			panic(err)
		}
		result.value = &temp
	}

	return result
}

func operatorResult(operator string, left int, right int) *int {
	if operator == "*" {
		temp := left * right
		return &temp
	}
	if operator == "+" {
		temp := left + right
		return &temp
	}
	if operator == "-" {
		temp := left - right
		return &temp
	}
	if operator == "/" {
		temp := left / right
		return &temp
	}
	return nil
}

func operatorResultInvert(operator string, left int, right int) *int {
	if operator == "*" {
		temp := right / left
		return &temp
	}
	if operator == "+" {
		temp := right - left
		return &temp
	}
	if operator == "-" {
		temp := left + right
		return &temp
	}
	if operator == "/" {
		temp := left * right
		return &temp
	}
	return nil
}

func getValueOfMap(parsedMap map[string]*LineParsed, name string) *int {
	entry, exists := parsedMap[name]
	if !exists {
		return nil
	}
	if entry.value != nil {
		return entry.value
	}
	leftValue := *getValueOfMap(parsedMap, *entry.leftValue)
	rightValue :=
		*getValueOfMap(parsedMap, *entry.rightValue)
	operator := *entry.operator
	return operatorResult(operator, leftValue, rightValue)

}

func part2NewMap(parsedMap map[string]*LineParsed, name string) (*LineParsed, *string) {
	for _, parsed := range parsedMap {
		if parsed.leftValue == nil || parsed.rightValue == nil {
			continue
		}
		if (*parsed.leftValue == name || *parsed.rightValue == name) && parsed.name == "root" {
			return &LineParsed{
				name:  name,
				value: parsed.value,
			}, nil
		}
		if *parsed.leftValue == name {
			operator := *parsed.operator

			if operator == "*" {
				return &LineParsed{
					name:       name,
					operator:   util.GetPtr("/"),
					rightValue: parsed.rightValue,
					leftValue:  util.GetPtr(parsed.name),
				}, util.GetPtr(parsed.name)
			}
			if operator == "/" {
				return &LineParsed{
					name:       name,
					operator:   util.GetPtr("*"),
					rightValue: parsed.rightValue,
					leftValue:  util.GetPtr(parsed.name),
				}, util.GetPtr(parsed.name)
			}
			if operator == "+" {
				return &LineParsed{
					name:       name,
					operator:   util.GetPtr("-"),
					rightValue: parsed.rightValue,
					leftValue:  util.GetPtr(parsed.name),
				}, util.GetPtr(parsed.name)
			}
			if operator == "-" {
				return &LineParsed{
					name:       name,
					operator:   util.GetPtr("+"),
					rightValue: parsed.rightValue,
					leftValue:  util.GetPtr(parsed.name),
				}, util.GetPtr(parsed.name)
			}
		}
		if *parsed.rightValue == name {
			operator := *parsed.operator

			if operator == "*" {
				return &LineParsed{
					name:       name,
					operator:   util.GetPtr("/"),
					rightValue: parsed.leftValue,
					leftValue:  util.GetPtr(parsed.name),
				}, util.GetPtr(parsed.name)
			}
			if operator == "/" {
				return &LineParsed{
					name:       name,
					operator:   util.GetPtr("/"),
					leftValue:  parsed.leftValue,
					rightValue: util.GetPtr(parsed.name),
				}, util.GetPtr(parsed.name)
			}
			if operator == "+" {
				return &LineParsed{
					name:       name,
					operator:   util.GetPtr("-"),
					rightValue: parsed.leftValue,
					leftValue:  util.GetPtr(parsed.name),
				}, util.GetPtr(parsed.name)
			}
			if operator == "-" {
				return &LineParsed{
					name:       name,
					operator:   util.GetPtr("-"),
					leftValue:  parsed.leftValue,
					rightValue: util.GetPtr(parsed.name),
				}, util.GetPtr(parsed.name)
			}
		}
	}
	return nil, nil
}

//func getValueOfMapPart2(parsedMap map[string]*LineParsed, name string) *int {
//	for _, parsed := range parsedMap {
//		if parsed.leftValue == nil || parsed.rightValue == nil {
//			continue
//		}
//		if (*parsed.leftValue == name || *parsed.rightValue == name) && parsed.name == "root" {
//			return parsed.value
//		}
//		if *parsed.leftValue == name {
//			result := *getValueOfMapPart2(parsedMap, parsed.name)
//			right := *getValueOfMap(parsedMap, *parsed.rightValue)
//			// left = X / right
//			// / -> left*right
//			// * -> left/right
//			// + -> left-right
//			// - -> left+right
//			operator := *parsed.operator
//			if operator == "*" {
//				temp := result / right
//				return &temp
//			}
//			if operator == "+" {
//				temp := result - right
//				return &temp
//			}
//			if operator == "-" {
//				temp := result + right
//				return &temp
//			}
//			if operator == "/" {
//				temp := result * right
//				return &temp
//			}
//
//		}
//		if *parsed.rightValue == name {
//			result := *getValueOfMapPart2(parsedMap, parsed.name)
//			left := *getValueOfMap(parsedMap, *parsed.leftValue)
//
//			// result = right / x
//			// / -> result*right
//			// * -> result/right
//			// + -> result-right
//			// - -> result+right
//			operator := *parsed.operator
//
//			if operator == "*" {
//				temp := result / left
//				return &temp
//			}
//			if operator == "+" {
//				temp := result - left
//				return &temp
//			}
//			if operator == "-" {
//				temp := result + left
//				return &temp
//			}
//			if operator == "/" {
//				temp := result * left
//				return &temp
//			}
//		}
//	}
//	return nil
//}

func evalPart1(parsedMap map[string]*LineParsed) {
	final := getValueOfMap(parsedMap, "root")

	log.Println(*final)
}
func evalPart2(parsedMap map[string]*LineParsed) {
	//parsedMap["humn"].value = nil

	//ideal := 72950437237500

	//for i := 0; i < 100000000; i++ {
	//	parsedMap["humn"].value = &i
	//	res := getValueOfMap(parsedMap, "jwcq")
	//
	//	if (i%10000 == 0) {
	//		log.Println(*res, i)
	//	}
	//	if *res == ideal {
	//		log.Println("stop", i)
	//		return
	//	}
	//}

	equalValue := getValueOfMap(parsedMap, "swbn")
	//equalValue := getValueOfMap(parsedMap, "sjmn")
	parsedMap["root"].value = equalValue

	result := make(map[string]*LineParsed)

	util.CloneMap(result, parsedMap)

	var nextKey = util.GetPtr("humn")
	for nextKey != nil {
		newLine, newKey := part2NewMap(parsedMap, *nextKey)
		if newLine != nil {
			result[newLine.name] = newLine
		}
		nextKey = newKey
	}

	log.Println(*equalValue)

	final := getValueOfMap(result, "humn")
	log.Println(*final)
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	parsedMap := make(map[string]*LineParsed)

	for _, result := range results {
		parsed := parseLine(result)
		parsedMap[parsed.name] = parsed
	}

	if part2 {
		evalPart2(parsedMap)
	} else {
		evalPart1(parsedMap)
	}

}
