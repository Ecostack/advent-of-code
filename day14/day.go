package day14

import (
	"AdventOfCode2022/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Part int

const (
	Empty Part = 0
	Stone Part = 1
	Sand  Part = 2
)

type Matrix struct {
	matrix       [][]Part
	lowestStoneY int
}

type Position struct {
	row int
	col int
}

func parseLine(line string) []*Position {
	split := strings.Split(line, " -> ")
	positions := make([]*Position, 0)
	for _, s := range split {
		numbers := strings.Split(s, ",")
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		pos := &Position{
			row: y,
			col: x,
		}
		positions = append(positions, pos)
	}
	return positions
}

func fillMatrixInternal(matrix *Matrix, pos1, pos2 *Position) {
	colFrom := pos1.col
	colTo := pos2.col

	rowFrom := pos1.row
	rowTo := pos2.row
	x := colFrom

	for {
		y := rowFrom
		for {
			if matrix.lowestStoneY < y {
				matrix.lowestStoneY = y
			}
			matrix.matrix[y][x] = Stone
			if y == rowTo {
				break
			}
			if rowFrom > rowTo {
				y--
			} else {
				y++
			}
		}
		if x == colTo {
			break
		}
		if colFrom > colTo {
			x--
		} else {
			x++
		}
	}
}

func fillMatrix(matrix *Matrix, positions []*Position) {
	for i := 0; i < len(positions)-1; i++ {
		pos1 := positions[i]
		pos2 := positions[i+1]
		fillMatrixInternal(matrix, pos1, pos2)
	}
}

func printMatrix(matrix *Matrix) {
	rowLen := len(matrix.matrix)
	colLen := len(matrix.matrix[0])
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if matrix.matrix[i][j] == Stone {
				fmt.Print("#")
			} else if matrix.matrix[i][j] == Sand {
				fmt.Print("+")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func createMatrix() *Matrix {
	matrix := &Matrix{matrix: make([][]Part, 0), lowestStoneY: -1}
	for i := 0; i < 200; i++ {
		parts := make([]Part, 0)
		for j := 0; j < 700; j++ {
			parts = append(parts, Empty)
		}
		matrix.matrix = append(matrix.matrix, parts)
	}
	return matrix
}

func moveCornUntilRest(matrix *Matrix) bool {
	startX := 500
	startY := 0
	for {
		if startY > matrix.lowestStoneY {
			return true
		}
		if matrix.matrix[startY][startX] == Sand {
			return true
		}
		if matrix.matrix[startY+1][startX] == Empty {
			startY++
			continue
		}
		if matrix.matrix[startY+1][startX-1] == Empty {
			startY++
			startX--
			continue
		}
		if matrix.matrix[startY+1][startX+1] == Empty {
			startY++
			startX++
			continue
		}
		matrix.matrix[startY][startX] = Sand
		break
	}
	return false
}

func runSandCorns(matrix *Matrix) int {
	cornIndex := 0
	for {
		cornIndex++
		res := moveCornUntilRest(matrix)
		if res {
			break
		}
	}
	return cornIndex - 1
}

func fillFloor(matrix *Matrix) {
	for x := 0; x < len(matrix.matrix[0]); x++ {
		matrix.matrix[matrix.lowestStoneY+2][x] = Stone
	}
	matrix.lowestStoneY += 2
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	matrix := createMatrix()
	for _, line := range results {
		positions := parseLine(line)
		fillMatrix(matrix, positions)
	}
	if part2 {
		fillFloor(matrix)
	}

	printMatrix(matrix)
	res := runSandCorns(matrix)
	printMatrix(matrix)
	log.Println("result", res)
}
