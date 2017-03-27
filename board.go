package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

type Board struct {
	colors       [boardWidth][boardHeight]termbox.Attribute
	rotation     [boardWidth][boardHeight]int
	previewMino  *Mino
	currentMino  *Mino
	dropDistance int
}

func NewBoard() *Board {
	board := &Board{}
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			board.colors[i][j] = blankColor
		}
	}
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			board.rotation[i][j] = 0
		}
	}
	board.previewMino = NewMino()
	board.currentMino = NewMino()
	return board
}

func (board *Board) MinoMoveLeft() {
	board.dropDistance = 0
	mino := board.currentMino.CloneMoveLeft()
	if mino.ValidLocation(false) {
		board.currentMino = mino
		board.StartLockDelayIfBottom()
	}
}

func (board *Board) MinoMoveRight() {
	board.dropDistance = 0
	mino := board.currentMino.CloneMoveRight()
	if mino.ValidLocation(false) {
		board.currentMino = mino
		board.StartLockDelayIfBottom()
	}
}

func (board *Board) MinoRotateRight() {
	board.dropDistance = 0
	mino := board.currentMino.CloneRotateRight()
	if mino.ValidLocation(false) {
		board.currentMino = mino
		board.StartLockDelayIfBottom()
		return
	}
	mino.MoveLeft()
	if mino.ValidLocation(false) {
		board.currentMino = mino
		board.StartLockDelayIfBottom()
		return
	}
	mino.MoveRight()
	mino.MoveRight()
	if mino.ValidLocation(false) {
		board.currentMino = mino
		board.StartLockDelayIfBottom()
		return
	}
}

func (board *Board) MinoRotateLeft() {
	board.dropDistance = 0
	mino := board.currentMino.CloneRotateLeft()
	if mino.ValidLocation(false) {
		board.currentMino = mino
		board.StartLockDelayIfBottom()
		return
	}
	mino.MoveLeft()
	if mino.ValidLocation(false) {
		board.currentMino = mino
		board.StartLockDelayIfBottom()
		return
	}
	mino.MoveRight()
	mino.MoveRight()
	if mino.ValidLocation(false) {
		board.currentMino = mino
		board.StartLockDelayIfBottom()
		return
	}
}

func (board *Board) MinoMoveDown() {
	mino := board.currentMino.CloneMoveDown()
	if mino.ValidLocation(false) {
		board.dropDistance = 0
		board.currentMino = mino
		if !board.StartLockDelayIfBottom() {
			engine.ResetTimer(0)
		}
		return
	}
	if !board.currentMino.ValidLocation(true) {
		engine.GameOver()
		return
	}
	board.nextMino()
}

func (board *Board) MinoDrop() {
	board.dropDistance = 0
	mino := board.currentMino.CloneMoveDown()
	for mino.ValidLocation(false) {
		board.dropDistance++
		mino.MoveDown()
	}
	for i := 0; i < board.dropDistance; i++ {
		board.currentMino.MoveDown()
	}
	if !board.currentMino.ValidLocation(true) {
		engine.GameOver()
		return
	}
	if board.dropDistance < 1 {
		return
	}
	if !board.StartLockDelayIfBottom() {
		engine.ResetTimer(0)
	}
}

func (board *Board) StartLockDelayIfBottom() bool {
	mino := board.currentMino.CloneMoveDown()
	if mino.ValidLocation(false) {
		return false
	}
	engine.ResetTimer(300 * time.Millisecond)
	return true
}

func (board *Board) nextMino() {
	engine.AddScore(board.dropDistance)

	board.currentMino.SetOnBoard()

	board.deleteCheck()

	if !board.previewMino.ValidLocation(false) {
		board.previewMino.MoveUp()
		if !board.previewMino.ValidLocation(false) {
			engine.GameOver()
			return
		}
	}

	board.currentMino = board.previewMino
	board.previewMino = NewMino()
	engine.ResetAiTimer()
	engine.ResetTimer(0)
}

func (board *Board) deleteCheck() {
	lines := board.fullLines()
	if len(lines) < 1 {
		return
	}

	view.ShowDeleteAnimation(lines)
	for _, line := range lines {
		board.deleteLine(line)
	}

	engine.AddDeleteLines(len(lines))
}

func (board *Board) fullLines() []int {
	fullLines := make([]int, 0, 1)
	for j := 0; j < boardHeight; j++ {
		if board.isFullLine(j) {
			fullLines = append(fullLines, j)
		}
	}
	return fullLines
}

func (board *Board) isFullLine(j int) bool {
	for i := 0; i < boardWidth; i++ {
		if board.colors[i][j] == blankColor {
			return false
		}
	}
	return true
}

func (board *Board) deleteLine(line int) {
	for i := 0; i < boardWidth; i++ {
		board.colors[i][line] = blankColor
	}
	for j := line; j > 0; j-- {
		for i := 0; i < boardWidth; i++ {
			board.colors[i][j] = board.colors[i][j-1]
			board.rotation[i][j] = board.rotation[i][j-1]
		}
	}
	for i := 0; i < boardWidth; i++ {
		board.colors[i][0] = blankColor
	}
}

func (board *Board) SetColor(x int, y int, color termbox.Attribute, rotation int) {
	board.colors[x][y] = color
	board.rotation[x][y] = rotation
}

func ValidBlockLocation(x int, y int, mustBeOnBoard bool) bool {
	if x < 0 || x >= boardWidth || y >= boardHeight {
		return false
	}
	if mustBeOnBoard {
		if y < 0 {
			return false
		}
	} else {
		if y < -2 {
			return false
		}
	}
	if y > -1 {
		if board.colors[x][y] != blankColor {
			return false
		}
	}
	return true
}

func ValidDisplayLocation(x int, y int) bool {
	return x >= 0 && x < boardWidth && y >= 0 && y < boardHeight
}

func (board *Board) DrawBoard() {
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight; j++ {
			if board.colors[i][j] != blankColor {
				view.DrawBlock(i, j, board.colors[i][j], board.rotation[i][j])
			}
		}
	}
}

func (board *Board) DrawPreviewMino() {
	board.previewMino.DrawMino(MinoPreview)
}

func (board *Board) DrawCurrentMino() {
	board.currentMino.DrawMino(MinoCurrent)
}

func (board *Board) DrawDropMino() {
	mino := board.currentMino.CloneMoveDown()
	if !mino.ValidLocation(false) {
		return
	}
	for mino.ValidLocation(false) {
		mino.MoveDown()
	}
	mino.MoveUp()
	mino.DrawMino(MinoDrop)
}

// for debuging
func (board *Board) drawDebugBoard() {
	for j := 0; j < boardHeight; j++ {
		for i := 0; i < boardWidth; i++ {
			switch board.colors[i][j] {
			case blankColor:
				fmt.Print(".")
			case termbox.ColorBlue:
				fmt.Print("B")
			case termbox.ColorCyan:
				fmt.Print("C")
			case termbox.ColorGreen:
				fmt.Print("G")
			case termbox.ColorMagenta:
				fmt.Print("M")
			case termbox.ColorRed:
				fmt.Print("R")
			case termbox.ColorWhite:
				fmt.Print("W")
			case termbox.ColorYellow:
				fmt.Print("Y")
			}
		}
		fmt.Println("")
	}
}
