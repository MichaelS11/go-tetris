package main

import (
	"runtime"

	"github.com/nsf/termbox-go"
)

// NewKeyInput creates a new KeyInput
func NewKeyInput() *KeyInput {
	return &KeyInput{
		chanStop:     make(chan struct{}, 1),
		chanKeyInput: make(chan *termbox.Event, 8),
	}
}

// Run starts the KeyInput engine
func (keyInput *KeyInput) Run() {
	logger.Println("KeyInput Run start")

loop:
	for {
		select {
		case <-keyInput.chanStop:
			break loop
		default:
		}
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey && len(keyInput.chanKeyInput) < 8 {
			select {
			case <-keyInput.chanStop:
				break loop
			case keyInput.chanKeyInput <- &event:
			}
		}
	}

	logger.Println("KeyInput Run end")
}

// ProcessEvent process the key input event
func (keyInput *KeyInput) ProcessEvent(event *termbox.Event) {
	if event.Key == termbox.KeyCtrlL {
		// Ctrl l (lower case L) to log stack trace
		buffer := make([]byte, 1<<16)
		length := runtime.Stack(buffer, true)
		logger.Println("Stack trace")
		logger.Println(string(buffer[:length]))
		return
	}

	if engine.editMode {
		if edit.boardSize {
			switch event.Ch {
			case 0:
				switch event.Key {
				case termbox.KeyArrowUp:
					edit.BoardHeightIncrement()
				case termbox.KeyArrowDown:
					edit.BoardHeightDecrement()
				case termbox.KeyArrowLeft:
					edit.BoardWidthDecrement()
				case termbox.KeyArrowRight:
					edit.BoardWidthIncrement()
				}
			case 'q':
				edit.ChangeBoardSize()
			}
		} else {
			switch event.Ch {
			case 0:
				switch event.Key {
				case termbox.KeyArrowUp:
					edit.CursorUp()
				case termbox.KeyArrowDown:
					edit.CursorDown()
				case termbox.KeyArrowLeft:
					edit.CursorLeft()
				case termbox.KeyArrowRight:
					edit.CursorRight()
				case termbox.KeyCtrlB:
					edit.BoardSizeMode()
				case termbox.KeyCtrlS:
					edit.SaveBoard()
				case termbox.KeyCtrlN:
					edit.SaveBoardNew()
				case termbox.KeyCtrlK:
					edit.DeleteBoard()
				case termbox.KeyCtrlO:
					edit.EmptyBoard()
				case termbox.KeyCtrlQ, termbox.KeyCtrlC:
					engine.DisableEditMode()
				}
			case 'c':
				edit.SetColor(termbox.ColorCyan)
			case 'b':
				edit.SetColor(termbox.ColorBlue)
			case 'w':
				edit.SetColor(termbox.ColorWhite)
			case 'e':
				edit.SetColor(termbox.ColorYellow)
			case 'g':
				edit.SetColor(termbox.ColorGreen)
			case 'a':
				edit.SetColor(termbox.ColorMagenta)
			case 'r':
				edit.SetColor(termbox.ColorRed)
			case 'f':
				edit.SetColor(blankColor)
			case 'z':
				edit.RotateLeft()
			case 'x':
				edit.RotateRight()
			}
		}
		return
	}

	if event.Ch == 'q' || event.Key == termbox.KeyCtrlC {
		if !keyInput.stopped {
			keyInput.stopped = true
			close(keyInput.chanStop)
		}
		engine.Stop()
		return
	}

	if engine.gameOver {
		if event.Ch == 0 {
			switch event.Key {
			case termbox.KeySpace:
				engine.NewGame()
			case termbox.KeyArrowLeft:
				board.PreviousBoard()
			case termbox.KeyArrowRight:
				board.NextBoard()
			case termbox.KeyCtrlE:
				engine.EnabledEditMode()
			}
		}
		return
	}

	if engine.paused {
		if event.Ch == 'p' {
			engine.UnPause()
		}
		return
	}

	if engine.aiEnabled {
		switch event.Ch {
		case 'p':
			engine.Pause()
		case 'i':
			engine.DisableAi()
		}
		return
	}

	switch event.Ch {
	case 0:
		switch event.Key {
		case termbox.KeySpace:
			board.MinoDrop()
		case termbox.KeyArrowUp:
			board.MinoDrop()
		case termbox.KeyArrowDown:
			board.MinoMoveDown()
		case termbox.KeyArrowLeft:
			board.MinoMoveLeft()
		case termbox.KeyArrowRight:
			board.MinoMoveRight()
		}
	case 'z':
		board.MinoRotateLeft()
	case 'x':
		board.MinoRotateRight()
	case 'p':
		engine.Pause()
	case 'i':
		engine.EnabledAi()
	}

}
