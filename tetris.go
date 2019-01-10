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
		logger.Fatal("error opening log file:", err)
	}
	defer logFile.Close()
	logger.SetOutput(logFile)

	rand.Seed(time.Now().UnixNano())

	err = loadBoards()
	if err != nil {
		logger.Fatal("error loading internal boards:", err)
	}

	err = loadUserBoards()
	if err != nil {
		logger.Fatal("error loading user boards:", err)
	}

	NewMinos()
	NewBoard()
	NewView()
	NewEngine()
	NewEdit()

	engine.Run()

	view.Stop()
}
