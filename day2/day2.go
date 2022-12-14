package day2

import (
	"AdventOfCode2022/util"
	"log"
)

type GameHand uint

const (
	HandUnknown GameHand = iota
	Rock
	Paper
	Scissor
)

type GameResult uint

const (
	ResultUnknown GameResult = iota
	Lost
	Draw
	Won
)

type Game struct {
	enemy  GameHand
	player GameHand
	result GameResult
	points uint
}

func (game *Game) calcSecondHand() {
	if game.result == Draw {
		game.player = game.enemy
		return
	}
	if game.result == Won {
		if game.enemy == Rock {
			game.player = Paper
		}
		if game.enemy == Scissor {
			game.player = Rock
		}
		if game.enemy == Paper {
			game.player = Scissor
		}
	}

	if game.result == Lost {
		if game.enemy == Rock {
			game.player = Scissor
		}
		if game.enemy == Scissor {
			game.player = Paper
		}
		if game.enemy == Paper {
			game.player = Rock
		}
	}
}

func (game *Game) calcResult() {
	if game.enemy == game.player {
		game.result = Draw
		return
	}
	if game.enemy == Rock && game.player == Paper {
		game.result = Won
	}
	if game.enemy == Scissor && game.player == Paper {
		game.result = Lost
	}

	if game.enemy == Paper && game.player == Rock {
		game.result = Lost
	}
	if game.enemy == Scissor && game.player == Rock {
		game.result = Won
	}

	if game.enemy == Paper && game.player == Scissor {
		game.result = Won
	}
	if game.enemy == Rock && game.player == Scissor {
		game.result = Lost
	}
}

func (game *Game) calcPoints() {
	pointsHand := uint(0)
	pointsResult := uint(0)
	if game.player == Rock {
		pointsHand = 1
	}
	if game.player == Paper {
		pointsHand = 2
	}
	if game.player == Scissor {
		pointsHand = 3
	}
	if game.result == Draw {
		pointsResult = 3
	}
	if game.result == Won {
		pointsResult = 6
	}
	game.points = pointsResult + pointsHand
}

func parseHand(val rune) GameHand {
	if val == 'A' || val == 'X' {
		return Rock
	}
	if val == 'B' || val == 'Y' {
		return Paper
	}
	if val == 'C' || val == 'Z' {
		return Scissor
	}
	return HandUnknown
}

func parseResult(val rune) GameResult {
	if val == 'X' {
		return Lost
	}
	if val == 'Y' {
		return Draw
	}
	if val == 'Z' {
		return Won
	}
	return ResultUnknown
}

func parseLineToGame(line string, secondColumnResult bool) Game {
	result := Game{
		enemy:  HandUnknown,
		player: HandUnknown,
	}

	runes := []rune(line)
	result.enemy = parseHand(runes[0])

	if secondColumnResult {
		result.result = parseResult(runes[2])
	} else {
		result.player = parseHand(runes[2])
	}

	if secondColumnResult {
		result.calcSecondHand()
	} else {
		result.calcResult()
	}

	result.calcPoints()

	return result
}

func getValue(file string, secondColumnResult bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)
	games := make([]Game, 0)
	totalPoints := uint(0)
	for _, result := range results {
		if len(result) == 0 {
			continue
		}
		games = append(games, parseLineToGame(result, secondColumnResult))
		totalPoints += games[len(games)-1].points
	}
	log.Println("game points: ", totalPoints)
}
