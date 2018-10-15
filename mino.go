package main

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

// NewMino creates a new Mino
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

// CloneMoveLeft creates copy of the mino and moves it left
func (mino *Mino) CloneMoveLeft() *Mino {
	newMino := *mino
	newMino.MoveLeft()
	return &newMino
}

// MoveLeft moves the mino left
func (mino *Mino) MoveLeft() {
	mino.x--
}

// CloneMoveRight creates copy of the mino and moves it right
func (mino *Mino) CloneMoveRight() *Mino {
	newMino := *mino
	newMino.MoveRight()
	return &newMino
}

// MoveRight moves the mino right
func (mino *Mino) MoveRight() {
	mino.x++
}

// CloneRotateRight creates copy of the mino and rotates it right
func (mino *Mino) CloneRotateRight() *Mino {
	newMino := *mino
	newMino.RotateRight()
	return &newMino
}

// RotateRight rotates the mino right
func (mino *Mino) RotateRight() {
	mino.rotation++
	if mino.rotation > 3 {
		mino.rotation = 0
	}
}

// CloneRotateLeft creates copy of the mino and rotates it left
func (mino *Mino) CloneRotateLeft() *Mino {
	newMino := *mino
	newMino.RotateLeft()
	return &newMino
}

// RotateLeft rotates the mino left
func (mino *Mino) RotateLeft() {
	if mino.rotation < 1 {
		mino.rotation = 3
		return
	}
	mino.rotation--
}

// CloneMoveDown creates copy of the mino and moves it down
func (mino *Mino) CloneMoveDown() *Mino {
	newMino := *mino
	newMino.MoveDown()
	return &newMino
}

// MoveDown moves the mino down
func (mino *Mino) MoveDown() {
	mino.y++
}

// MoveUp moves the mino up
func (mino *Mino) MoveUp() {
	mino.y--
}

// ValidLocation check if the mino is in a valid location
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

// SetOnBoard attaches mino to the board
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

// DrawMino draws the mino on the board
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

// minoOverlap check if a mino overlaps another mino
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

// isMinoAtLocation check if a mino block is in a location
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

// getMinoColorAtLocation gets the mino color at a location
func (mino *Mino) getMinoColorAtLocation(x int, y int) termbox.Attribute {
	xIndex := x - mino.x
	yIndex := y - mino.y
	if xIndex < 0 || xIndex >= mino.length || yIndex < 0 || yIndex >= mino.length {
		return blankColor
	}

	minoBlocks := mino.minoRotation[mino.rotation]
	return minoBlocks[xIndex][yIndex]
}
