package main

import (
	"math/rand"
	"time"
)

type MinoType int

const (
	MinoPreview MinoType = iota
	MinoCurrent          = iota
	MinoDrop             = iota
)

var (
	bagRand  []int
	bagIndex int
)

func init() {
	rand.Seed(time.Now().UnixNano())
	bagRand = rand.Perm(7)
}

type Mino struct {
	x            int
	y            int
	length       int
	rotation     int
	minoRotation MinoRotation
}

func NewMino() *Mino {
	minoRotation := minoBag[bagRand[bagIndex]]
	bagIndex++
	if bagIndex > 6 {
		bagIndex = 0
		bagRand = rand.Perm(7)
	}
	mino := &Mino{
		minoRotation: minoRotation,
		length:       len(minoRotation[0]),
	}
	mino.x = boardWidth/2 - (mino.length+1)/2
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
				if !ValidBlockLocation(mino.x+i, mino.y+j, mustBeOnBoard) {
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
