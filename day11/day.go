package day11

import (
	"AdventOfCode2022/util"
	"log"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items        []*big.Int
	op           func(val *big.Int) *big.Int
	testDivision *big.Int
	testTrue     int
	testFalse    int
	inspected    int
}

func parsingMonkeys(lines []string) []*Monkey {
	monkeys := make([]*Monkey, 0)
	temp := &Monkey{
		items: make([]*big.Int, 0),
	}
	for _, line := range lines {
		if len(line) == 0 {
			monkeys = append(monkeys, temp)
			temp = &Monkey{
				items: make([]*big.Int, 0),
			}
		}
		if strings.Contains(line, "Starting items") {
			split := strings.Split(line, "  Starting items: ")
			items := strings.Split(split[1], ", ")
			for _, item := range items {
				integer, _ := strconv.Atoi(item)
				temp.items = append(temp.items, big.NewInt(int64(integer)))
			}
		}
		if strings.Contains(line, "Test: divisible by") {
			split := strings.Split(line, "  Test: divisible by ")
			integer, _ := strconv.Atoi(split[1])
			temp.testDivision = big.NewInt(int64(integer))
		}
		if strings.HasPrefix(line, "    If true: throw to monkey ") {
			split := strings.Split(line, "    If true: throw to monkey ")
			integer, _ := strconv.Atoi(split[1])
			temp.testTrue = integer
		}
		if strings.HasPrefix(line, "    If false: throw to monkey ") {
			split := strings.Split(line, "    If false: throw to monkey ")
			integer, _ := strconv.Atoi(split[1])
			temp.testFalse = integer
		}
		if strings.Contains(line, "Operation") {
			split := strings.Split(line, "  Operation: new = old ")
			operators := strings.Split(split[1], " ")
			value := big.NewInt(-1)
			isOld := true
			if operators[1] != "old" {
				isOld = false
				integer, _ := strconv.Atoi(operators[1])
				value = big.NewInt(int64(integer))
			}

			temp.op = func(val *big.Int) *big.Int {
				if isOld {
					if operators[0] == "+" {
						return val.Add(val, val)
					}
					if operators[0] == "-" {
						return val.Sub(val, val)
					}
					if operators[0] == "*" {
						return val.Mul(val, val)
					}
					if operators[0] == "/" {
						return val.Div(val, val)
					}
				}
				if operators[0] == "+" {
					return val.Add(val, value)
				}
				if operators[0] == "-" {
					return val.Sub(val, value)
				}
				if operators[0] == "*" {
					return val.Mul(val, value)
				}
				if operators[0] == "/" {
					return val.Div(val, value)
				}
				panic("missing operator")
				return nil
			}
		}
	}
	monkeys = append(monkeys, temp)
	return monkeys
}

func runTheMonkeys(monkeys []*Monkey, part2 bool) {

	var bigIntDiv = big.NewInt(3)
	if part2 {
		total := big.NewInt(1)
		for _, monkey := range monkeys {
			total.Mul(total, monkey.testDivision)
		}
		bigIntDiv = total
	}

	for _, monkey := range monkeys {
		for _, item := range monkey.items {
			monkey.inspected++
			newItem := monkey.op(item)
			if !part2 {
				newItem.Div(newItem, bigIntDiv)
			} else {
				if newItem.Cmp(bigIntDiv) == 1 {
					newItem.Mod(newItem, bigIntDiv)
				}
			}

			mod := big.NewInt(-1)
			mod.Mod(newItem, monkey.testDivision)

			if mod.Int64() == int64(0) {
				monkeys[monkey.testTrue].items = append(monkeys[monkey.testTrue].items, newItem)
			} else {
				monkeys[monkey.testFalse].items = append(monkeys[monkey.testFalse].items, newItem)
			}
		}
		monkey.items = make([]*big.Int, 0)
	}
}

func printInspected(monkeys []*Monkey) {
	inspected := make([]int, 0)
	for _, monkey := range monkeys {
		inspected = append(inspected, monkey.inspected)
	}

	sort.Ints(inspected)
	inspected = util.Reverse(inspected)

	log.Println(inspected)
	log.Println("result", inspected[0]*inspected[1])
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	monkeys := parsingMonkeys(results)

	rounds := 20
	if part2 {
		rounds = 10000
	}
	for i := 0; i < rounds; i++ {
		runTheMonkeys(monkeys, part2)
		if i == 19 {
			printInspected(monkeys)
		}
		//fmt.Println("i", i)
	}
	printInspected(monkeys)
}
