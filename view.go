package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

const (
	boardXOffset = 4
	boardYOffset = 2
)

type View struct {
	drawDropMarkerDisabled bool
}

func NewView() *View {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()
	return &View{}
}

func (view *View) Stop() {
	logger.Info("View Stop start")

	termbox.Close()

	logger.Info("View Stop end")
}

func (view *View) RefreshScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	view.drawBackground()
	view.drawTexts()
	board.DrawPreviewMino()

	if engine.gameOver {
		view.drawGameOver()
	} else {
		board.DrawBoard()
		board.DrawDropMino()
		board.DrawCurrentMino()
	}

	termbox.Flush()
}

func (view *View) drawBackground() {
	// playing board
	xOffset := boardXOffset
	yOffset := boardYOffset
	xEnd := boardXOffset + boardWidth*2 + 4
	yEnd := boardYOffset + boardHeight + 2
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
	xOffset = boardXOffset + boardWidth*2 + 8
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

func (view *View) drawTexts() {
	xOffset := boardXOffset + boardWidth*2 + 8
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
	xOffset := 2*x + 2*boardWidth + boardXOffset + 11 + (4 - length)
	termbox.SetCell(xOffset, y+boardYOffset+2, char1, color, color^termbox.AttrBold)
	termbox.SetCell(xOffset+1, y+boardYOffset+2, char2, color, color^termbox.AttrBold)
}

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

func (view *View) drawGameOver() {
	xOffset := boardXOffset + 4
	yOffset := boardYOffset + 2

	view.drawText(xOffset, yOffset, "   GAME OVER", termbox.ColorWhite, termbox.ColorBlack)
	yOffset += 2
	view.drawText(xOffset, yOffset, "sbar for new game", termbox.ColorWhite, termbox.ColorBlack)
	yOffset += 2
	xOffset += 2
	for index, line := range engine.ranking.scores {
		view.drawText(xOffset, yOffset+index, fmt.Sprintf("%2d: %6d", index+1, line), termbox.ColorWhite, termbox.ColorBlack)
	}
}

func (view *View) drawText(x int, y int, text string, fg termbox.Attribute, bg termbox.Attribute) {
	for index, char := range text {
		termbox.SetCell(x+index, y, rune(char), fg, bg)
	}
}

func (view *View) ShowDeleteAnimation(lines []int) {
	view.drawDropMarkerDisabled = true

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

	view.drawDropMarkerDisabled = false
}

func (view *View) ShowGameOverAnimation() {
	view.drawDropMarkerDisabled = true

	view.RefreshScreen()

	for y := boardHeight - 1; y >= 0; y-- {
		view.colorizeLine(y, termbox.ColorBlack)
		termbox.Flush()
		time.Sleep(60 * time.Millisecond)
	}

	view.drawDropMarkerDisabled = false
}

func (view *View) colorizeLine(y int, color termbox.Attribute) {
	for x := 0; x < boardWidth; x++ {
		termbox.SetCell(x*2+boardXOffset+2, y+boardYOffset+1, ' ', termbox.ColorDefault, color)
		termbox.SetCell(x*2+boardXOffset+3, y+boardYOffset+1, ' ', termbox.ColorDefault, color)
	}
}
