package main

const (
	aiQueueSize  = (boardWidth+1)/2 + 6
	aiMoveLength = (boardWidth + 1) / 2
)

type Ai struct {
	queue [aiQueueSize]rune
	index int
}

func NewAi() *Ai {
	ai := Ai{}
	for i := 0; i < aiQueueSize; i++ {
		ai.queue[i] = 'x'
	}
	return &ai
}

func (ai *Ai) ProcessQueue() bool {
	// wasd + qe keyboard keys
	switch ai.queue[ai.index] {
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
		return false
	}
	ai.index++
	if ai.index == aiQueueSize {
		ai.index = 0
	}
	view.RefreshScreen()
	return true
}

func (ai *Ai) GetBestQueue() {
	ai.addMovesToQueue(ai.getBestQueue())
}

func (ai *Ai) addMovesToQueue(queue []rune) {
	insertIndex := ai.index
	for _, char := range queue {
		ai.queue[insertIndex] = char
		insertIndex++
		if insertIndex == aiQueueSize {
			insertIndex = 0
		}
	}
}

func (ai *Ai) getBestQueue() []rune {
	bestQueue := make([]rune, 0, 0)
	bestScore := -9999999
	var slideScore int
	bestSlide := 6

	for move1 := 0; move1 <= boardWidth; move1++ {
		for rotate1 := 0; rotate1 < 5; rotate1++ {
			for slide1 := 0; slide1 <= 5; slide1++ {

				queue, mino1 := ai.getMovesforMino(rotate1, move1, slide1, nil)
				if mino1 == nil {
					continue
				}

				for move2 := 0; move2 <= boardWidth; move2++ {
					for rotate2 := 0; rotate2 < 5; rotate2++ {
						for slide2 := 0; slide2 <= 5; slide2++ {

							_, mino2 := ai.getMovesforMino(rotate2, move2, slide2, mino1)
							if mino2 == nil {
								continue
							}

							fullLines, holes, bumpy := board.boardStatsWithMinos(mino1, mino2)
							score := ai.getScoreFromBoardStats(fullLines, holes, bumpy)

							if slide1 < 3 {
								slideScore = slide1
							} else {
								slideScore = slide1 - 2
							}

							if score > bestScore || (score == bestScore && slideScore < bestSlide) {
								bestScore = score
								bestQueue = queue
								bestSlide = slideScore
							}

						}
					}
				}

			}
		}
	}

	return bestQueue
}

func (ai *Ai) getMovesforMino(rotate int, move int, slide int, mino1 *Mino) ([]rune, *Mino) {
	queue := make([]rune, 0, 4)
	var mino Mino
	if mino1 != nil {
		mino = *board.previewMino
	} else {
		mino = *board.currentMino
	}
	if rotate < 3 {
		for i := 0; i < rotate; i++ {
			mino.RotateRight()
			queue = append(queue, 'e')
			if !mino.ValidLocation(false) || (mino1 != nil && mino1.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	} else {
		for i := 0; i < rotate-2; i++ {
			mino.RotateLeft()
			queue = append(queue, 'q')
			if !mino.ValidLocation(false) || (mino1 != nil && mino1.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	}
	if move <= aiMoveLength {
		move = aiMoveLength - move
		for i := 0; i < move; i++ {
			mino.MoveLeft()
			queue = append(queue, 'a')
			if !mino.ValidLocation(false) || (mino1 != nil && mino1.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	} else {
		move = move - aiMoveLength + 1
		for i := 0; i < move; i++ {
			mino.MoveRight()
			queue = append(queue, 'd')
			if !mino.ValidLocation(false) || (mino1 != nil && mino1.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	}
	for mino.ValidLocation(false) && (mino1 == nil || !mino1.minoOverlap(&mino)) {
		mino.MoveDown()
	}
	mino.MoveUp()
	queue = append(queue, 'w')
	if slide < 3 {
		for i := 0; i < slide; i++ {
			mino.MoveLeft()
			queue = append(queue, 'a')
			if !mino.ValidLocation(false) || (mino1 != nil && mino1.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	} else {
		slide = slide - 2
		for i := 0; i < slide; i++ {
			mino.MoveRight()
			queue = append(queue, 'd')
			if !mino.ValidLocation(false) || (mino1 != nil && mino1.minoOverlap(&mino)) {
				return queue, nil
			}
		}
	}
	queue = append(queue, 'x')
	return queue, &mino
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

func (board *Board) boardStatsWithMinos(mino1 *Mino, mino2 *Mino) (fullLines int, holes int, bumpy int) {
	// fullLines
	fullLinesY := make(map[int]bool, 2)
	for j := 0; j < boardHeight; j++ {
		isFullLine := true
		for i := 0; i < boardWidth; i++ {
			if board.colors[i][j] == blankColor && !mino1.isMinoAtLocation(i, j) && !mino2.isMinoAtLocation(i, j) {
				isFullLine = false
				break
			}
		}
		if isFullLine {
			fullLinesY[j] = true
			fullLines++
		}
	}

	// holes and bumpy
	indexLast := 0
	for i := 0; i < boardWidth; i++ {
		index := boardHeight
		indexOffset := 0
		for j := 0; j < boardHeight; j++ {
			if _, found := fullLinesY[j]; found {
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
		for j := index; j < boardHeight; j++ {
			if board.colors[i][j] == blankColor && !mino1.isMinoAtLocation(i, j) && !mino2.isMinoAtLocation(i, j) {
				holes++
			}
		}
	}
	return
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

func (ai *Ai) getScoreFromBoardStats(fullLines int, holes int, bumpy int) (score int) {
	if fullLines == 4 {
		score += 256
	}
	score -= 75 * holes
	score -= 25 * bumpy
	return score
}
