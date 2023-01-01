package day15

import (
	"AdventOfCode2022/util"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"sync"
)

type Part int8

const (
	EmptyPart    Part = 0
	BeaconPart   Part = 1
	SensorPart   Part = 2
	NoBeaconPart Part = 3
)

type Position struct {
	x int
	y int
}

type Sensor struct {
	position  *Position
	beaconPos *Position
	distance  int
}

type Matrix struct {
	matrix   map[int]map[int]Part
	sensors  []*Sensor
	mapMutex sync.RWMutex
}

func getMatrixValueY(matrix *Matrix, y int) map[int]Part {
	matrix.mapMutex.Lock()
	if _, exists := matrix.matrix[y]; !exists {
		matrix.matrix[y] = make(map[int]Part)
	}
	val := matrix.matrix[y]
	matrix.mapMutex.Unlock()
	return val
}

func setMatrixValue(matrix *Matrix, x, y int, part Part) {
	matrix.mapMutex.Lock()
	if _, exists := matrix.matrix[y]; !exists {
		matrix.matrix[y] = make(map[int]Part)
	}
	matrix.matrix[y][x] = part
	matrix.mapMutex.Unlock()
}

func getMatrixValue(matrix *Matrix, x, y int) Part {
	matrix.mapMutex.Lock()
	if _, exists := matrix.matrix[y]; !exists {
		matrix.matrix[y] = make(map[int]Part)
	}
	_, exists := matrix.matrix[y][x]
	if !exists {
		matrix.matrix[y][x] = EmptyPart
	}
	val := matrix.matrix[y][x]
	matrix.mapMutex.Unlock()
	return val
}

func parseLine(matrix *Matrix, line string) {
	//Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	expr, err := regexp.Compile("(-?\\d+)")
	util.PanicOnError(err)
	matches := expr.FindAllString(line, -1)

	sensor := &Sensor{
		position:  nil,
		beaconPos: nil,
	}

	values := make([]int, 0)
	for i, match := range matches {
		temp, err := strconv.Atoi(match)
		util.PanicOnError(err)
		values = append(values, temp)
		if i == 1 {
			setMatrixValue(matrix, values[0], values[1], SensorPart)

			sensor.position = &Position{
				x: values[0],
				y: values[1],
			}
			values = make([]int, 0)
		}
		if i == 3 {
			setMatrixValue(matrix, values[0], values[1], BeaconPart)

			sensor.beaconPos = &Position{
				x: values[0],
				y: values[1],
			}
		}
	}
	distance := int(math.Abs(float64(sensor.position.x-sensor.beaconPos.x)) + math.Abs(float64(sensor.position.y-sensor.beaconPos.y)))
	sensor.distance = distance
	fmt.Println("(assert  (> (+ (abs (- x " + strconv.Itoa(sensor.position.x) + ")) (abs (- y " + strconv.Itoa(sensor.position.y) + ")) ) " + strconv.Itoa(sensor.distance) + "))")

	matrix.sensors = append(matrix.sensors, sensor)
}

func createMatrix() *Matrix {
	//RANGE := 2000000
	matrix := &Matrix{matrix: make(map[int]map[int]Part), mapMutex: sync.RWMutex{}}
	//for y := -RANGE; y < RANGE; y++ {
	//	matrix.matrix[y] = make(map[int]Part)
	//	for x := -RANGE; x < RANGE; x++ {
	//		matrix.matrix[y][x] = EmptyPart
	//	}
	//}
	return matrix
}

