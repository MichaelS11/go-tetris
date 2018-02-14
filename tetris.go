package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/inconshreveable/log15.v2"
)

func main() {

	baseDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	logger = log15.New()
	if baseDir != "" {
		logger.SetHandler(log15.Must.FileHandler(baseDir+"/tetris.log", log15.LogfmtFormat()))
	}

	rand.Seed(time.Now().UnixNano())

	NewMinos()
	NewBoard()
	NewView()
	NewEngine()

	engine.Run()

	view.Stop()

}
