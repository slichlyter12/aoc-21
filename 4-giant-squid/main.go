package main

import (
	"fmt"

	"github.com/slichlyter12/aoc-21/4-giant-squid/bingo"
	"github.com/slichlyter12/aoc-21/common"
)

func main() {
	input := common.GetFileStrs("4-giant-squid/input.txt")
	calls := bingo.GetCalls(input[0])
	var boards []bingo.Board

	// Turn input into boards
	var boardRows []string
	for rowIndex, row := range input {
		// Skip first two lines, we've dealt with those
		if rowIndex < 2 {
			continue
		}

		// if we've reached a break it means we've seen all the rows for this
		// board
		if row == "" {
			boards = append(boards, bingo.NewBoard(boardRows))
			boardRows = make([]string, 0)
			continue
		}

		boardRows = append(boardRows, row)
	}

	// Play and print winner
	winningBoard, winningCall := bingo.Play(calls, boards)
	score := bingo.Score(winningBoard, winningCall)

	fmt.Println("Winning Board:")
	winningBoard.Print()
	fmt.Println("Winning Call:", winningCall)
	fmt.Println("Winning Score:", score)

	fmt.Println("-----------")

	lastToWinBoard, lastToWinCall := bingo.LastToWin(calls, boards)
	lastToWinScore := bingo.Score(lastToWinBoard, lastToWinCall)

	fmt.Println("Last to win board:")
	lastToWinBoard.Print()
	fmt.Println("Last to win call:", lastToWinCall)
	fmt.Println("Last to win score:", lastToWinScore)
}
