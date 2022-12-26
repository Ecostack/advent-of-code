package day5

import (
	"AdventOfCode2022/util"
	"log"
	"math"
)

func dec2snafuPat(dec int) string {
	if dec == 2 {
		return "2"
	}
	if dec == 1 {
		return "1"
	}
	if dec == 0 {
		return "0"
	}
	if dec == -1 {
		return "-"
	}

	if dec == -2 {
		return "="
	}
	panic("something is wrong")
	return ""
}

func parseSNAFUPart(part rune) int {
	if part == '1' {
		return 1
	}
	if part == '2' {
		return 2
	}
	if part == '-' {
		return -1
	}
	if part == '=' {
		return -2
	}
	return 0
}

func parseSNAFU(snafu string) int {
	temp := make([]int, len(snafu))
	for i, char := range snafu {
		temp[i] = parseSNAFUPart(char)
	}
	reversed := util.Reverse(temp)
	result := 0
	for i := len(reversed) - 1; i >= 0; i-- {
		val := reversed[i] * int(math.Pow(5, float64(i)))
		result += val
	}
	return result
}

func dec2Snafu(val int) string {
	outerIndex := 0
	multiplierForMaximumValue := 0
	found := false
	for !found {
		temp := int(math.Pow(5, float64(outerIndex)))
		if temp >= val {
			multiplierForMaximumValue = 1
			found = true
			continue
		}
		if temp*2 >= val {
			multiplierForMaximumValue = 2
			found = true
			continue
		}
		outerIndex++
	}

	result := dec2snafuPat(multiplierForMaximumValue)
	currentTemp := int(math.Pow(5, float64(outerIndex))) * multiplierForMaximumValue
	for i := outerIndex - 1; i >= 0; i-- {
		indexValue := int(math.Pow(5, float64(i)))
		lowestMultiplier := 9999999999999999
		lowestDifference := 9999999999999999
		for j := -2; j <= 2; j++ {
			multiValue := indexValue * j
			combined := currentTemp + multiValue
			if int(math.Abs(float64(val-combined))) < lowestDifference {
				lowestMultiplier = j
				lowestDifference = int(math.Abs(float64(val - combined)))
			}
		}
		currentTemp += indexValue * lowestMultiplier
		result += dec2snafuPat(lowestMultiplier)
	}

	log.Println("result", result, "currentTemp", currentTemp)

	return result
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	sum := 0
	for _, result := range results {
		sum += parseSNAFU(result)
	}

	log.Println(sum, dec2Snafu(sum))
}
