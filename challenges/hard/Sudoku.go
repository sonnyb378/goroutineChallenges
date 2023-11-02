package hard

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 9

var solutionCh = make(chan bool, 1)
var wg = sync.WaitGroup{}

var boardCh = make(chan [][]int)

func Sudoku() {

	// Challenge 6: Concurrent Sudoku Solver

	// Create a Go program that solves Sudoku puzzles concurrently.
	// The program should be able to take Sudoku boards of varying
	// difficulty levels and use goroutines to attempt to solve them.
	// The Sudoku puzzle should be represented as a 9x9 grid, with
	// some cells filled, and the program should find the solution by
	// concurrently trying different possibilities.

	// Input:
	// A Sudoku puzzle board (9x9 grid) with some cells filled.

	// Desired Output:
	// The program should find and print the solved Sudoku board.

	board := GenerateEmptyBoard()
	FillInitialValues(board, 17)

	fmt.Println("Initial Board:")
	PrintBoard(board)
	fmt.Println()

	SolveSudokuCon(board)

}

func SolveSudokuCon(board [][]int) {

	emptyCells := CheckEmptyCells(board)

	numberOfGoroutines := len(emptyCells) / N

	for i := 1; i <= numberOfGoroutines; i++ {
		wg.Add(1)
		go Solve(emptyCells, board)
	}

	go func() {
		wg.Wait()
		close(solutionCh)
		close(boardCh)
	}()

	timeout := time.Second * 5

	select {
	case solution := <-solutionCh:
		if solution {
			fmt.Println("Solution:")
			PrintBoard(<-boardCh)
			fmt.Println()
		} else {
			fmt.Println("Sudoku puzzle unsolvable")
		}

	case <-time.After(timeout):
		fmt.Println("Timeout: No solution found within the specified time.")
		return
	}

}

func Solve(emptyCells [][2]int, board [][]int) {
	defer wg.Done()

	boardCopy := CopyBoard(board)
	index := 0

	for index < len(emptyCells) {
		row, col := emptyCells[index][0], emptyCells[index][1]
		num := boardCopy[row][col] + 1
		found := false

		for num <= N {
			if IsPossibleNumber(boardCopy, row, col, num) {
				boardCopy[row][col] = num
				index++
				found = true
				break
			}
			num++
		}
		if !found {
			boardCopy[row][col] = 0
			index--
			found = false
			if index < 0 {
				break
			}
		}
	}

	if index < 0 {
		solutionCh <- false
		boardCh <- boardCopy
	} else {
		solutionCh <- true
		boardCh <- boardCopy
	}

}

func CopyBoard(board [][]int) [][]int {
	copyBoard := make([][]int, len(board))
	for i, row := range board {
		copyBoard[i] = append([]int{}, row...)
	}

	return copyBoard
}

func CheckEmptyCells(board [][]int) [][2]int {

	var empty [][2]int

	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if board[row][col] == 0 {
				empty = append(empty, [2]int{row, col})
			}
		}
	}

	return empty
}

func GenerateEmptyBoard() [][]int {
	matrix := make([][]int, N)

	for i := range matrix {
		matrix[i] = make([]int, N)
	}

	return matrix
}

func FillInitialValues(board [][]int, fillCount int) {

	for fillCount > 0 {
		row := rand.Intn(N)
		col := rand.Intn(N)
		if board[row][col] == 0 {
			num := RandomNumber(board, row, col)
			board[row][col] = num
			fillCount--
		}
	}
}

func RandomNumber(board [][]int, row, col int) int {
	var validNumbers []int
	for num := 1; num <= N; num++ {
		if IsPossibleNumber(board, row, col, num) {
			validNumbers = append(validNumbers, num)
		}
	}
	if len(validNumbers) == 0 {
		panic("No valid numbers for the cell.")
	}
	return validNumbers[rand.Intn(len(validNumbers))]
}

func IsPossibleNumber(board [][]int, row, col, num int) bool {
	return !usedInBoard(board, row-row%3, col-col%3, num) &&
		!usedInRow(board, row, num) &&
		!usedInCol(board, col, num)
}

func usedInBoard(board [][]int, row, col, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[row+i][col+j] == num {
				return true
			}
		}
	}
	return false
}

func usedInRow(board [][]int, row, num int) bool {
	for i := 0; i < N; i++ {
		if board[row][i] == num {
			return true
		}
	}
	return false
}
func usedInCol(board [][]int, col, num int) bool {
	for i := 0; i < N; i++ {
		if board[i][col] == num {
			return true
		}
	}
	return false
}

func PrintBoard(board [][]int) {
	for row := 0; row < N; row++ {
		if row%3 == 0 && row != 0 {
			fmt.Println("-----------")
		}
		for col := 0; col < N; col++ {
			if col%3 == 0 && col != 0 {
				fmt.Print("|")
			}
			fmt.Print(board[row][col])
		}
		fmt.Println()
	}
}
