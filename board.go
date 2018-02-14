package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func NewBoard() {
	board = &Board{}
	board.Clear()
}

func (board *Board) Clear() {
	board.width = len(boards[board.boardsIndex].colors)
	board.height = len(boards[board.boardsIndex].colors[0])
	board.colors = make([][]termbox.Attribute, len(boards[board.boardsIndex].colors))
	for i := 0; i < len(boards[board.boardsIndex].colors); i++ {
		board.colors[i] = make([]termbox.Attribute, len(boards[board.boardsIndex].colors[i]))
		copy(board.colors[i], boards[board.boardsIndex].colors[i])
	}
	board.rotation = make([][]int, len(boards[board.boardsIndex].rotation))
	for i := 0; i < len(boards[board.boardsIndex].rotation); i++ {
		board.rotation[i] = make([]int, len(boards[board.boardsIndex].rotation[i]))
		copy(board.rotation[i], boards[board.boardsIndex].rotation[i])
	}
	board.previewMino = NewMino()
	board.currentMino = NewMino()
}

func (board *Board) PreviousBoard() {
	board.boardsIndex--
	if board.boardsIndex < 0 {
		board.boardsIndex = len(boards) - 1
	}
	engine.PreviewBoard()
	board.Clear()
}

func (board *Board) NextBoard() {
	board.boardsIndex++
	if board.boardsIndex == len(boards) {
		board.boardsIndex = 0
	}
	engine.PreviewBoard()
	board.Clear()
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
	engine.AiGetBestQueue()
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
	for j := 0; j < board.height; j++ {
		if board.isFullLine(j) {
			fullLines = append(fullLines, j)
		}
	}
	return fullLines
}

func (board *Board) isFullLine(j int) bool {
	for i := 0; i < board.width; i++ {
		if board.colors[i][j] == blankColor {
			return false
		}
	}
	return true
}

func (board *Board) deleteLine(line int) {
	for i := 0; i < board.width; i++ {
		board.colors[i][line] = blankColor
	}
	for j := line; j > 0; j-- {
		for i := 0; i < board.width; i++ {
			board.colors[i][j] = board.colors[i][j-1]
			board.rotation[i][j] = board.rotation[i][j-1]
		}
	}
	for i := 0; i < board.width; i++ {
		board.colors[i][0] = blankColor
	}
}

func (board *Board) SetColor(x int, y int, color termbox.Attribute, rotation int) {
	board.colors[x][y] = color
	board.rotation[x][y] = rotation
}

func ValidBlockLocation(x int, y int, mustBeOnBoard bool) bool {
	if x < 0 || x >= board.width || y >= board.height {
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
	return x >= 0 && x < board.width && y >= 0 && y < board.height
}

func (board *Board) DrawBoard() {
	for i := 0; i < board.width; i++ {
		for j := 0; j < board.height; j++ {
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
func (board *Board) printDebugBoard() {
	for j := 0; j < board.height; j++ {
		for i := 0; i < board.width; i++ {
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
			default:
				fmt.Print("U")
			}
		}
		fmt.Println("")
	}
}

// for debuging
func (board *Board) getDebugBoard() []string {
	lines := make([]string, board.height)
	for j := 0; j < board.height; j++ {
		for i := 0; i < board.width; i++ {
			switch board.colors[i][j] {
			case blankColor:
				lines[j] += "."
			case termbox.ColorBlue:
				lines[j] += "B"
			case termbox.ColorCyan:
				lines[j] += "C"
			case termbox.ColorGreen:
				lines[j] += "G"
			case termbox.ColorMagenta:
				lines[j] += "M"
			case termbox.ColorRed:
				lines[j] += "R"
			case termbox.ColorWhite:
				lines[j] += "W"
			case termbox.ColorYellow:
				lines[j] += "Y"
			default:
				lines[j] += "U"
			}
		}
	}
	return lines
}

// for debuging
func (board *Board) getDebugBoardWithMino(mino *Mino) []string {
	lines := make([]string, board.height)
	for j := 0; j < board.height; j++ {
		for i := 0; i < board.width; i++ {
			switch mino.getMinoColorAtLocation(i, j) {
			case blankColor:
				switch board.colors[i][j] {
				case blankColor:
					lines[j] += "."
				case termbox.ColorBlue:
					lines[j] += "B"
				case termbox.ColorCyan:
					lines[j] += "C"
				case termbox.ColorGreen:
					lines[j] += "G"
				case termbox.ColorMagenta:
					lines[j] += "M"
				case termbox.ColorRed:
					lines[j] += "R"
				case termbox.ColorWhite:
					lines[j] += "W"
				case termbox.ColorYellow:
					lines[j] += "Y"
				default:
					lines[j] += "U"
				}
			case termbox.ColorBlue:
				lines[j] += "b"
			case termbox.ColorCyan:
				lines[j] += "c"
			case termbox.ColorGreen:
				lines[j] += "g"
			case termbox.ColorMagenta:
				lines[j] += "m"
			case termbox.ColorRed:
				lines[j] += "r"
			case termbox.ColorWhite:
				lines[j] += "w"
			case termbox.ColorYellow:
				lines[j] += "y"
			default:
				lines[j] += "u"
			}
		}
	}
	return lines
}
