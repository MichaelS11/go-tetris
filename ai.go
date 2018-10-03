package main

import (
	"github.com/nsf/termbox-go"
)

// NewAi creates a new AI
func NewAi() *Ai {
	ai := Ai{}
	queue := make([]rune, 1)
	queue[0] = 'x'
	ai.queue = &queue
	return &ai
}

// ProcessQueue checks AI queue and process key moments
func (ai *Ai) ProcessQueue() {
	if ai.newQueue != nil {
		ai.queue = ai.newQueue
		ai.newQueue = nil
		ai.index = 0
	}
	queue := *ai.queue
	// wasd + qe keyboard keys
	switch queue[ai.index] {
	case 'w':
		board.MinoDrop()
	case 'a':
		board.MinoMoveLeft()
	case 'd':
		board.MinoMoveRight()
	case 'q':
		board.MinoRotateLeft()
	case 'e':
		board.MinoRotateRight()
	case 'x':
		return
	}
	ai.index++
	view.RefreshScreen()
}

// GetBestQueue gets the best queue
func (ai *Ai) GetBestQueue() {
	bestScore := -9999999
	bestQueue := make([]rune, 0, 0)
	currentMino := *board.currentMino
	rotations1 := 5
	rotations2 := 5

	switch currentMino.minoRotation[0][1][1] {
	case termbox.ColorCyan, termbox.ColorGreen, termbox.ColorRed:
		rotations1 = 2
	case termbox.ColorYellow:
		rotations1 = 1
	}
	switch board.previewMino.minoRotation[0][1][1] {
	case termbox.ColorCyan, termbox.ColorGreen, termbox.ColorRed:
		rotations2 = 2
	case termbox.ColorYellow:
		rotations2 = 1
	}

	for slide1 := 0; slide1 < 5; slide1++ {
		for move1 := board.width; move1 >= 0; move1-- {
			for rotate1 := 0; rotate1 < rotations1; rotate1++ {

				queue, mino1 := board.getMovesforMino(rotate1, move1, slide1, &currentMino, nil)
				if mino1 == nil {
					continue
				}

				for slide2 := 0; slide2 < 5; slide2++ {
					for move2 := board.width; move2 >= 0; move2-- {
						for rotate2 := 0; rotate2 < rotations2; rotate2++ {

							_, mino2 := board.getMovesforMino(rotate2, move2, slide2, board.previewMino, mino1)
							if mino2 == nil {
								continue
							}

							fullLines, holes, bumpy := board.boardStatsWithMinos(mino1, mino2)
							score := ai.getScoreFromBoardStats(fullLines, holes, bumpy)

							if score > bestScore {
								bestScore = score
								bestQueue = queue
							}

						}
					}
				}

			}
		}
	}

	if len(bestQueue) < 1 {
		bestQueue = append(bestQueue, 'x')
	}

	ai.newQueue = &bestQueue
}

func (board *Board) getMovesforMino(rotate int, move int, slide int, mino1 *Mino, mino2 *Mino) ([]rune, *Mino) {
	queue := make([]rune, 0, 4)
	mino := *mino1

	if rotate%2 == 0 {
		rotate /= 2
		for i := 0; i < rotate; i++ {
			mino.RotateRight()
			queue = append(queue, 'e')
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	} else {
		rotate = rotate/2 + 1
		for i := 0; i < rotate; i++ {
			mino.RotateLeft()
			queue = append(queue, 'q')
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	}

	if move%2 == 0 {
		move /= 2
		for i := 0; i < move; i++ {
			mino.MoveLeft()
			queue = append(queue, 'a')
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	} else {
		move = move/2 + 1
		for i := 0; i < move; i++ {
			mino.MoveRight()
			queue = append(queue, 'd')
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	}
	for mino.ValidLocation(false) && (mino2 == nil || !mino2.minoOverlap(&mino)) {
		mino.MoveDown()
	}
	mino.MoveUp()
	queue = append(queue, 'w')

	if slide%2 == 0 {
		slide /= 2
		for i := 0; i < slide; i++ {
			mino.MoveLeft()
			queue = append(queue, 'a')
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	} else {
		slide = slide/2 + 1
		for i := 0; i < slide; i++ {
			mino.MoveRight()
			queue = append(queue, 'd')
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	}

	if !mino.ValidLocation(true) {
		return queue, nil
	}
	queue = append(queue, 'x')
	return queue, &mino
}

func (board *Board) boardStatsWithMinos(mino1 *Mino, mino2 *Mino) (fullLines int, holes int, bumpy int) {
	// fullLines
	fullLinesY := make([]bool, board.height)
	for j := 0; j < board.height; j++ {
		fullLinesY[j] = true
		for i := 0; i < board.width; i++ {
			if board.colors[i][j] == blankColor && !mino1.isMinoAtLocation(i, j) && !mino2.isMinoAtLocation(i, j) {
				fullLinesY[j] = false
				break
			}
		}
		if fullLinesY[j] {
			fullLines++
		}
	}

	// holes and bumpy
	indexLast := 0
	for i := 0; i < board.width; i++ {
		index := board.height
		indexOffset := 0
		for j := 0; j < board.height; j++ {
			if fullLinesY[j] {
				indexOffset++
			} else {
				if board.colors[i][j] != blankColor || mino1.isMinoAtLocation(i, j) || mino2.isMinoAtLocation(i, j) {
					index = j
					break
				}
			}
		}

		if i != 0 {
			diffrence := (index + fullLines - indexOffset) - indexLast
			if diffrence < 0 {
				diffrence = -diffrence
			}
			bumpy += diffrence

		}
		indexLast = index + fullLines - indexOffset

		index++
		for j := index; j < board.height; j++ {
			if board.colors[i][j] == blankColor && !mino1.isMinoAtLocation(i, j) && !mino2.isMinoAtLocation(i, j) {
				holes++
			}
		}
	}
	return
}

func (ai *Ai) getScoreFromBoardStats(fullLines int, holes int, bumpy int) int {
	score := 0
	if fullLines == 4 {
		score += 512
	}
	score -= 80 * holes
	score -= 20 * bumpy
	return score
}
