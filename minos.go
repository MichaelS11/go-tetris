package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

// NewMinos creates the minos and minoBag
func NewMinos() {
	minoI := MinoBlocks{
		[]termbox.Attribute{blankColor, termbox.ColorCyan, blankColor, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorCyan, blankColor, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorCyan, blankColor, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorCyan, blankColor, blankColor},
	}
	minoJ := MinoBlocks{
		[]termbox.Attribute{termbox.ColorBlue, termbox.ColorBlue, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorBlue, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorBlue, blankColor},
	}
	minoL := MinoBlocks{
		[]termbox.Attribute{blankColor, termbox.ColorWhite, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorWhite, blankColor},
		[]termbox.Attribute{termbox.ColorWhite, termbox.ColorWhite, blankColor},
	}
	minoO := MinoBlocks{
		[]termbox.Attribute{termbox.ColorYellow, termbox.ColorYellow},
		[]termbox.Attribute{termbox.ColorYellow, termbox.ColorYellow},
	}
	minoS := MinoBlocks{
		[]termbox.Attribute{blankColor, termbox.ColorGreen, blankColor},
		[]termbox.Attribute{termbox.ColorGreen, termbox.ColorGreen, blankColor},
		[]termbox.Attribute{termbox.ColorGreen, blankColor, blankColor},
	}
	minoT := MinoBlocks{
		[]termbox.Attribute{blankColor, termbox.ColorMagenta, blankColor},
		[]termbox.Attribute{termbox.ColorMagenta, termbox.ColorMagenta, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorMagenta, blankColor},
	}
	minoZ := MinoBlocks{
		[]termbox.Attribute{termbox.ColorRed, blankColor, blankColor},
		[]termbox.Attribute{termbox.ColorRed, termbox.ColorRed, blankColor},
		[]termbox.Attribute{blankColor, termbox.ColorRed, blankColor},
	}

	var minoRotationI MinoRotation
	minoRotationI[0] = minoI
	for i := 1; i < 4; i++ {
		minoRotationI[i] = minosCloneRotateRight(minoRotationI[i-1])
	}
	var minoRotationJ MinoRotation
	minoRotationJ[0] = minoJ
	for i := 1; i < 4; i++ {
		minoRotationJ[i] = minosCloneRotateRight(minoRotationJ[i-1])
	}
	var minoRotationL MinoRotation
	minoRotationL[0] = minoL
	for i := 1; i < 4; i++ {
		minoRotationL[i] = minosCloneRotateRight(minoRotationL[i-1])
	}
	var minoRotationO MinoRotation
	minoRotationO[0] = minoO
	minoRotationO[1] = minoO
	minoRotationO[2] = minoO
	minoRotationO[3] = minoO
	var minoRotationS MinoRotation
	minoRotationS[0] = minoS
	for i := 1; i < 4; i++ {
		minoRotationS[i] = minosCloneRotateRight(minoRotationS[i-1])
	}
	var minoRotationT MinoRotation
	minoRotationT[0] = minoT
	for i := 1; i < 4; i++ {
		minoRotationT[i] = minosCloneRotateRight(minoRotationT[i-1])
	}
	var minoRotationZ MinoRotation
	minoRotationZ[0] = minoZ
	for i := 1; i < 4; i++ {
		minoRotationZ[i] = minosCloneRotateRight(minoRotationZ[i-1])
	}

	minos = &Minos{
		minoBag: [7]MinoRotation{minoRotationI, minoRotationJ, minoRotationL, minoRotationO, minoRotationS, minoRotationT, minoRotationZ},
		bagRand: rand.Perm(7),
	}
}

// minosCloneRotateRight clones a mino and rotates the mino to the right
func minosCloneRotateRight(minoBlocks MinoBlocks) MinoBlocks {
	length := len(minoBlocks)
	newMinoBlocks := make(MinoBlocks, length, length)
	for i := 0; i < length; i++ {
		newMinoBlocks[i] = make([]termbox.Attribute, length, length)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			newMinoBlocks[length-j-1][i] = minoBlocks[i][j]
		}
	}

	return newMinoBlocks
}
