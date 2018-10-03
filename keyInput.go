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
	if event.Key == termbox.KeyCtrlI {
		// Ctrl l (lower case L) to log stack trace
		buffer := make([]byte, 1<<16)
		length := runtime.Stack(buffer, true)
		logger.Println("Stack trace")
		logger.Println(string(buffer[:length]))
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

	if event.Ch == 0 {
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
	} else {
		switch event.Ch {
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

}
