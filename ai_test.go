package main

import (
	"testing"
)

type testAiStruct struct {
	info      string
	minos     []testMinoStruct
	fullLines int
	holes     int
	bumpy     int
}

func TestAI(t *testing.T) {
	// this must be set to the blank 10x20 board
	board.boardsIndex = 0
	board.Clear()

	tests := []testAiStruct{
		{info: "fullLines 2x minoI", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 18}, // minoI
		}, fullLines: 0, holes: 0, bumpy: 1},
		{info: "fullLines 2x2 minoI", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 0, y: 17}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 17}, // minoI
		}, fullLines: 0, holes: 0, bumpy: 2},
		{info: "fullLines 2x minoI minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 18}, // minoI
			{minoRotation: minos.minoBag[3], x: 8, y: 18}, // minoO
		}, fullLines: 1, holes: 0, bumpy: 1},
		{info: "fullLines 2x2 minoI minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 0, y: 17}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 17}, // minoI
			{minoRotation: minos.minoBag[3], x: 8, y: 18}, // minoO
		}, fullLines: 2, holes: 0, bumpy: 0},
		{info: "fullLines 5x minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[3], x: 0, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 2, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 4, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 6, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 8, y: 18}, // minoO
		}, fullLines: 2, holes: 0, bumpy: 0},
		{info: "fullLines 4x4 minoI 2x minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 0, y: 17}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 17}, // minoI
			{minoRotation: minos.minoBag[0], x: 0, y: 16}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 16}, // minoI
			{minoRotation: minos.minoBag[0], x: 0, y: 15}, // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 15}, // minoI
			{minoRotation: minos.minoBag[3], x: 8, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 8, y: 16}, // minoO
		}, fullLines: 4, holes: 0, bumpy: 0},
		{info: "holes 2x minoI minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 6, y: 18}, // minoI
			{minoRotation: minos.minoBag[3], x: 4, y: 17}, // minoO
		}, fullLines: 0, holes: 2, bumpy: 4},
		{info: "holes 6x minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[3], x: 0, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 4, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 8, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 2, y: 16}, // minoO
			{minoRotation: minos.minoBag[3], x: 6, y: 16}, // minoO
		}, fullLines: 0, holes: 8, bumpy: 8},
		{info: "holes 4x minoT 2x minoI", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[5], x: 0, y: 18}, // minoT
			{minoRotation: minos.minoBag[5], x: 7, y: 18}, // minoT
			{minoRotation: minos.minoBag[5], x: 3, y: 16}, // minoT
			{minoRotation: minos.minoBag[5], x: 7, y: 16}, // minoT
			{minoRotation: minos.minoBag[0], x: 2, y: 14}, // minoI
			{minoRotation: minos.minoBag[0], x: 6, y: 14}, // minoI
		}, fullLines: 0, holes: 19, bumpy: 4},
		{info: "holes 3x minoZ", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[6], x: 0, y: 18}, // minoZ
			{minoRotation: minos.minoBag[6], x: 3, y: 18}, // minoZ
			{minoRotation: minos.minoBag[6], x: 6, y: 18}, // minoZ
		}, fullLines: 0, holes: 3, bumpy: 6},
		{info: "holes 4x minoT 2x minoI 2x minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[5], x: 0, y: 18}, // minoT
			{minoRotation: minos.minoBag[5], x: 7, y: 18}, // minoT
			{minoRotation: minos.minoBag[3], x: 0, y: 16}, // minoO
			{minoRotation: minos.minoBag[5], x: 3, y: 16}, // minoT
			{minoRotation: minos.minoBag[5], x: 7, y: 16}, // minoT
			{minoRotation: minos.minoBag[3], x: 0, y: 14}, // minoO
			{minoRotation: minos.minoBag[0], x: 2, y: 14}, // minoI
			{minoRotation: minos.minoBag[0], x: 6, y: 14}, // minoI
		}, fullLines: 1, holes: 9, bumpy: 16},
		{info: "bumpy 2x minoT - 1", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[5], x: 0, y: 18}, // minoT
			{minoRotation: minos.minoBag[5], x: 5, y: 18}, // minoT
		}, fullLines: 0, holes: 0, bumpy: 7},
		{info: "bumpy 2x minoT - 2", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[5], x: 1, y: 18}, // minoT
			{minoRotation: minos.minoBag[5], x: 6, y: 18}, // minoT
		}, fullLines: 0, holes: 0, bumpy: 8},
		{info: "bumpy 2x minoT - 3", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[5], x: 2, y: 18}, // minoT
			{minoRotation: minos.minoBag[5], x: 7, y: 18}, // minoT
		}, fullLines: 0, holes: 0, bumpy: 7},
		{info: "bumpy 2x minoJ - 1", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[1], x: 0, y: 18}, // minoJ
			{minoRotation: minos.minoBag[1], x: 5, y: 18}, // minoJ
		}, fullLines: 0, holes: 0, bumpy: 6},
		{info: "bumpy 2x minoJ - 2", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[1], x: 1, y: 18}, // minoJ
			{minoRotation: minos.minoBag[1], x: 6, y: 18}, // minoJ
		}, fullLines: 0, holes: 0, bumpy: 8},
		{info: "bumpy 2x minoJ - 2", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[1], x: 2, y: 18}, // minoJ
			{minoRotation: minos.minoBag[1], x: 7, y: 18}, // minoJ
		}, fullLines: 0, holes: 0, bumpy: 7},
		{info: "bumpy 2x minoL - 1", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[2], x: 0, y: 18}, // minoL
			{minoRotation: minos.minoBag[2], x: 5, y: 18}, // minoL
		}, fullLines: 0, holes: 0, bumpy: 7},
		{info: "bumpy 2x minoL - 2", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[2], x: 1, y: 18}, // minoL
			{minoRotation: minos.minoBag[2], x: 6, y: 18}, // minoL
		}, fullLines: 0, holes: 0, bumpy: 8},
		{info: "bumpy 2x minoL - 3", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[2], x: 2, y: 18}, // minoL
			{minoRotation: minos.minoBag[2], x: 7, y: 18}, // minoL
		}, fullLines: 0, holes: 0, bumpy: 6},
	}

	runAiTests(t, tests)
}

