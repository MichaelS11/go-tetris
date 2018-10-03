package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

// NewView creates a new view
func NewView() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()
	view = &View{}
}

// Stop stops the view
func (view *View) Stop() {
	logger.Println("View Stop start")

	termbox.Close()

	logger.Println("View Stop end")
}

// RefreshScreen refreshs the updated view to the screen
func (view *View) RefreshScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	view.drawBackground()
	view.drawTexts()

	if engine.previewBoard {
		board.DrawBoard()
		view.drawGameOver()
	} else if engine.gameOver {
		view.drawGameOver()
		view.drawRankingScores()
	} else if engine.paused {
		view.drawPaused()
	} else {
		board.DrawBoard()
		board.DrawPreviewMino()
		board.DrawDropMino()
		board.DrawCurrentMino()
	}

	termbox.Flush()
}

// drawBackground draws the background
func (view *View) drawBackground() {
	// playing board
	xOffset := boardXOffset
	yOffset := boardYOffset
	xEnd := boardXOffset + board.width*2 + 4
	yEnd := boardYOffset + board.height + 2
	for x := xOffset; x < xEnd; x++ {
		for y := yOffset; y < yEnd; y++ {
			if x == xOffset || x == xOffset+1 || x == xEnd-1 || x == xEnd-2 ||
				y == yOffset || y == yEnd-1 {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorWhite)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorBlack)
			}
		}
	}

	// piece preview
	xOffset = boardXOffset + board.width*2 + 8
	yOffset = boardYOffset
	xEnd = xOffset + 14
	yEnd = yOffset + 6
	for x := xOffset; x < xEnd; x++ {
		for y := yOffset; y < yEnd; y++ {
			if x == xOffset || x == xOffset+1 || x == xEnd-1 || x == xEnd-2 ||
				y == yOffset || y == yEnd-1 {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorWhite)
			} else {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorBlack)
			}
		}
	}

}

// drawTexts draws the text
func (view *View) drawTexts() {
	xOffset := boardXOffset + board.width*2 + 8
	yOffset := boardYOffset + 7

	view.drawText(xOffset, yOffset, "SCORE:", termbox.ColorWhite, termbox.ColorBlue)
	view.drawText(xOffset+7, yOffset, fmt.Sprintf("%7d", engine.score), termbox.ColorBlack, termbox.ColorWhite)

	yOffset += 2

	view.drawText(xOffset, yOffset, "LINES:", termbox.ColorWhite, termbox.ColorBlue)
	view.drawText(xOffset+7, yOffset, fmt.Sprintf("%7d", engine.deleteLines), termbox.ColorBlack, termbox.ColorWhite)

	yOffset += 2

	view.drawText(xOffset, yOffset, "LEVEL:", termbox.ColorWhite, termbox.ColorBlue)
	view.drawText(xOffset+7, yOffset, fmt.Sprintf("%4d", engine.level), termbox.ColorBlack, termbox.ColorWhite)

	yOffset += 2

	// ascii arrow characters add extra two spaces
	view.drawText(xOffset, yOffset, "←  - left", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "z    - rotate left", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "x    - rotate right", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "→  - right", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "↓  - soft drop", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "↑  - hard drop", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "sbar - hard drop", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "p    - pause", termbox.ColorWhite, termbox.ColorBlack)
	yOffset++
	view.drawText(xOffset, yOffset, "q    - quit", termbox.ColorWhite, termbox.ColorBlack)
}

// DrawPreviewMinoBlock draws the preview mino
func (view *View) DrawPreviewMinoBlock(x int, y int, color termbox.Attribute, rotation int, length int) {
	var char1 rune
	var char2 rune
	if rotation < 2 {
		char1 = '▓'
		char2 = ' '
	} else {
		char1 = ' '
		char2 = '▓'
	}
	xOffset := 2*x + 2*board.width + boardXOffset + 11 + (4 - length)
	termbox.SetCell(xOffset, y+boardYOffset+2, char1, color, color^termbox.AttrBold)
	termbox.SetCell(xOffset+1, y+boardYOffset+2, char2, color, color^termbox.AttrBold)
}

// DrawBlock draws a block
func (view *View) DrawBlock(x int, y int, color termbox.Attribute, rotation int) {
	var char1 rune
	var char2 rune
	if rotation < 2 {
		char1 = '▓'
		char2 = ' '
	} else {
		char1 = ' '
		char2 = '▓'
	}
	if color == blankColor {
		// blankColor means drop Mino
		termbox.SetCell(2*x+boardXOffset+2, y+boardYOffset+1, char1, termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite)
		termbox.SetCell(2*x+boardXOffset+3, y+boardYOffset+1, char2, termbox.ColorBlack|termbox.AttrBold, termbox.ColorWhite)
	} else {
		termbox.SetCell(2*x+boardXOffset+2, y+boardYOffset+1, char1, color, color^termbox.AttrBold)
		termbox.SetCell(2*x+boardXOffset+3, y+boardYOffset+1, char2, color, color^termbox.AttrBold)
	}
}

