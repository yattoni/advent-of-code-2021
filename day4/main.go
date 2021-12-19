package day4

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type BingoGame struct {
	randomNumbers []int
	boards        []*BingoBoard
	winningNumber int
}

type BingoBoard struct {
	// store numbers how we need to query them to make life easier
	rows    []BingoSquares
	columns []BingoSquares
}

type BingoSquare struct {
	number int
	marked bool
}

// https://stackoverflow.com/questions/54302129/how-to-print-the-slice-of-pointers-to-get-the-values-instead-of-their-address-wi
type BingoSquares []*BingoSquare

func (squares BingoSquares) String() string {
	s := "["
	for i, user := range squares {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", user)
	}
	return s + "]"
}

func ReadFileToBingoGame(fileName string) *BingoGame {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	fileFields := strings.Fields(string(fileBytes))

	randomNumberStrs := strings.Split(fileFields[0], ",")
	randomNumbers := []int{}
	for _, n := range randomNumberStrs {
		num, _ := strconv.Atoi(n)
		randomNumbers = append(randomNumbers, num)
	}

	boards := []*BingoBoard{}
	for i := 1; i < len(fileFields); i += 25 {
		board := &BingoBoard{}
		board.rows = []BingoSquares{}
		for j := i; j < i+25; j += 5 {
			n, _ := strconv.Atoi(fileFields[j])
			n2, _ := strconv.Atoi(fileFields[j+1])
			n3, _ := strconv.Atoi(fileFields[j+2])
			n4, _ := strconv.Atoi(fileFields[j+3])
			n5, _ := strconv.Atoi(fileFields[j+4])
			board.rows = append(board.rows, []*BingoSquare{{n, false}, {n2, false}, {n3, false}, {n4, false}, {n5, false}})
		}
		board.columns = []BingoSquares{}
		for j := i; j < i+5; j++ {
			n, _ := strconv.Atoi(fileFields[j])
			n2, _ := strconv.Atoi(fileFields[j+5])
			n3, _ := strconv.Atoi(fileFields[j+10])
			n4, _ := strconv.Atoi(fileFields[j+15])
			n5, _ := strconv.Atoi(fileFields[j+20])
			board.columns = append(board.columns, []*BingoSquare{{n, false}, {n2, false}, {n3, false}, {n4, false}, {n5, false}})
		}
		boards = append(boards, board)
	}

	return &BingoGame{randomNumbers, boards, -1}
}

func PlayOneTurnOneBoard(number int, board *BingoBoard) {
	for _, row := range board.rows {
		for _, square := range row {
			if square.number == number {
				square.marked = true
			}
		}
	}
	for _, column := range board.columns {
		for _, square := range column {
			if square.number == number {
				square.marked = true
			}
		}
	}
}

func IsBoardWinner(board *BingoBoard) bool {
	for _, row := range board.rows {
		if allSquaresMarked(row) {
			fmt.Printf("Row %v is a winner", row)
			return true
		}
	}
	for _, column := range board.columns {
		if allSquaresMarked(column) {
			fmt.Printf("Column %v is a winner", column)
			return true
		}
	}
	return false
}

func FindFirstWinningBoard(game *BingoGame) *BingoBoard {
	for _, number := range game.randomNumbers {
		for j, board := range game.boards {
			fmt.Printf("Playing number %d on board %d\n", number, j)
			PlayOneTurnOneBoard(number, board)
			if IsBoardWinner(board) {
				game.winningNumber = number
				return board
			}
		}
	}
	return nil
}

func allSquaresMarked(squares []*BingoSquare) bool {
	for i := 0; i < len(squares); i++ {
		if !squares[i].marked {
			return false
		}
	}
	return true
}

func Solve(game *BingoGame) int {
	winner := FindFirstWinningBoard(game)

	sumUnmarkedSquares := 0

	for _, row := range winner.rows {
		for _, s := range row {
			if !s.marked {
				sumUnmarkedSquares += s.number
			}
		}
	}

	return sumUnmarkedSquares * game.winningNumber
}
