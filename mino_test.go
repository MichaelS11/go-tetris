package main

import (
	"log"
	"math/rand"
	"os"
	"testing"
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

	NewMinos()
	NewBoard()
	NewEngine()
}

func TestMinoValidLocation(t *testing.T) {
	// this must be set to the blank boards
	for _, i := range []int{0, 3} {
		board.boardsIndex = i
		board.Clear()

		tests := []struct {
			info           string
			mino           testMinoStruct
			changeLocation bool
			mustBeOnBoard  bool
			validLocation  bool
		}{
			{info: "start 0 false", mino: testMinoStruct{minoRotation: minos.minoBag[0]}, mustBeOnBoard: false, validLocation: true},
			{info: "start 0 true", mino: testMinoStruct{minoRotation: minos.minoBag[0]}, mustBeOnBoard: true, validLocation: true},
			{info: "start 1 false", mino: testMinoStruct{minoRotation: minos.minoBag[1]}, mustBeOnBoard: false, validLocation: true},
			{info: "start 1 true", mino: testMinoStruct{minoRotation: minos.minoBag[1]}, mustBeOnBoard: true, validLocation: false},
			{info: "start 2 false", mino: testMinoStruct{minoRotation: minos.minoBag[2]}, mustBeOnBoard: false, validLocation: true},
			{info: "start 2 true", mino: testMinoStruct{minoRotation: minos.minoBag[2]}, mustBeOnBoard: true, validLocation: false},
			{info: "start 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3]}, mustBeOnBoard: false, validLocation: true},
			{info: "start 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3]}, mustBeOnBoard: true, validLocation: false},
			{info: "start 4 false", mino: testMinoStruct{minoRotation: minos.minoBag[4]}, mustBeOnBoard: false, validLocation: true},
			{info: "start 4 true", mino: testMinoStruct{minoRotation: minos.minoBag[4]}, mustBeOnBoard: true, validLocation: false},
			{info: "start 5 false", mino: testMinoStruct{minoRotation: minos.minoBag[5]}, mustBeOnBoard: false, validLocation: true},
			{info: "start 5 true", mino: testMinoStruct{minoRotation: minos.minoBag[5]}, mustBeOnBoard: true, validLocation: false},
			{info: "start 6 false", mino: testMinoStruct{minoRotation: minos.minoBag[6]}, mustBeOnBoard: false, validLocation: true},
			{info: "start 6 true", mino: testMinoStruct{minoRotation: minos.minoBag[6]}, mustBeOnBoard: true, validLocation: false},

			{info: "top left 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: 0}, changeLocation: true, mustBeOnBoard: false, validLocation: true},
			{info: "top left 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: 0}, changeLocation: true, mustBeOnBoard: true, validLocation: true},
			{info: "top right 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: 0}, changeLocation: true, mustBeOnBoard: false, validLocation: true},
			{info: "top right 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: 0}, changeLocation: true, mustBeOnBoard: true, validLocation: true},
			{info: "bottom right 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: board.height - 2}, changeLocation: true, mustBeOnBoard: false, validLocation: true},
			{info: "bottom right 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: board.height - 2}, changeLocation: true, mustBeOnBoard: true, validLocation: true},
			{info: "bottom left 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: board.height - 2}, changeLocation: true, mustBeOnBoard: false, validLocation: true},
			{info: "bottom left 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: board.height - 2}, changeLocation: true, mustBeOnBoard: true, validLocation: true},

			{info: "up 1 top left 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: -1}, changeLocation: true, mustBeOnBoard: false, validLocation: true},
			{info: "up 1 top left 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: -1}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
			{info: "up 1 top right 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: -1}, changeLocation: true, mustBeOnBoard: false, validLocation: true},
			{info: "up 1 top right 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: -1}, changeLocation: true, mustBeOnBoard: true, validLocation: false},

			{info: "up 2 top left 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: -2}, changeLocation: true, mustBeOnBoard: false, validLocation: true},
			{info: "up 2 top left 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: -2}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
			{info: "up 2 top right 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: -2}, changeLocation: true, mustBeOnBoard: false, validLocation: true},
			{info: "up 2 top right 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: -2}, changeLocation: true, mustBeOnBoard: true, validLocation: false},

			{info: "up 3 top left 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: -3}, changeLocation: true, mustBeOnBoard: false, validLocation: false},
			{info: "up 3 top left 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: -3}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
			{info: "up 3 top right 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: -3}, changeLocation: true, mustBeOnBoard: false, validLocation: false},
			{info: "up 3 top right 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: -3}, changeLocation: true, mustBeOnBoard: true, validLocation: false},

			{info: "off top left 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: -1, y: 0}, changeLocation: true, mustBeOnBoard: false, validLocation: false},
			{info: "off top left 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: -1, y: 0}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
			{info: "off top right 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 1, y: 0}, changeLocation: true, mustBeOnBoard: false, validLocation: false},
			{info: "off top right 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 1, y: 0}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
			{info: "off bottom right 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 1, y: board.height - 2}, changeLocation: true, mustBeOnBoard: false, validLocation: false},
			{info: "off bottom right 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 1, y: board.height - 2}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
			{info: "off bottom right 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: board.height - 1}, changeLocation: true, mustBeOnBoard: false, validLocation: false},
			{info: "off bottom right 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: board.width - 2, y: board.height - 1}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
			{info: "off bottom left 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: -1, y: board.height - 2}, changeLocation: true, mustBeOnBoard: false, validLocation: false},
			{info: "off bottom left 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: -1, y: board.height - 2}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
			{info: "off bottom left 3 false", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: board.height - 1}, changeLocation: true, mustBeOnBoard: false, validLocation: false},
			{info: "off bottom left 3 true", mino: testMinoStruct{minoRotation: minos.minoBag[3], x: 0, y: board.height - 1}, changeLocation: true, mustBeOnBoard: true, validLocation: false},
		}

		for _, test := range tests {
			mino := NewMino()
			mino.minoRotation = test.mino.minoRotation
			mino.length = len(mino.minoRotation[0])
			if test.changeLocation {
				mino.x = test.mino.x
				mino.y = test.mino.y
			}
			validLocation := mino.ValidLocation(test.mustBeOnBoard)
			if validLocation != test.validLocation {
				lines := board.getDebugBoardWithMino(mino)
				for i := 0; i < len(lines); i++ {
					t.Log(lines[i])
				}
				t.Errorf("MinoValidLocation validLocation - received: %v - expected: %v - info %v", validLocation, test.validLocation, test.info)
			}
		}

	}

}
