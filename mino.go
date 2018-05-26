package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

func NewMino() *Mino {
	minoRotation := minos.minoBag[minos.bagRand[minos.bagIndex]]
	minos.bagIndex++
	if minos.bagIndex > 6 {
		minos.bagIndex = 0
		minos.bagRand = rand.Perm(7)
	}
	mino := &Mino{
		minoRotation: minoRotation,
		length:       len(minoRotation[0]),
	}
	mino.x = board.width/2 - (mino.length+1)/2
	mino.y = -1
	return mino
}

func (mino *Mino) CloneMoveLeft() *Mino {
	newMino := *mino
	newMino.MoveLeft()
	return &newMino
}

func (mino *Mino) MoveLeft() {
	mino.x--
}

func (mino *Mino) CloneMoveRight() *Mino {
	newMino := *mino
	newMino.MoveRight()
	return &newMino
}

func (mino *Mino) MoveRight() {
	mino.x++
}

func (mino *Mino) CloneRotateRight() *Mino {
	newMino := *mino
	newMino.RotateRight()
	return &newMino
}

func (mino *Mino) RotateRight() {
	mino.rotation++
	if mino.rotation > 3 {
		mino.rotation = 0
	}
}

func (mino *Mino) CloneRotateLeft() *Mino {
	newMino := *mino
	newMino.RotateLeft()
	return &newMino
}

func (mino *Mino) RotateLeft() {
	if mino.rotation < 1 {
		mino.rotation = 3
		return
	}
	mino.rotation--
}

func (mino *Mino) CloneMoveDown() *Mino {
	newMino := *mino
	newMino.MoveDown()
	return &newMino
}

func (mino *Mino) MoveDown() {
	mino.y++
}

func (mino *Mino) MoveUp() {
	mino.y--
}

func (mino *Mino) ValidLocation(mustBeOnBoard bool) bool {
	minoBlocks := mino.minoRotation[mino.rotation]
	for i := 0; i < mino.length; i++ {
		for j := 0; j < mino.length; j++ {
			if minoBlocks[i][j] != blankColor {
				if !board.ValidBlockLocation(mino.x+i, mino.y+j, mustBeOnBoard) {
					return false
				}
			}
		}
	}
	return true
}

func (mino *Mino) SetOnBoard() {
	minoBlocks := mino.minoRotation[mino.rotation]
	for i := 0; i < mino.length; i++ {
		for j := 0; j < mino.length; j++ {
			if minoBlocks[i][j] != blankColor {
				board.SetColor(mino.x+i, mino.y+j, minoBlocks[i][j], mino.rotation)
			}
		}
	}
}

func (mino *Mino) DrawMino(minoType MinoType) {
	minoBlocks := mino.minoRotation[mino.rotation]
	for i := 0; i < mino.length; i++ {
		for j := 0; j < mino.length; j++ {
			if minoBlocks[i][j] != blankColor {
				switch minoType {
				case MinoPreview:
					view.DrawPreviewMinoBlock(i, j, minoBlocks[i][j], mino.rotation, mino.length)
				case MinoDrop:
					view.DrawBlock(mino.x+i, mino.y+j, blankColor, mino.rotation)
				case MinoCurrent:
					if ValidDisplayLocation(mino.x+i, mino.y+j) {
						view.DrawBlock(mino.x+i, mino.y+j, minoBlocks[i][j], mino.rotation)
					}
				}
			}
		}
	}
}

func (mino *Mino) minoOverlap(mino1 *Mino) bool {
	minoBlocks := mino.minoRotation[mino.rotation]
	for i := 0; i < mino.length; i++ {
		for j := 0; j < mino.length; j++ {
			if minoBlocks[i][j] != blankColor {
				if mino1.isMinoAtLocation(mino.x+i, mino.y+j) {
					return true
				}
			}
		}
	}
	return false
}

func (mino *Mino) isMinoAtLocation(x int, y int) bool {
	xIndex := x - mino.x
	yIndex := y - mino.y
	if xIndex < 0 || xIndex >= mino.length || yIndex < 0 || yIndex >= mino.length {
		return false
	}

	minoBlocks := mino.minoRotation[mino.rotation]
	if minoBlocks[xIndex][yIndex] != blankColor {
		return true
	}

	return false
}

func (mino *Mino) getMinoColorAtLocation(x int, y int) termbox.Attribute {
	xIndex := x - mino.x
	yIndex := y - mino.y
	if xIndex < 0 || xIndex >= mino.length || yIndex < 0 || yIndex >= mino.length {
		return blankColor
	}

	minoBlocks := mino.minoRotation[mino.rotation]
	return minoBlocks[xIndex][yIndex]
}
