package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToBingoGame(t *testing.T) {
	randomNumbers := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	game := ReadFileToBingoGame("prompt-input")

	assert.Equal(t, randomNumbers, game.randomNumbers)
	assert.Len(t, game.boards, 3)
	assert.Equal(t, BingoSquares{{22, false}, {13, false}, {17, false}, {11, false}, {0, false}}, game.boards[0].rows[0])
	assert.Equal(t, BingoSquares{{15, false}, {18, false}, {8, false}, {11, false}, {21, false}}, game.boards[1].columns[1])
}

func TestPlayOneTurn(t *testing.T) {
	game := ReadFileToBingoGame("prompt-input")

	PlayOneTurnOneBoard(2, game.boards[0])

	assert.Equal(t, true, game.boards[0].rows[1][1].marked)
	assert.Equal(t, true, game.boards[0].columns[1][1].marked)

	PlayOneTurnOneBoard(5, game.boards[0])

	assert.Equal(t, true, game.boards[0].rows[3][4].marked)
	assert.Equal(t, true, game.boards[0].columns[4][3].marked)
}

func TestIsBoardWinner(t *testing.T) {
	game := ReadFileToBingoGame("prompt-input")

	board := game.boards[0]

	assert.False(t, IsBoardWinner(board))

	PlayOneTurnOneBoard(17, board)
	assert.False(t, IsBoardWinner(board))

	PlayOneTurnOneBoard(23, board)
	assert.False(t, IsBoardWinner(board))

	PlayOneTurnOneBoard(14, board)
	assert.False(t, IsBoardWinner(board))
	PlayOneTurnOneBoard(3, board)
	assert.False(t, IsBoardWinner(board))

	PlayOneTurnOneBoard(20, board)

	assert.True(t, IsBoardWinner(board))
}

func TestFindFirstWinningBoard(t *testing.T) {
	game := ReadFileToBingoGame("prompt-input")

	winner := FindFirstWinningBoard(game)

	assert.NotNil(t, winner)

	assert.Equal(t, 24, winner.rows[0][3].number)
	assert.True(t, winner.rows[0][3].marked)
	assert.Equal(t, 24, winner.columns[3][0].number)
	assert.True(t, winner.columns[3][0].marked)

	assert.Equal(t, game.boards[2], winner)
}

func TestSolvePromptInput(t *testing.T) {
	game := ReadFileToBingoGame("prompt-input")
	assert.Equal(t, 4512, Solve(game))
}

func TestSolve(t *testing.T) {
	game := ReadFileToBingoGame("input")
	assert.Equal(t, 54275, Solve(game))
}
