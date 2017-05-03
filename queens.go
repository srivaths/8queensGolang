package main

import "fmt"

const (
	maxRows             = 8
	maxCols             = 8
	invalidateIncrement = -1 // Indicates that this location cannot host a queen
	validateIncrement   = 1
	queen = 1
	empty = 0
)

func main() {
	// Create the chess board
	board := make([][]int, maxRows)
	// Allocate the board
	for i := range board {
		board[i] = make([]int, maxCols)
	}

	QueenForRow(board, 0, 0)
}

// QueenForRow - Adds the next queen
func QueenForRow(board [][]int, row, qcount int) {
	for col := range board[row] {
		if board[row][col] == empty {
			qcount = AddQueen(board, row, col, qcount)
			// Test if we are done
			if qcount == maxRows {
				PrintBoard(board)
			} else {
				QueenForRow(board, row+1, qcount)
			}
			// Prepare to try next column
			qcount = RemoveQueen(qcount, board, row, col)
		}
	}
}

// RemoveQueen -- Performs steps involved with removing a queen off the board
func RemoveQueen(qcount int, board [][]int, row int, col int) (int) {
	qcount--
	board[row][col] = empty
	Unblocker(board, row, col)
	return qcount
}

// AddQueen -- Performs steps involved with adding a queen to the board
func AddQueen(board [][]int, row, col, qcount int) (int) {
	board[row][col] = queen
	qcount++
	Blocker(board, row, col)
	return qcount
}

// Blocker -- Blocks squares where there can't be a queen or unblocks
func Blocker(board [][]int, row, col int) {
	Setter(board, row, col, invalidateIncrement)
}

// Unblocker -- Unblocks squares where there can't be a queen or unblocks
func Unblocker(board [][]int, row, col int) {
	// Note: Cannot unblock by setting value to 0 because a square may have
	// been blocked more than once (i.e. by 2 different queens)
	Setter(board, row, col, validateIncrement)
}

// Setter - Sets target board squares to specified value
func Setter(board [][]int, row, col, valueToSet int) {
	// block east
	for x := col + 1; x < maxCols; x++ {
		board[row][x] += valueToSet
	}
	// block south
	for x := row + 1; x < maxRows; x++ {
		board[x][col] += valueToSet
	}
	// block se
	var y = col + 1
	for x := row + 1; x < maxRows; x++ {
		if y == maxCols {
			break
		}
		board[x][y] += valueToSet
		y++
	}
	// block sw
	y = col - 1
	for x := row + 1; x < maxRows; x++ {
		if y == -1 {
			break
		}
		board[x][y] += valueToSet
		y--
	}
}

// PrintBoard -- Prints the board
func PrintBoard(board [][]int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] != 1 {
				fmt.Print("_ ")
			} else {
				fmt.Print("Q ")
			}
		}
		fmt.Println()
	}
	fmt.Println("********************")
}
