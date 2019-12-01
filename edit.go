package main

import (
	"time"

	"github.com/gdamore/tcell"
)

// NewEdit creates a new edit mode
func NewEdit() {
	edit = &Edit{moved: true}
}

// EnabledEditMode enable edit mode
func (edit *Edit) EnabledEditMode() {
	if edit.y > board.height-1 {
		edit.y = board.height - 1
	}
	if edit.x > board.width-1 {
		edit.x = board.width - 1
	}
	edit.moved = true
}

// DisableEditMode disable edit mode
func (edit *Edit) DisableEditMode() {
	err := saveUserBoards()
	if err != nil {
		logger.Fatal("error saving user boards:", err)
	}
}

// BoardSizeMode changed to board size edit mode
func (edit *Edit) BoardSizeMode() {
	edit.width = board.width
	edit.height = board.height
	edit.boardSize = true
}

// BoardWidthIncrement board width increment
func (edit *Edit) BoardWidthIncrement() {
	if edit.width > 39 {
		return
	}
	edit.width++
}

// BoardWidthDecrement board width decrement
func (edit *Edit) BoardWidthDecrement() {
	if edit.width < 9 {
		return
	}
	edit.width--
}

// BoardHeightIncrement board height increment
func (edit *Edit) BoardHeightIncrement() {
	if edit.height > 39 {
		return
	}
	edit.height++
}

// BoardHeightDecrement board height decrement
func (edit *Edit) BoardHeightDecrement() {
	if edit.height < 9 {
		return
	}
	edit.height--
}

// ChangeBoardSize create new board
func (edit *Edit) ChangeBoardSize() {
	ChangeBoardSize(edit.width, edit.height)
	edit.saved = false
	edit.boardSize = false
}

// EmptyBoard removes all blocks/colors from the board
func (edit *Edit) EmptyBoard() {
	board.EmptyBoard()
}

// CursorUp move cursor up
func (edit *Edit) CursorUp() {
	if !edit.moved {
		edit.moved = true
	}
	if edit.y < 1 {
		return
	}
	edit.y--
}

// CursorDown move cursor down
func (edit *Edit) CursorDown() {
	if !edit.moved {
		edit.moved = true
	}
	if edit.y == board.height-1 {
		return
	}
	edit.y++
}

// CursorRight move cursor right
func (edit *Edit) CursorRight() {
	if !edit.moved {
		edit.moved = true
	}
	if edit.x == board.width-1 {
		return
	}
	edit.x++
}

// CursorLeft move cursor left
func (edit *Edit) CursorLeft() {
	if !edit.moved {
		edit.moved = true
	}
	if edit.x < 1 {
		return
	}
	edit.x--
}

// SetColor sets the board color
func (edit *Edit) SetColor(color tcell.Color) {
	if edit.moved {
		edit.moved = false
	}
	if edit.saved {
		edit.saved = false
	}
	board.SetColor(edit.x, edit.y, color, -1)
}

// RotateLeft rotates cell left
func (edit *Edit) RotateLeft() {
	if edit.moved {
		edit.moved = false
	}
	if edit.saved {
		edit.saved = false
	}
	board.RotateLeft(edit.x, edit.y)
}

// RotateRight rotates cell right
func (edit *Edit) RotateRight() {
	if edit.moved {
		edit.moved = false
	}
	if edit.saved {
		edit.saved = false
	}
	board.RotateRight(edit.x, edit.y)
}

// DrawCursor draws the cursor location when cursor moves
func (edit *Edit) DrawCursor() {
	if !edit.moved {
		return
	}
	board.DrawCursor(edit.x, edit.y)
}

// SaveBoard save board
func (edit *Edit) SaveBoard() {
	if board.boardsIndex < numInternalBoards {
		edit.SaveBoardNew()
		return
	}
	boards[board.boardsIndex].colors = board.colors
	boards[board.boardsIndex].rotation = board.rotation
	if !edit.saved {
		edit.saved = true
	}
}

// SaveBoardNew save board as new board
func (edit *Edit) SaveBoardNew() {
	aBoards := Boards{name: time.Now().Format("Jan 2 3:4:5")}
	aBoards.colors = make([][]tcell.Color, len(board.colors))
	for i := 0; i < len(board.colors); i++ {
		aBoards.colors[i] = append(aBoards.colors[i], board.colors[i]...)
	}
	aBoards.rotation = make([][]int, len(board.rotation))
	for i := 0; i < len(board.rotation); i++ {
		aBoards.rotation[i] = append(aBoards.rotation[i], board.rotation[i]...)
	}
	boards = append(boards, aBoards)
	board.boardsIndex = len(boards) - 1
	if !edit.saved {
		edit.saved = true
	}
}

// DeleteBoard deletes a board
func (edit *Edit) DeleteBoard() {
	if board.boardsIndex < numInternalBoards {
		return
	}
	boards = append(boards[:board.boardsIndex], boards[board.boardsIndex+1:]...)
	board.boardsIndex--
	board.Clear()
}
