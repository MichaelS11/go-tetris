package main

import (
	"github.com/nsf/termbox-go"
	"gopkg.in/inconshreveable/log15.v2"
	"os"
	"path/filepath"
)

const (
	boardWidth  = 10
	boardHeight = 20
	blankColor  = termbox.ColorBlack
)

var (
	baseDir string
	logger  log15.Logger
	view    *View
	engine  *Engine
	board   *Board
)

func main() {

	baseDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	logger = log15.New()
	if baseDir != "" {
		logger.SetHandler(log15.Must.FileHandler(baseDir+"/tetris.log", log15.LogfmtFormat()))
	}

	view = NewView()
	engine = NewEngine()

	engine.Run()

	view.Stop()

}
