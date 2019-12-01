package main

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/gdamore/tcell"
)

type testMinoStruct struct {
	minoRotation MinoRotation
	x            int
	y            int
}

func TestMain(m *testing.M) {
	setupForTesting()
	retCode := m.Run()
	os.Exit(retCode)
}

func setupForTesting() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)

	rand.Seed(1)

	err := loadBoards()
	if err != nil {
		log.Fatal("error loading boards:", err)
	}

	screen, err = tcell.NewScreen()
	if err != nil {
		logger.Fatal("NewScreen error:", err)
	}

	NewMinos()
	NewBoard()
	NewEngine()
}
