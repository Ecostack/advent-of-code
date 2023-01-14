package day17

import (
	"AdventOfCode2022/util"
	"fmt"
	"log"
)

const CHAMBER_WIDTH = 7

//####
//
//.#.
//###
//.#.
//
//..#
//..#
//###
//
//#
//#
//#
//#
//
//##
//##

const DEBUG_PRINT_MATRIX_EVERY_STEP = false

type Matrix [][]Part

type System struct {
	matrix            Matrix
	rocks             []Rock
	jetMoves          []JetMove
	currentRockIndexX int
	currentRockIndexY int
	lastPutIndex      int
	jetMoveIndex      int
	partIndex         int
	rocksPlaced       int
}

func (s *System) crashIntoStone(rock Rock, rockX, rockY int) bool {
	rockWidth := len(rock[0])
	rockHeight := len(rock)

	if rockY < 0 {
		return true
	}
	if rockX < 0 {
		return true
	}
	if rockX+rockWidth > CHAMBER_WIDTH {
		return true
	}

	for y := 0; y < rockHeight; y++ {
		for x := 0; x < rockWidth; x++ {
			newY := rockY + y
			newX := rockX + x

			part := rock[y][x]
			if part == EmptyPart {
				continue
			}
			matrixPart := s.matrix[newY][newX]
			if matrixPart == EmptyPart {
				continue
			}
			if matrixPart == StonePart && part == StonePart {
				return true
			}
		}
	}
	return false
}

func (s *System) putRock() {
	s.currentRockIndexX = 2
	s.currentRockIndexY = s.lastPutIndex + 3
	currentRock := s.rocks[s.partIndex]
	rockWidth := len(currentRock[0])
	rockHeight := len(currentRock)
	hasLanded := false
	if DEBUG_PRINT_MATRIX_EVERY_STEP {
		fmt.Println("Rock begins to fall:")
	}
	if DEBUG_PRINT_MATRIX_EVERY_STEP {
		printMatrix(s)
	}

	//lastJetHappened := false
	for !hasLanded {
		//printMatrix(s)
		jetMoved := false
		jetmove := s.jetMoves[s.jetMoveIndex]
		if jetmove == MoveLeft {
			if !s.crashIntoStone(currentRock, s.currentRockIndexX-1, s.currentRockIndexY) {
				if DEBUG_PRINT_MATRIX_EVERY_STEP {
					fmt.Println("Jet of gas pushes rock left:")
				}
				jetMoved = true
				s.currentRockIndexX--
			}
		}
		if jetmove == MoveRight {
			if !s.crashIntoStone(currentRock, s.currentRockIndexX+1, s.currentRockIndexY) {

				if DEBUG_PRINT_MATRIX_EVERY_STEP {
					fmt.Println("Jet of gas pushes rock right:")
				}
				jetMoved = true
				s.currentRockIndexX++
			}
		}
		if !jetMoved {
			if DEBUG_PRINT_MATRIX_EVERY_STEP {
				fmt.Println("Jet of gas pushes rock but nothing happens:")
			}

		}
		if DEBUG_PRINT_MATRIX_EVERY_STEP {
			printMatrix(s)
		}

		s.jetMoveIndex = (s.jetMoveIndex + 1) % len(s.jetMoves)

		if s.crashIntoStone(currentRock, s.currentRockIndexX, s.currentRockIndexY-1) {

			if DEBUG_PRINT_MATRIX_EVERY_STEP {
				fmt.Println("Rock falls 1 unit, causing it to come to rest:")
			}
			hasLanded = true
		} else {
			if DEBUG_PRINT_MATRIX_EVERY_STEP {
				fmt.Println("Rock falls 1 unit:")
			}

			s.currentRockIndexY--
		}
		if DEBUG_PRINT_MATRIX_EVERY_STEP {
			printMatrix(s)
		}
	}
	if s.currentRockIndexY+rockHeight > s.lastPutIndex {
		s.lastPutIndex = s.currentRockIndexY + rockHeight
	}
	for y := 0; y < rockHeight; y++ {
		for x := 0; x < rockWidth; x++ {
			newY := s.currentRockIndexY + y
			newX := s.currentRockIndexX + x
			part := currentRock[y][x]
			if part == StonePart {
				s.matrix[newY][newX] = part
			}
		}
	}

	s.partIndex = (s.partIndex + 1) % len(s.rocks)

	s.currentRockIndexX = -1
	s.currentRockIndexY = -1
}

