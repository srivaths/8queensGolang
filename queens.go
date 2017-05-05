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

var solutionCount int
func main() {
	// Create the chess board
	board := make([][]int, maxRows)
	// Allocate the board
	for i := range board {
		board[i] = make([]int, maxCols)
	}

	queenForRow(board, 0, 0)
	fmt.Printf("%d solutions.\n", solutionCount)
}

// queenForRow - Adds the next queen
func queenForRow(board [][]int, row, qcount int) {
	for col, value := range board[row] {
		if value == empty {
			qcount = addQueen(board, row, col, qcount)
			// Test if we are done
			if qcount == maxRows {
				printBoard(board)
			} else {
				queenForRow(board, row+1, qcount)
			}
			// Prepare to try next column
			qcount = removeQueen(qcount, board, row, col)
		}
	}
}

// removeQueen -- Performs steps involved with removing a queen off the board
func removeQueen(qcount int, board [][]int, row int, col int) (int) {
	qcount--
	board[row][col] = empty
	unblocker(board, row, col)
	return qcount
}

// addQueen -- Performs steps involved with adding a queen to the board
func addQueen(board [][]int, row, col, qcount int) (int) {
	board[row][col] = queen
	qcount++
	blocker(board, row, col)
	return qcount
}

// blocker -- Blocks squares where there can't be a queen or unblocks
func blocker(board [][]int, row, col int) {
	setter(board, row, col, invalidateIncrement)
}

// unblocker -- Unblocks squares where there can't be a queen or unblocks
func unblocker(board [][]int, row, col int) {
	// Note: Cannot unblock by setting value to 0 because a square may have
	// been blocked more than once (i.e. by 2 different queens)
	setter(board, row, col, validateIncrement)
}

// setter - Sets target board squares to specified value
func setter(board [][]int, row, col, valueToSet int) {
	// Set east
	for x := col + 1; x < maxCols; x++ {
		board[row][x] += valueToSet
	}
	// Set south
	for x := row + 1; x < maxRows; x++ {
		board[x][col] += valueToSet
	}
	// Set se
	for x, y := row + 1, col + 1 ; x < maxRows && y < maxCols; x, y = x + 1, y + 1 {
		board[x][y] += valueToSet
	}
	// Set sw
	for x, y := row + 1, col - 1; x < maxRows && y > -1; x, y = x+1, y-1 {
		board[x][y] += valueToSet
	}
}

// printBoard -- Prints the board
func printBoard(board [][]int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] != queen {
				fmt.Print("- ")
			} else {
				fmt.Print("Q ")
			}
		}
		fmt.Println()
	}
	solutionCount++
	fmt.Println("********************")
}
