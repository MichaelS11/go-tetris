package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell"
)

var (
	screen tcell.Screen
	x      = 0
	y      = 1
)

func main() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Println("NewScreen error:", err)
	}
	err = screen.Init()
	if err != nil {
		fmt.Println("screen Init error:", err)
	}

	screen.Clear()

	for i := 0; i < 379; i++ {
		printNum(i)
		style := tcell.StyleDefault.Foreground(tcell.Color(i)).Background(tcell.Color(i)).Dim(true)
		screen.SetContent(x, y, '▄', nil, style)
		x++
		screen.SetContent(x, y, '▄', nil, style)
		x += 2
		if x > 80 {
			x = 0
			y += 2
		}
		if i == 15 {
			i = 255
		}
	}

	screen.Show()
}

func printNum(num int) {
	word := strconv.FormatInt(int64(num), 10) + ":"
	if num < 10 {
		word = "  " + word
	} else if num < 100 {
		word = " " + word
	}
	if len(word)+x+2 > 80 {
		x = 0
		y += 2
	}
	style := tcell.StyleDefault.Foreground(tcell.ColorLightGray).Background(tcell.ColorBlack)
	for _, char := range word {
		screen.SetContent(x, y, char, nil, style)
		x++
	}

}
