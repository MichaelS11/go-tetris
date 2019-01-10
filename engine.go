package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

// NewEngine creates new engine
func NewEngine() {
	engine = &Engine{
		chanStop: make(chan struct{}, 1),
		gameOver: true,
		tickTime: time.Hour,
		ai:       NewAi(),
	}
}

// Run runs the engine
func (engine *Engine) Run() {
	logger.Println("Engine Run start")

	var event *termbox.Event

	engine.timer = time.NewTimer(engine.tickTime)
	engine.timer.Stop()
	engine.aiTimer = time.NewTimer(engine.tickTime)
	engine.aiTimer.Stop()

	engine.ranking = NewRanking()
	board.Clear()
	view.RefreshScreen()

	engine.keyInput = NewKeyInput()
	go engine.keyInput.Run()

loop:
	for {
		select {
		case <-engine.chanStop:
			break loop
		default:
			select {
			case event = <-engine.keyInput.chanKeyInput:
				engine.keyInput.ProcessEvent(event)
				view.RefreshScreen()
			case <-engine.timer.C:
				engine.tick()
			case <-engine.aiTimer.C:
				engine.ai.ProcessQueue()
				engine.aiTimer.Reset(engine.tickTime / aiTickDivider)
			case <-engine.chanStop:
				break loop
			}
		}
	}

	logger.Println("Engine Run end")
}

// Stop stops the engine
func (engine *Engine) Stop() {
	logger.Println("Engine Stop start")

	if !engine.stopped {
		engine.stopped = true
		close(engine.chanStop)
	}
	engine.timer.Stop()
	engine.aiTimer.Stop()

	logger.Println("Engine Stop end")
}

// Pause pauses the engine
func (engine *Engine) Pause() {
	if !engine.timer.Stop() {
		select {
		case <-engine.timer.C:
		default:
		}
	}
	if !engine.aiTimer.Stop() {
		select {
		case <-engine.aiTimer.C:
		default:
		}
	}
	engine.paused = true
}

// UnPause resumes running the engine
func (engine *Engine) UnPause() {
	engine.timer.Reset(engine.tickTime)
	if engine.aiEnabled {
		engine.aiTimer.Reset(engine.tickTime / aiTickDivider)
	}
	engine.paused = false
}

// PreviewBoard sets previewBoard to true
func (engine *Engine) PreviewBoard() {
	engine.previewBoard = true
}

// NewGame resets board and starts a new game
func (engine *Engine) NewGame() {
	logger.Println("Engine NewGame start")

	board.Clear()
	engine.tickTime = 480 * time.Millisecond
	engine.score = 0
	engine.level = 1
	engine.deleteLines = 0

loop:
	for {
		select {
		case <-engine.keyInput.chanKeyInput:
		default:
			break loop
		}
	}

	engine.previewBoard = false
	engine.gameOver = false
	if engine.aiEnabled {
		engine.ai.GetBestQueue()
	}
	engine.UnPause()

	logger.Println("Engine NewGame end")
}

// ResetTimer resets the time for lock delay or tick time
func (engine *Engine) ResetTimer(duration time.Duration) {
	if !engine.timer.Stop() {
		select {
		case <-engine.timer.C:
		default:
		}
	}
	if duration == 0 {
		// duration 0 means tick time
		engine.timer.Reset(engine.tickTime)
	} else {
		// duration is lock delay
		engine.timer.Reset(duration)
	}
}

// AiGetBestQueue calls AI to get best queue
func (engine *Engine) AiGetBestQueue() {
	if !engine.aiEnabled {
		return
	}
	go engine.ai.GetBestQueue()
}

// tick move mino down and refreshes screen
func (engine *Engine) tick() {
	board.MinoMoveDown()
	view.RefreshScreen()
}

// AddDeleteLines adds deleted lines to score
func (engine *Engine) AddDeleteLines(lines int) {
	engine.deleteLines += lines
	if engine.deleteLines > 999999 {
		engine.deleteLines = 999999
	}

	switch lines {
	case 1:
		engine.AddScore(40 * (engine.level + 1))
	case 2:
		engine.AddScore(100 * (engine.level + 1))
	case 3:
		engine.AddScore(300 * (engine.level + 1))
	case 4:
		engine.AddScore(1200 * (engine.level + 1))
	}

	if engine.level < engine.deleteLines/10 {
		engine.LevelUp()
	}
}

// AddScore adds to score
func (engine *Engine) AddScore(add int) {
	engine.score += add
	if engine.score > 9999999 {
		engine.score = 9999999
	}
}

// LevelUp goes up a level
func (engine *Engine) LevelUp() {
	if engine.level >= 30 {
		return
	}

	engine.level++
	switch {
	case engine.level > 29:
		engine.tickTime = 10 * time.Millisecond
	case engine.level > 25:
		engine.tickTime = 20 * time.Millisecond
	case engine.level > 19:
		// 50 to 30
		engine.tickTime = time.Duration(10*(15-engine.level/2)) * time.Millisecond
	case engine.level > 9:
		// 150 to 60
		engine.tickTime = time.Duration(10*(25-engine.level)) * time.Millisecond
	default:
		// 480 to 160
		engine.tickTime = time.Duration(10*(52-4*engine.level)) * time.Millisecond
	}
}

// GameOver pauses engine and sets to game over
func (engine *Engine) GameOver() {
	logger.Println("Engine GameOver start")

	engine.Pause()
	engine.gameOver = true

	view.ShowGameOverAnimation()

loop:
	for {
		select {
		case <-engine.keyInput.chanKeyInput:
		default:
			break loop
		}
	}

	engine.ranking.InsertScore(uint64(engine.score))
	engine.ranking.Save()

	logger.Println("Engine GameOver end")
}

// EnabledAi enables the AI
func (engine *Engine) EnabledAi() {
	engine.aiEnabled = true
	go engine.ai.GetBestQueue()
	engine.aiTimer.Reset(engine.tickTime / aiTickDivider)
}

// DisableAi disables the AI
func (engine *Engine) DisableAi() {
	engine.aiEnabled = false
	if !engine.aiTimer.Stop() {
		select {
		case <-engine.aiTimer.C:
		default:
		}
	}
}

// EnabledEditMode enables edit mode
func (engine *Engine) EnabledEditMode() {
	edit.EnabledEditMode()
	engine.editMode = true
}

// DisableEditMode disables edit mode
func (engine *Engine) DisableEditMode() {
	edit.DisableEditMode()
	engine.editMode = false
}
