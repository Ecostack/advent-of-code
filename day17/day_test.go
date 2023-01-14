package day17

import (
	"testing"
)

func TestDayPart1Example(t *testing.T) {
	getValue("./input-example.txt", true, false)
}

func TestDayPart2Example(t *testing.T) {
	getValue("./input-example.txt", true, true)
}

func TestDayPart1(t *testing.T) {
	getValue("./input.txt", false, false)
}

func TestDayPart2(t *testing.T) {
	getValue("./input.txt", false, true)
}
