package day13

import (
	"encoding/json"
	"log"
	"testing"
)

func TestDayPart1Example(t *testing.T) {
	getValue("./input-example.txt", false)
}

func TestDayPart2Example(t *testing.T) {
	getValue("./input-example.txt", true)
}

func TestDayPart1(t *testing.T) {
	getValue("./input.txt", false)
}

func TestDayPart2(t *testing.T) {
	getValue("./input.txt", true)
}

func TestCopy(t *testing.T) {
	val := "[[4,4],4,4]"
	arrays := make([]any, 0)
	json.Unmarshal([]byte(val), &arrays)
	log.Println(arrays)
}
