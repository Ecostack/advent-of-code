package day12

import (
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
	bla := make([]string, 0)
	bla = append(bla, "11")
	bla2 := append(bla, "12")
	bla3 := append(bla, "13")
	log.Println(bla, bla2, bla3)
}