// drawPaused draws Paused
func (view *View) drawPaused() {
	yOffset := (board.height+1)/2 + boardYOffset
	view.drawTextCenter(yOffset, "Paused", termbox.ColorWhite, termbox.ColorBlack)
}

// drawGameOver draws GAME OVER
func (view *View) drawGameOver() {
	yOffset := boardYOffset + 2
	view.drawTextCenter(yOffset, " GAME OVER", termbox.ColorWhite, termbox.ColorBlack)
	yOffset += 2
	view.drawTextCenter(yOffset, "sbar for new game", termbox.ColorWhite, termbox.ColorBlack)

	if engine.previewBoard {
		return
	}

	yOffset += 2
	// ascii arrow characters add extra two spaces
	view.drawTextCenter(yOffset, "←previous board", termbox.ColorWhite, termbox.ColorBlack)
	yOffset += 2
	view.drawTextCenter(yOffset, "→next board", termbox.ColorWhite, termbox.ColorBlack)
}

// drawRankingScores draws the ranking scores
func (view *View) drawRankingScores() {
	yOffset := boardYOffset + 10
	for index, line := range engine.ranking.scores {
		view.drawTextCenter(yOffset+index, fmt.Sprintf("%1d: %6d", index+1, line), termbox.ColorWhite, termbox.ColorBlack)
	}
}

// drawText draws the provided text
func (view *View) drawText(x int, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for index, char := range text {
		termbox.SetCell(x+index, y, rune(char), fg, bg)
	}
}

// drawTextCenter draws text in the center of the board
func (view *View) drawTextCenter(y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	xOffset := board.width - (len(text)+1)/2 + boardXOffset + 2
	for index, char := range text {
		termbox.SetCell(index+xOffset, y, rune(char), fg, bg)
	}
}

// ShowDeleteAnimation draws the delete animation
func (view *View) ShowDeleteAnimation(lines []int) {
	view.RefreshScreen()

	for times := 0; times < 3; times++ {
		for _, y := range lines {
			view.colorizeLine(y, termbox.ColorCyan)
		}
		termbox.Flush()
		time.Sleep(140 * time.Millisecond)

		view.RefreshScreen()
		time.Sleep(140 * time.Millisecond)
	}
}

// ShowGameOverAnimation draws one randomily picked gave over animation
func (view *View) ShowGameOverAnimation() {
	logger.Println("View ShowGameOverAnimation start")

	switch rand.Intn(3) {
	case 0:
		for y := board.height - 1; y >= 0; y-- {
			view.colorizeLine(y, termbox.ColorBlack)
			termbox.Flush()
			time.Sleep(60 * time.Millisecond)
		}

	case 1:
		for y := 0; y < board.height; y++ {
			view.colorizeLine(y, termbox.ColorBlack)
			termbox.Flush()
			time.Sleep(60 * time.Millisecond)
		}

	case 2:
		sleepTime := 50 * time.Millisecond
		topStartX := boardXOffset + 3
		topEndX := board.width*2 + boardXOffset + 1
		topY := boardYOffset + 1
		rightStartY := boardYOffset + 1
		rightEndY := board.height + boardYOffset + 1
		rightX := board.width*2 + boardXOffset + 1
		bottomStartX := topEndX - 1
		bottomEndX := topStartX - 1
		bottomY := board.height + boardYOffset
		leftStartY := rightEndY - 1
		leftEndY := rightStartY - 1
		leftX := boardXOffset + 2

		for topStartX <= topEndX && rightStartY <= rightEndY {
			for x := topStartX; x < topEndX; x++ {
				termbox.SetCell(x, topY, ' ', termbox.ColorBlack, termbox.ColorBlack)
			}
			topStartX++
			topEndX--
			topY++
			for y := rightStartY; y < rightEndY; y++ {
				termbox.SetCell(rightX, y, ' ', termbox.ColorBlack, termbox.ColorBlack)
			}
			rightStartY++
			rightEndY--
			rightX--
			for x := bottomStartX; x > bottomEndX; x-- {
				termbox.SetCell(x, bottomY, ' ', termbox.ColorBlack, termbox.ColorBlack)
			}
			bottomStartX--
			bottomEndX++
			bottomY--
			for y := leftStartY; y > leftEndY; y-- {
				termbox.SetCell(leftX, y, ' ', termbox.ColorBlack, termbox.ColorBlack)
			}
			leftStartY--
			leftEndY++
			leftX++
			termbox.Flush()
			time.Sleep(sleepTime)
			sleepTime += 4 * time.Millisecond
		}
	}

	logger.Println("View ShowGameOverAnimation end")
}

// colorizeLine changes the color of a line
func (view *View) colorizeLine(y int, color termbox.Attribute) {
	for x := 0; x < board.width; x++ {
		termbox.SetCell(x*2+boardXOffset+2, y+boardYOffset+1, ' ', termbox.ColorDefault, color)
		termbox.SetCell(x*2+boardXOffset+3, y+boardYOffset+1, ' ', termbox.ColorDefault, color)
	}
}
