package main

import (
	"errors"
	"fmt"
)

const size = 9

type Board [size][size]int

func main() {
	board := &Board{
		{0, 0, 8, 0, 0, 0, 0, 1, 6},
		{5, 0, 0, 0, 9, 2, 0, 0, 8},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{9, 0, 0, 3, 0, 0, 8, 2, 0},
		{0, 2, 0, 0, 0, 0, 0, 7, 0},
		{0, 8, 4, 0, 0, 6, 0, 0, 5},
		{0, 0, 0, 0, 0, 3, 0, 0, 0},
		{4, 0, 0, 9, 6, 0, 0, 0, 2},
		{1, 6, 0, 0, 0, 0, 7, 0, 0},
	}

	fmt.Println("Initial Sudoku puzzle:")
	printBoard(board)

	if err := validateBoard(board); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if solveSudoku(board) {
		fmt.Println("\nSolved Sudoku:")
		printBoard(board)
	} else {
		fmt.Println("\nNo solution exists")
	}
}

func printBoard(board *Board) {
	for i := 0; i < size; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("------+-------+------")
		}
		for j := 0; j < size; j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			if j == 8 {
				fmt.Printf("%d\n", board[i][j])
			} else {
				fmt.Printf("%d ", board[i][j])
			}
		}
	}
}

func validateBoard(board *Board) error {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] != 0 {
				temp := board[i][j]
				board[i][j] = 0
				if !isValid(board, i, j, temp) {
					return errors.New("invalid initial board configuration")
				}
				board[i][j] = temp
			}
		}
	}
	return nil
}

func solveSudoku(board *Board) bool {
	row, col, found := findEmptyCell(board)
	if !found {
		return true // puzzle solved
	}

	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num

			if solveSudoku(board) {
				return true
			}

			board[row][col] = 0 // backtrack
		}
	}

	return false
}

func findEmptyCell(board *Board) (int, int, bool) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func isValid(board *Board, row, col, num int) bool {
	// Check row
	for x := 0; x < size; x++ {
		if board[row][x] == num {
			return false
		}
	}

	// Check column
	for x := 0; x < size; x++ {
		if board[x][col] == num {
			return false
		}
	}

	// Check 3x3 box
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}
