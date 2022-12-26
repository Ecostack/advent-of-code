package day5

import (
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
