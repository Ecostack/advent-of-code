package day1

import "testing"

func TestDay1Part1(t *testing.T) {
	result, _ := getValue("./input.txt")
	if result != 70369 {
		t.Fail()
	}
}

func TestDay1Part2(t *testing.T) {
	_, topThree := getValue("./input.txt")
	if topThree != 31216 {
		t.Fail()
	}
}
