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
		log.Fatal("error opening logFile:", err)
	}
	defer logFile.Close()
	logger.SetOutput(logFile)

	rand.Seed(time.Now().UnixNano())

	NewMinos()
	NewBoard()
	NewView()
	NewEngine()

	engine.Run()

	view.Stop()

}
