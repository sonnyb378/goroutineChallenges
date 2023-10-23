package hard

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func MatrixMultiplication(rows int, cols int) [][]int {

	// Challenge 1: Parallel Matrix Multiplication
	// Write a Go program that multiplies two large matrices concurrently using goroutines.
	// The program should efficiently distribute the work among multiple goroutines to speed
	// up the computation.

	// Input:
	// Two matrices to be multiplied.

	// Desired Output:
	// The resulting product matrix.

	wg := sync.WaitGroup{}
	ch := make(chan []int, cols)

	matrixC := GenerateEmptyMatrix(rows, cols) // resulting matrix

	matrixA := GenerateMatrix(rows, cols)
	fmt.Println("Matrix A: ", matrixA)
	matrixB := GenerateMatrix(rows, cols)
	fmt.Println("Matrix B: ", matrixB)

	timeStart := time.Now()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			wg.Add(1)
			go ComputeCell(&wg, ch, matrixA[i][j], matrixB[i][j], i, j)
		}
	}
	fmt.Printf("Elapsed time: %v microseconds\n", time.Since(timeStart).Microseconds())

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Plot producy matrix
	for c := range ch {
		matrixC[c[0]][c[1]] = c[2]
	}

	return matrixC
}

func ComputeCell(wg *sync.WaitGroup, ch chan []int, cellA, cellB, i, j int) {
	defer wg.Done()
	cellResult := []int{i, j, MultiplyCells(cellA, cellB)}
	// cellResult = []int{ cellA index, cellB index, product of cellA and cellB}
	ch <- cellResult
}

func MultiplyCells(cellA int, cellB int) int {
	return cellA * cellB
}

func GenerateMatrix(rows int, cols int) [][]int {

	matrix := GenerateEmptyMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = GenerateCellNumber()
		}
	}

	return matrix
}

func GenerateEmptyMatrix(rows int, cols int) [][]int {
	// creates empty matrix
	// 3 x 3 = [[] [] []]
	matrix := make([][]int, rows)

	// creates matrix with 0 value
	// sample: 3 x 3 = [[0 0 0] [0 0 0] [0 0 0]]
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	return matrix
}

func GenerateCellNumber() int {
	// generates random number between 1 - 10
	randomNumber := rand.Intn(10) + 1
	return randomNumber
}