func TestBigBoardAI(t *testing.T) {
	// this must be set to the blank 20x20 board
	board.boardsIndex = 3
	board.Clear()

	tests := []testAiStruct{
		{info: "fullLines 4x minoI", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 8, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 12, y: 18}, // minoI
		}, fullLines: 0, holes: 0, bumpy: 1},
		{info: "fullLines 5x minoI", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 8, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 12, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 16, y: 18}, // minoI
		}, fullLines: 1, holes: 0, bumpy: 0},
		{info: "fullLines 5x2 minoI", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[0], x: 0, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 8, y: 18},  // minoI
			{minoRotation: minos.minoBag[0], x: 12, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 16, y: 18}, // minoI
			{minoRotation: minos.minoBag[0], x: 0, y: 17},  // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 17},  // minoI
			{minoRotation: minos.minoBag[0], x: 8, y: 17},  // minoI
			{minoRotation: minos.minoBag[0], x: 12, y: 17}, // minoI
			{minoRotation: minos.minoBag[0], x: 16, y: 17}, // minoI
		}, fullLines: 2, holes: 0, bumpy: 0},
		{info: "fullLines 9x minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[3], x: 0, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 2, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 4, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 6, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 8, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 10, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 12, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 14, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 16, y: 18}, // minoO
		}, fullLines: 0, holes: 0, bumpy: 2},
		{info: "fullLines 10x minoO", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[3], x: 0, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 2, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 4, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 6, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 8, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 10, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 12, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 14, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 16, y: 18}, // minoO
			{minoRotation: minos.minoBag[3], x: 18, y: 18}, // minoO
		}, fullLines: 2, holes: 0, bumpy: 0},
		{info: "holes 3x minoO 3x minoI", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[3], x: 0, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 6, y: 18},  // minoO
			{minoRotation: minos.minoBag[3], x: 12, y: 18}, // minoO
			{minoRotation: minos.minoBag[0], x: 2, y: 16},  // minoI
			{minoRotation: minos.minoBag[0], x: 8, y: 16},  // minoI
			{minoRotation: minos.minoBag[0], x: 14, y: 16}, // minoI
		}, fullLines: 0, holes: 24, bumpy: 8},
		{info: "holes 5x minoZ", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[6], x: 0, y: 18},  // minoZ
			{minoRotation: minos.minoBag[6], x: 4, y: 18},  // minoZ
			{minoRotation: minos.minoBag[6], x: 8, y: 18},  // minoZ
			{minoRotation: minos.minoBag[6], x: 12, y: 18}, // minoZ
			{minoRotation: minos.minoBag[6], x: 16, y: 18}, // minoZ
		}, fullLines: 0, holes: 5, bumpy: 18},
		{info: "holes 6x minoT 2x minoO 5x minoI", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[5], x: 0, y: 18},  // minoT
			{minoRotation: minos.minoBag[5], x: 6, y: 18},  // minoT
			{minoRotation: minos.minoBag[5], x: 12, y: 18}, // minoT
			{minoRotation: minos.minoBag[3], x: 18, y: 18}, // minoO
			{minoRotation: minos.minoBag[5], x: 3, y: 16},  // minoT
			{minoRotation: minos.minoBag[5], x: 9, y: 16},  // minoT
			{minoRotation: minos.minoBag[5], x: 15, y: 16}, // minoT
			{minoRotation: minos.minoBag[3], x: 18, y: 16}, // minoO
			{minoRotation: minos.minoBag[0], x: 0, y: 14},  // minoI
			{minoRotation: minos.minoBag[0], x: 4, y: 14},  // minoI
			{minoRotation: minos.minoBag[0], x: 8, y: 14},  // minoI
			{minoRotation: minos.minoBag[0], x: 12, y: 14}, // minoI
			{minoRotation: minos.minoBag[0], x: 16, y: 14}, // minoI
		}, fullLines: 1, holes: 18, bumpy: 23},
		{info: "bumpy 4x minoJ - 1", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[1], x: 0, y: 18},  // minoJ
			{minoRotation: minos.minoBag[1], x: 5, y: 18},  // minoJ
			{minoRotation: minos.minoBag[1], x: 10, y: 18}, // minoJ
			{minoRotation: minos.minoBag[1], x: 15, y: 18}, // minoJ
		}, fullLines: 0, holes: 0, bumpy: 14},
		{info: "bumpy 4x minoJ - 2", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[1], x: 1, y: 18},  // minoJ
			{minoRotation: minos.minoBag[1], x: 6, y: 18},  // minoJ
			{minoRotation: minos.minoBag[1], x: 11, y: 18}, // minoJ
			{minoRotation: minos.minoBag[1], x: 16, y: 18}, // minoJ
		}, fullLines: 0, holes: 0, bumpy: 16},
		{info: "bumpy 4x minoJ - 3", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[1], x: 2, y: 18},  // minoJ
			{minoRotation: minos.minoBag[1], x: 7, y: 18},  // minoJ
			{minoRotation: minos.minoBag[1], x: 12, y: 18}, // minoJ
			{minoRotation: minos.minoBag[1], x: 17, y: 18}, // minoJ
		}, fullLines: 0, holes: 0, bumpy: 15},
		{info: "bumpy 4x minoL - 1", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[2], x: 0, y: 18},  // minoL
			{minoRotation: minos.minoBag[2], x: 5, y: 18},  // minoL
			{minoRotation: minos.minoBag[2], x: 10, y: 18}, // minoL
			{minoRotation: minos.minoBag[2], x: 15, y: 18}, // minoL
		}, fullLines: 0, holes: 0, bumpy: 15},
		{info: "bumpy 4x minoL - 2", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[2], x: 1, y: 18},  // minoL
			{minoRotation: minos.minoBag[2], x: 6, y: 18},  // minoL
			{minoRotation: minos.minoBag[2], x: 11, y: 18}, // minoL
			{minoRotation: minos.minoBag[2], x: 16, y: 18}, // minoL
		}, fullLines: 0, holes: 0, bumpy: 16},
		{info: "bumpy 4x minoL - 3", minos: []testMinoStruct{
			{minoRotation: minos.minoBag[2], x: 2, y: 18},  // minoL
			{minoRotation: minos.minoBag[2], x: 7, y: 18},  // minoL
			{minoRotation: minos.minoBag[2], x: 12, y: 18}, // minoL
			{minoRotation: minos.minoBag[2], x: 17, y: 18}, // minoL
		}, fullLines: 0, holes: 0, bumpy: 14},
	}

	runAiTests(t, tests)
}

