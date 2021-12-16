package bingo

import (
	"fmt"
	"strings"

	"github.com/slichlyter12/aoc-21/common"
)

type Board struct {
	Values [][]int
	Seen   [][]bool
}

func NewBoard(boardInput []string) Board {
	var boardRows [][]int
	var seenRows [][]bool
	for _, row := range boardInput {
		var boardRow []int
		var seenRow []bool
		values := strings.Fields(row)
		for _, value := range values {
			boardRow = append(boardRow, common.Atoi(value))
			seenRow = append(seenRow, false)
		}

		boardRows = append(boardRows, boardRow)
		seenRows = append(seenRows, seenRow)
	}

	board := Board{
		boardRows,
		seenRows,
	}

	return board
}

func GetCalls(calls string) []int {
	var intCalls []int
	strCalls := strings.Split(calls, ",")
	for _, call := range strCalls {
		intCalls = append(intCalls, common.Atoi(call))
	}

	return intCalls
}

func (board Board) hasWon() bool {
	return board.hasWinningRow() || board.hasWinningColumn()
}

func (board Board) hasWinningRow() bool {
	for _, row := range board.Seen {
		for columnIndex, seen := range row {
			if !seen {
				break
			}

			// if we reach the last column and have seen all these values,
			// we've won
			if columnIndex == len(row)-1 {
				return true
			}
		}
	}

	return false
}

func (board Board) hasWinningColumn() bool {
	for columnIndex := 0; columnIndex < len(board.Seen[0]); columnIndex++ {
		for rowIndex := 0; rowIndex < len(board.Seen); rowIndex++ {
			if !board.Seen[rowIndex][columnIndex] {
				break
			}

			// if we reach the last row and have seen all these values,
			// we've won
			if rowIndex == len(board.Seen)-1 {
				return true
			}
		}
	}

	return false
}

func (board *Board) markBoard(call int) {
	for rowIndex, row := range board.Values {
		for columnIndex, value := range row {
			if value == call {
				board.Seen[rowIndex][columnIndex] = true
			}
		}
	}
}

func Play(calls []int, boards []Board) (Board, int) {
	for _, call := range calls {
		for boardIndex := 0; boardIndex < len(boards); boardIndex++ {
			boards[boardIndex].markBoard(call)
			if boards[boardIndex].hasWon() {
				return boards[boardIndex], call
			}
		}
	}

	// No one won :(
	return Board{}, -1
}

func LastToWin(calls []int, boards []Board) (Board, int) {
	for _, call := range calls {
		for boardIndex := 0; boardIndex < len(boards); boardIndex++ {
			boards[boardIndex].markBoard(call)
		}

		var notWon []Board
		for boardIndex := 0; boardIndex < len(boards); boardIndex++ {
			if !boards[boardIndex].hasWon() {
				notWon = append(notWon, boards[boardIndex])
			}
		}

		if len(notWon) == 1 {
			return Play(calls, []Board{notWon[0]})
		}
	}

	return Board{}, -1
}

func Score(board Board, winningCall int) int {
	sum := 0
	for rowIndex := 0; rowIndex < len(board.Values); rowIndex++ {
		for columnIndex := 0; columnIndex < len(board.Values[rowIndex]); columnIndex++ {
			if !board.Seen[rowIndex][columnIndex] {
				sum += board.Values[rowIndex][columnIndex]
			}
		}
	}

	return sum * winningCall
}

func (board Board) Print() {
	for _, row := range board.Values {
		fmt.Println(row)
	}
}
