package day3

import (
	"log"
	"testing"
)

func TestDayPart1(t *testing.T) {
	getValue("./input.txt", false)
}

func TestDayPart2(t *testing.T) {
	getValue("./input.txt", true)
}

func TestSplit(t *testing.T) {
	bla := []int{4, 5, 6, 7}
	log.Println(bla[:2])
	log.Println(bla[2:])
}
