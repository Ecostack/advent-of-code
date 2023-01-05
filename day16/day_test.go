package day16

import (
	"fmt"
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

func TestBla(t *testing.T) {
	byteArray := []int{0, 3, 5, 6}
	str1 := fmt.Sprintf("%v", byteArray)
	fmt.Println("String =", str1)

	fmt.Println(combinations(4, 3))
}