func runAiTests(t *testing.T, tests []testAiStruct) {
	var mino1 *Mino
	var mino2 *Mino

	for _, test := range tests {
		board.Clear()

		for i, minoTest := range test.minos {
			mino := NewMino()
			mino.minoRotation = minoTest.minoRotation
			mino.length = len(mino.minoRotation[0])
			mino.x = minoTest.x
			mino.y = minoTest.y
			if i < len(test.minos)-2 {
				mino.SetOnBoard()
			} else if i == len(test.minos)-2 {
				mino1 = mino
			} else {
				mino2 = mino
			}
		}

		fullLines, holes, bumpy := board.boardStatsWithMinos(mino1, mino2)

		if fullLines != test.fullLines {
			mino1.SetOnBoard()
			lines := board.getDebugBoardWithMino(mino2)
			for i := 0; i < len(lines); i++ {
				t.Log(lines[i])
			}
			t.Errorf("AI fullLines - received: %v - expected: %v - info %v", fullLines, test.fullLines, test.info)
			continue
		}
		if holes != test.holes {
			mino1.SetOnBoard()
			lines := board.getDebugBoardWithMino(mino2)
			for i := 0; i < len(lines); i++ {
				t.Log(lines[i])
			}
			t.Errorf("AI holes - received: %v - expected: %v - info %v", holes, test.holes, test.info)
			continue
		}
		if bumpy != test.bumpy {
			mino1.SetOnBoard()
			lines := board.getDebugBoardWithMino(mino2)
			for i := 0; i < len(lines); i++ {
				t.Log(lines[i])
			}
			t.Errorf("AI bumpy - received: %v - expected: %v - info %v", bumpy, test.bumpy, test.info)
			continue
		}

	}
}
