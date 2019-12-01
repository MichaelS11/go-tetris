package main

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
	bestQueue := []rune{'x'}
	currentMino := board.currentMino
	previewMino := board.previewMino
	rotations1 := 5
	rotations2 := 5
	slides := 5
	if board.width > 10 {
		slides = 3
	}

	switch currentMino.minoRotation[0][1][1] {
	case colorCyan, colorGreen, colorRed:
		rotations1 = 2
	case colorYellow:
		rotations1 = 1
	}
	switch previewMino.minoRotation[0][1][1] {
	case colorCyan, colorGreen, colorRed:
		rotations2 = 2
	case colorYellow:
		rotations2 = 1
	}

	for slide1 := 0; slide1 < slides; slide1++ {
		for move1 := board.width; move1 >= 0; move1-- {
			for rotate1 := 0; rotate1 < rotations1; rotate1++ {

				queue, mino1 := board.getMovesforMino(rotate1, move1, slide1, currentMino, nil)
				if mino1 == nil {
					continue
				}

				for slide2 := 0; slide2 < slides; slide2++ {
					for move2 := board.width; move2 >= 0; move2-- {
						for rotate2 := 0; rotate2 < rotations2; rotate2++ {

							_, mino2 := board.getMovesforMino(rotate2, move2, slide2, previewMino, mino1)
							if mino2 == nil {
								continue
							}

							fullLines, holes, bumpy, heightEnds := board.boardStatsWithMinos(mino1, mino2)
							score := ai.getScoreFromBoardStats(fullLines, holes, bumpy, heightEnds, mino1.y, mino2.y)

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

	ai.newQueue = &bestQueue
}

func (board *Board) getMovesforMino(rotate int, move int, slide int, mino1 *Mino, mino2 *Mino) ([]rune, *Mino) {
	var i int
	queue := make([]rune, 0, (rotate/2+1)+(move/2+1)+(slide/2+1)+1)
	mino := *mino1

	mino.MoveDown()

	if rotate%2 == 0 {
		rotate /= 2
		for i = 0; i < rotate; i++ {
			mino.RotateRight()
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
			queue = append(queue, 'e')
		}
	} else {
		rotate = rotate/2 + 1
		for i = 0; i < rotate; i++ {
			mino.RotateLeft()
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
			queue = append(queue, 'q')
		}
	}

	if move%2 == 0 {
		move /= 2
		for i = 0; i < move; i++ {
			mino.MoveRight()
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
			queue = append(queue, 'd')
		}
	} else {
		move = move/2 + 1
		for i = 0; i < move; i++ {
			mino.MoveLeft()
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
			queue = append(queue, 'a')
		}
	}
	for mino.ValidLocation(false) && (mino2 == nil || !mino2.minoOverlap(&mino)) {
		mino.MoveDown()
	}
	mino.MoveUp()
	queue = append(queue, 'w')

	if slide%2 == 0 {
		slide /= 2
		for i = 0; i < slide; i++ {
			mino.MoveLeft()
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
			queue = append(queue, 'a')
		}
	} else {
		slide = slide/2 + 1
		for i = 0; i < slide; i++ {
			mino.MoveRight()
			if !mino.ValidLocation(false) || (mino2 != nil && mino2.minoOverlap(&mino)) {
				return queue, nil
			}
			queue = append(queue, 'd')
		}
	}

	if !mino.ValidLocation(true) {
		return queue, nil
	}

	return append(queue, 'x'), &mino
}

func (board *Board) boardStatsWithMinos(mino1 *Mino, mino2 *Mino) (fullLines int, holes int, bumpy int, heightEnds int) {
	var i int
	var j int

	// fullLines
	for j = 0; j < board.height; j++ {
		board.fullLinesY[j] = true
		for i = 0; i < board.width; i++ {
			if board.colors[i][j] == colorBlank && !mino1.isMinoAtLocation(i, j) && !mino2.isMinoAtLocation(i, j) {
				board.fullLinesY[j] = false
				break
			}
		}
		if board.fullLinesY[j] {
			fullLines++
		}
	}

	// holes and bumpy
	var foundLast int
	var fullLinesFound int
	for i = 0; i < board.width; i++ {
		found := board.height
		fullLinesFound = 0
		for j = 0; j < board.height; j++ {
			if board.fullLinesY[j] {
				fullLinesFound++
			} else {
				if board.colors[i][j] != colorBlank || mino1.isMinoAtLocation(i, j) || mino2.isMinoAtLocation(i, j) {
					found = j
					break
				}
			}
		}

		if i == 0 {
			heightEnds = board.height - (found + fullLines - fullLinesFound)
		} else {
			diffrence := (found + fullLines - fullLinesFound) - foundLast
			if diffrence < 0 {
				diffrence = -diffrence
			}
			bumpy += diffrence
		}
		foundLast = found + fullLines - fullLinesFound

		for j++; j < board.height; j++ {
			if board.colors[i][j] == colorBlank && !mino1.isMinoAtLocation(i, j) && !mino2.isMinoAtLocation(i, j) {
				holes++
			}
		}
	}

	heightEnds += board.height - foundLast

	return
}

func (ai *Ai) getScoreFromBoardStats(fullLines int, holes int, bumpy int, heightEnds int, height1 int, height2 int) int {
	score := 8 * heightEnds
	if fullLines > 3 {
		score += 512
	}
	score -= 75 * holes
	score -= 15 * bumpy
	if height1 < 6 {
		score -= 10 * (5 - height1)
	}
	if height2 < 6 {
		score -= 10 * (5 - height2)
	}
	return score
}