func printMatrix(matrix *Matrix) {
	//yKeys := make([]int, 0, len(matrix.matrix))
	//for k := range matrix.matrix {
	//	yKeys = append(yKeys, k)
	//}
	//sort.Sort(sort.IntSlice(yKeys))
	//
	//for _, yKey := range yKeys {
	//	xKeys := make([]int, 0, len(matrix.matrix[yKey]))
	//	for k := range matrix.matrix[yKey] {
	//		xKeys = append(xKeys, k)
	//	}
	//	sort.Sort(sort.IntSlice(xKeys))
	//
	//	for _, xKey := range xKeys {
	//		part := getMatrixValue(matrix, xKey, yKey)
	//		if part == BeaconPart {
	//			fmt.Print("B")
	//		} else if part == SensorPart {
	//			fmt.Print("S")
	//		} else if part == NoBeaconPart {
	//			fmt.Print("#")
	//		} else {
	//			fmt.Print(".")
	//		}
	//	}
	//	fmt.Print("\n")
	//}
	//fmt.Print("\n")

	maximum := 20

	for y := 0; y < maximum; y++ {
		for x := 0; x < maximum; x++ {
			part := getMatrixValue(matrix, x, y)
			if part == BeaconPart {
				fmt.Print("B")
			} else if part == SensorPart {
				fmt.Print("S")
			} else if part == NoBeaconPart {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	//
	//for _, yKey := range yKeys {
	//	xKeys := make([]int, 0, len(matrix.matrix[yKey]))
	//	for k := range matrix.matrix[yKey] {
	//		xKeys = append(xKeys, k)
	//	}
	//	sort.Sort(sort.IntSlice(xKeys))
	//
	//	for _, xKey := range xKeys {
	//		part := getMatrixValue(matrix, xKey, yKey)
	//		if part == BeaconPart {
	//			fmt.Print("B")
	//		} else if part == SensorPart {
	//			fmt.Print("S")
	//		} else if part == NoBeaconPart {
	//			fmt.Print("#")
	//		} else {
	//			fmt.Print(".")
	//		}
	//	}
	//	fmt.Print("\n")
	//}
	//fmt.Print("\n")
}

func part1FnApproachOld(matrix *Matrix, example bool) {

	maximum := 20

	for _, sensor := range matrix.sensors {

		for x := -sensor.distance; x <= sensor.distance; x++ {
			tempY := sensor.distance - int(math.Abs(float64(x)))
			for y := 0; y <= tempY; y++ {
				newX := sensor.position.x + x
				newY := sensor.position.y + y
				if getMatrixValue(matrix, newX, newY) == EmptyPart {
					setMatrixValue(matrix, newX, newY, NoBeaconPart)
				}
				if getMatrixValue(matrix, newX, -newY) == EmptyPart {
					setMatrixValue(matrix, newX, -newY, NoBeaconPart)
				}
			}
		}
	}

	fmt.Println(getMatrixValue(matrix, 14, 12))
	fmt.Println(getMatrixValue(matrix, 14, 10))
	fmt.Println(getMatrixValue(matrix, 13, 11))
	fmt.Println(getMatrixValue(matrix, 14, 11))
	fmt.Println(getMatrixValue(matrix, 15, 11))

	checkPoint := func(x, y int) {
		if x >= 0 && x <= maximum && y >= 0 && y <= maximum {
			if getMatrixValue(matrix, x, y) != EmptyPart {
				return
			}
			if getMatrixValue(matrix, x+1, y) != EmptyPart &&
				getMatrixValue(matrix, x-1, y) != EmptyPart &&
				getMatrixValue(matrix, x, y+1) != EmptyPart &&
				getMatrixValue(matrix, x, y-1) != EmptyPart {
				log.Println("part2", x, y, x*maximum+y)
			}
		}
	}

	for _, sensor := range matrix.sensors {
		distance := sensor.distance
		newDistance := distance + 1
		for x := -newDistance; x <= newDistance; x++ {
			y := newDistance - int(math.Abs(float64(x)))

			newX := sensor.position.x + x
			newY := sensor.position.y + y

			checkPoint(newX, newY)
			checkPoint(newX, -newY)
			//newX := sensor.position.x + x
			//newY := sensor.position.y + y
			//setMatrixValue(matrix, newX, newY, NoBeaconPart)
			//setMatrixValue(matrix, newX, -newY, NoBeaconPart)
		}
	}

	//for y := 0; y < maximum; y++ {
	//	for x := 0; x < maximum; x++ {
	//		for _, sensor := range matrix.sensors {
	//			sensor := sensor
	//			//wg.Add(1)
	//			//task := func() {
	//			//	defer wg.Done()
	//			distance := int(math.Abs(float64(sensor.position.x-sensor.beaconPos.x)) + math.Abs(float64(sensor.position.y-sensor.beaconPos.y)))
	//			//for x := sensor.position.x - distance; x < sensor.position.x+distance; x++ {
	//			val := getMatrixValue(matrix, x, y)
	//			currentDistance := int(math.Abs(float64(sensor.position.x-x)) + math.Abs(float64(sensor.position.y-y)))
	//			if currentDistance <= distance {
	//				if val == EmptyPart {
	//					setMatrixValue(matrix, x, y, NoBeaconPart)
	//				}
	//			}
	//
	//			//}
	//		}
	//	}
	//	fmt.Print("\n")
	//}
	//fmt.Print("\n")

}

func part1FnApproach2(matrix *Matrix, example bool) {
	y := 2000000
	if example {
		y = 10
	}

	count := int64(0)
	for _, sensor := range matrix.sensors {
		sensor := sensor
		//wg.Add(1)
		//task := func() {
		//	defer wg.Done()
		distance := int(math.Abs(float64(sensor.position.x-sensor.beaconPos.x)) + math.Abs(float64(sensor.position.y-sensor.beaconPos.y)))
		for x := sensor.position.x - distance; x < sensor.position.x+distance; x++ {
			val := getMatrixValue(matrix, x, y)
			currentDistance := int(math.Abs(float64(sensor.position.x-x)) + math.Abs(float64(sensor.position.y-y)))
			if currentDistance <= distance {
				if val == EmptyPart {
					count++
					setMatrixValue(matrix, x, y, NoBeaconPart)
				}
			}

		}
	}
	log.Println("part1", count)
}

func part2Fn(matrix *Matrix) {
	maximum := 4000000

	for y := 0; y < maximum; y++ {
		for x := 0; x < maximum; x++ {

			hasFoundOne := false
			for _, sensor := range matrix.sensors {
				if hasFoundOne {
					continue
				}
				distance := int(math.Abs(float64(sensor.position.x-sensor.beaconPos.x)) + math.Abs(float64(sensor.position.y-sensor.beaconPos.y)))
				currentDistance := int(math.Abs(float64(sensor.position.x-x)) + math.Abs(float64(sensor.position.y-y)))
				//log.Println("distance", distance)
				if currentDistance <= distance {
					hasFoundOne = true
					break
				}
			}
			if !hasFoundOne {
				log.Println("part2 x,y, res ", x, y, (x*maximum)+y)
			}
		}
	}

	//for y := 0; y < maximum; y++ {
	//	for x := 0; x < maximum; x++ {
	//		//val := getMatrixValue(matrix, x, y)
	//		//if val == EmptyPart {
	//		//	log.Println("part2 x,y, res ", x, y, (x*maximum)+y)
	//		//}
	//	}
	//}

}

func getValue(file string, example bool, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	matrix := createMatrix()
	for _, result := range results {
		parseLine(matrix, result)
	}

	//determineNonAvailableBeacons(matrix)
	//printMatrix(matrix)

	//if part2 {
	//	part2Fn(matrix)
	//} else {
	part1FnApproachOld(matrix, example)
	//part1FnApproach2(matrix, example)
	printMatrix(matrix)
	//}
	//figurePart1Fn(matrix, example)
}

//; Variable declarations
//(declare-fun x () Int)
//(declare-fun y () Int)
//
//; Constraints
//(assert (>= x 0))
//(assert (<= x 4000000))
//(assert (>= y 0))
//(assert (<= y 4000000))
//
//
//(assert  (> (+ (abs (- x 2208586)) (abs (- y 2744871)) ) 749486))
//(assert  (> (+ (abs (- x 3937279)) (abs (- y 2452476)) ) 479626))
//(assert  (> (+ (abs (- x 3535638)) (abs (- y 3151860)) ) 785113))
//(assert  (> (+ (abs (- x 1867584)) (abs (- y 2125870)) ) 929181))
//(assert  (> (+ (abs (- x 2290971)) (abs (- y 1583182)) ) 1048482))
//(assert  (> (+ (abs (- x 3137806)) (abs (- y 2216828)) ) 312578))
//(assert  (> (+ (abs (- x 3393352)) (abs (- y 331000)) ) 1828796))
//(assert  (> (+ (abs (- x 1444302)) (abs (- y 821689)) ) 1260561))
//(assert  (> (+ (abs (- x 1084667)) (abs (- y 3412239)) ) 1041801))
//(assert  (> (+ (abs (- x 439341)) (abs (- y 3916996)) ) 915627))
//(assert  (> (+ (abs (- x 295460)) (abs (- y 2114590)) ) 1811587))
//(assert  (> (+ (abs (- x 2212046)) (abs (- y 3819484)) ) 556131))
//(assert  (> (+ (abs (- x 3413280)) (abs (- y 3862465)) ) 1174044))
//(assert  (> (+ (abs (- x 3744934)) (abs (- y 1572303)) ) 888692))
//(assert  (> (+ (abs (- x 3349047)) (abs (- y 2522469)) ) 457361))
//(assert  (> (+ (abs (- x 171415)) (abs (- y 591241)) ) 412283))
//(assert  (> (+ (abs (- x 3237499)) (abs (- y 2150414)) ) 154357))
//(assert  (> (+ (abs (- x 559077)) (abs (- y 454593)) ) 280839))
//(assert  (> (+ (abs (- x 3030733)) (abs (- y 2047512)) ) 250335))
//(assert  (> (+ (abs (- x 1667358)) (abs (- y 3956837)) ) 1003708))
//(assert  (> (+ (abs (- x 1850337)) (abs (- y 98963)) ) 464728))
//(assert  (> (+ (abs (- x 2699546)) (abs (- y 3157824)) ) 827493))
//(assert  (> (+ (abs (- x 1113195)) (abs (- y 98130)) ) 868109))
//(assert  (> (+ (abs (- x 59337)) (abs (- y 246804)) ) 426690))
//(assert  (> (+ (abs (- x 566043)) (abs (- y 29068)) ) 544518))
//(assert  (> (+ (abs (- x 2831421)) (abs (- y 2581088)) ) 489874))
//(assert  (> (+ (abs (- x 597818)) (abs (- y 749461)) ) 614448))
//
//; Solve
//(check-sat)
//(get-model)

//sat
//(model
//  (define-fun y () Int
//    2766584)
//  (define-fun x () Int
//    3135800)
//)
