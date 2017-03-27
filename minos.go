package main

import (
	"github.com/nsf/termbox-go"
)

type MinoBlocks [][]termbox.Attribute

type MinoRotation [4]MinoBlocks

var minoBag [7]MinoRotation

func init() {
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
		minoRotationI[i] = initCloneRotateRight(minoRotationI[i-1])
	}
	var minoRotationJ MinoRotation
	minoRotationJ[0] = minoJ
	for i := 1; i < 4; i++ {
		minoRotationJ[i] = initCloneRotateRight(minoRotationJ[i-1])
	}
	var minoRotationL MinoRotation
	minoRotationL[0] = minoL
	for i := 1; i < 4; i++ {
		minoRotationL[i] = initCloneRotateRight(minoRotationL[i-1])
	}
	var minoRotationO MinoRotation
	minoRotationO[0] = minoO
	minoRotationO[1] = minoO
	minoRotationO[2] = minoO
	minoRotationO[3] = minoO
	var minoRotationS MinoRotation
	minoRotationS[0] = minoS
	for i := 1; i < 4; i++ {
		minoRotationS[i] = initCloneRotateRight(minoRotationS[i-1])
	}
	var minoRotationT MinoRotation
	minoRotationT[0] = minoT
	for i := 1; i < 4; i++ {
		minoRotationT[i] = initCloneRotateRight(minoRotationT[i-1])
	}
	var minoRotationZ MinoRotation
	minoRotationZ[0] = minoZ
	for i := 1; i < 4; i++ {
		minoRotationZ[i] = initCloneRotateRight(minoRotationZ[i-1])
	}

	minoBag = [7]MinoRotation{minoRotationI, minoRotationJ, minoRotationL, minoRotationO, minoRotationS, minoRotationT, minoRotationZ}
}

func initCloneRotateRight(minoBlocks MinoBlocks) MinoBlocks {
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
