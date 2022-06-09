package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	b := SudokuBoard{
		board: &board,
	}

	return b.isValid()
}

type SudokuBoard struct {
	board *[][]byte
}

func (b *SudokuBoard) isValid() bool {
	for i := 0; i < 9; i++ {
		if b.hasDuplicates(b.Row(i)) {
			return false
		}
		if b.hasDuplicates(b.Col(i)) {
			return false
		}
		if b.hasDuplicates(b.Square(i)) {
			return false
		}
	}
	return true
}

func (b *SudokuBoard) Row(i int) []byte {
	return (*b.board)[i]
}

func (b *SudokuBoard) Col(i int) []byte {
	col := make([]byte, 9)

	for j := 0; j < 9; j++ {
		col[j] = (*b.board)[j][i]
	}
	return col
}

func (b *SudokuBoard) Square(i int) []byte {
	square := make([]byte, 9)
	row := (i / 3) * 3
	col := (i % 3) * 3

	for j := 0; j < 9; j++ {
		row := row + (j / 3)
		col := col + (j % 3)
		square[j] = (*b.board)[row][col]
	}
	return square
}

func (b *SudokuBoard) hasDuplicates(numbers []byte) bool {
	seen := make(map[byte]struct{}, 9)

	for _, num := range numbers {

		if num == '.' {
			continue
		}

		if _, ok := seen[num]; ok {
			return true
		}

		seen[num] = struct{}{}
	}
	return false
}

func main() {
	b := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(b))
}
