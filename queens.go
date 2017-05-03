package main

import "fmt"

const (
	maxRows      = 8
	maxCols      = 8
	blockValue   = -1
	unblockValue = 1
)

func main() {
	// Create the chess board
	board := make([][]int, maxRows)
	// Allocate the board
	for i := range board {
		board[i] = make([]int, maxCols)
	}

	AddAQueen(board, 0, 0)
	PrintBoard(board)
}

// AddAQueen - Adds the next queen
func AddAQueen(board [][]int, row, qcount int) bool {
	for col := range board[row] {
		if board[row][col] == 0 {
			board[row][col] = 1
			qcount++
			Blocker(board, row, col)
			// Test if we are done
			if qcount == maxRows {
				return true
			}
			// Add another queen
			if AddAQueen(board, row+1, qcount) {
				return true
			}
			// Adding this queen didn't work out.  Need to try next column
			qcount--
			board[row][col] = 0
			Unblocker(board, row, col)
		}
	}
	return false
}

// Blocker -- Blocks squares where there can't be a queen or unblocks
func Blocker(board [][]int, row, col int) {
	Setter(board, row, col, blockValue)
}

// Unblocker -- Unblocks squares where there can't be a queen or unblocks
func Unblocker(board [][]int, row, col int) {
	// Note: Cannot unblock by setting value to 0 because a square may have
	// been blocked more than once (i.e. by 2 different queens)
	Setter(board, row, col, unblockValue)
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
}
