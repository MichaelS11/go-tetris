package main

import (
	"github.com/nsf/termbox-go"
	"runtime"
)

type KeyInput struct {
	stopped      bool
	chanStop     chan struct{}
	chanKeyInput chan *termbox.Event
}

func NewKeyInput() *KeyInput {
	return &KeyInput{
		chanStop:     make(chan struct{}, 1),
		chanKeyInput: make(chan *termbox.Event, 8),
	}
}

func (keyInput *KeyInput) Run() {
	logger.Info("KeyInput Run start")

loop:
	for {
		select {
		case <-keyInput.chanStop:
			break loop
		default:
		}
		event := termbox.PollEvent()
		if event.Type == termbox.EventKey {
			select {
			case <-keyInput.chanStop:
				break loop
			default:
				select {
				case keyInput.chanKeyInput <- &event:
				case <-keyInput.chanStop:
					break loop
				}
			}
		}
	}

	logger.Info("KeyInput Run end")
}

func (keyInput *KeyInput) ProcessEvent(event *termbox.Event) {

	if event.Key == termbox.KeyCtrlI {
		// ctrl i to log stack trace
		buffer := make([]byte, 1<<16)
		length := runtime.Stack(buffer, true)
		logger.Debug("Stack trace", "buffer", string(buffer[:length]))
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
		if event.Key == termbox.KeySpace {
			engine.NewGame()
		}
		return
	}

	if engine.paused {
		if event.Ch == 'p' {
			engine.UnPause()
		}
		return
	}

	if event.Ch != 0 {
		switch event.Ch {
		case 'p':
			engine.Pause()
		case 'i':
			engine.EnabledAi()
		}
	}

	if engine.aiEnabled {
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
		}
	}

}
