package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	baseDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)
	logFile, err := os.OpenFile(baseDir+"/go-tetris.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatal("opening log file error:", err)
	}
	defer logFile.Close()
	logger.SetOutput(logFile)

	rand.Seed(time.Now().UnixNano())

	err = loadBoards()
	if err != nil {
		logger.Fatal("loading internal boards error:", err)
	}

	err = loadUserBoards()
	if err != nil {
		logger.Fatal("loading user boards error:", err)
	}

	NewView()
	NewMinos()
	NewBoard()
	NewEdit()
	NewEngine()

	engine.Start()

	view.Stop()
}