type JetMove uint8

const MoveLeft JetMove = 1
const MoveRight JetMove = 2

type Part uint8

const EmptyPart Part = 0
const StonePart Part = 1
const StoneMoving Part = 2

type Rock [][]Part

func getJetMoves(example bool) []JetMove {
	lines, err := util.GetFileContentsSplit("./jet.txt")
	util.PanicOnError(err)
	if example {
		lines, err = util.GetFileContentsSplit("./jet-example.txt")
		util.PanicOnError(err)
	}
	line := lines[0]

	moves := make([]JetMove, 0)
	for _, i2 := range line {
		if i2 == '<' {
			moves = append(moves, MoveLeft)
			continue
		}
		if i2 == '>' {
			moves = append(moves, MoveRight)
		}
	}
	return moves
}

func getRocks(lines []string) []Rock {
	result := make([]Rock, 0)
	rock := make(Rock, 0)
	for _, line := range lines {
		if len(line) == 0 {
			rock = util.Reverse(rock)
			result = append(result, rock)
			rock = make(Rock, 0)
			continue
		}
		rockParts := make([]Part, 0)
		for _, i3 := range line {
			if i3 == '.' {
				rockParts = append(rockParts, EmptyPart)
			} else {
				rockParts = append(rockParts, StonePart)
			}
		}
		rock = append(rock, rockParts)
	}
	rock = util.Reverse(rock)
	result = append(result, rock)
	return result
}

func createMatrix(width uint8) Matrix {
	matrix := make(Matrix, 10000)
	for i, _ := range matrix {
		matrix[i] = make([]Part, width)
	}
	return matrix
}

func printMatrix(s *System) {
	matrix := make(Matrix, len(s.matrix))
	for i, parts := range s.matrix {
		matrix[i] = util.CloneArray(parts)
	}
	currentRock := s.rocks[s.partIndex]
	rockWidth := len(currentRock[0])
	rockHeight := len(currentRock)
	if s.currentRockIndexY >= 0 {
		for y := 0; y < rockHeight; y++ {
			for x := 0; x < rockWidth; x++ {
				newY := s.currentRockIndexY + y
				newX := s.currentRockIndexX + x
				part := currentRock[y][x]
				if part == EmptyPart {
					continue
				}
				matrix[newY][newX] = StoneMoving
			}
		}
	}

	for y := s.lastPutIndex + 8; y >= 0; y-- {
		fmt.Print("|")
		for _, part := range matrix[y] {
			if part == EmptyPart {
				fmt.Print(".")
			}
			if part == StonePart {
				fmt.Print("#")
			}
			if part == StoneMoving {
				fmt.Print("@")
			}
		}
		fmt.Print("|")
		fmt.Print("\n")
	}
	fmt.Print("\n\n")
}

func part1Fn(lines []string, example bool) {
	rocks := getRocks(lines)
	moves := getJetMoves(example)
	matrix := createMatrix(CHAMBER_WIDTH)
	system := &System{matrix: matrix, rocks: rocks, jetMoves: moves, currentRockIndexX: -1, currentRockIndexY: -1}
	for i := 0; i < 2022; i++ {
		system.putRock()
		//log.Println("currentMax:   ", i, "-> ", system.lastPutIndex)
		//fmt.Println(system.lastPutIndex)
		if DEBUG_PRINT_MATRIX_EVERY_STEP {
			printMatrix(system)
		}
	}

	log.Print("part1 ", system.lastPutIndex)
}

func getValue(file string, example bool, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	part1Fn(results, example)
}
