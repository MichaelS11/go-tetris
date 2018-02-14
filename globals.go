package main

import (
	"time"

	"github.com/nsf/termbox-go"
	"gopkg.in/inconshreveable/log15.v2"
)

const (
	blankColor      = termbox.ColorBlack
	boardXOffset    = 4
	boardYOffset    = 2
	rankingFileName = "/tetris.db"

	MinoPreview MinoType = iota
	MinoCurrent          = iota
	MinoDrop             = iota
)

type (
	MinoType     int
	MinoBlocks   [][]termbox.Attribute
	MinoRotation [4]MinoBlocks

	Mino struct {
		x            int
		y            int
		length       int
		rotation     int
		minoRotation MinoRotation
	}

	Minos struct {
		minoBag  [7]MinoRotation
		bagRand  []int
		bagIndex int
	}

	Board struct {
		boardsIndex  int
		width        int
		height       int
		colors       [][]termbox.Attribute
		rotation     [][]int
		previewMino  *Mino
		currentMino  *Mino
		dropDistance int
	}

	Boards struct {
		colors   [][]termbox.Attribute
		rotation [][]int
	}

	KeyInput struct {
		stopped      bool
		chanStop     chan struct{}
		chanKeyInput chan *termbox.Event
	}

	View struct {
	}

	Ai struct {
		queue    *[]rune
		newQueue *[]rune
		index    int
	}

	Ranking struct {
		scores []uint64
	}

	Engine struct {
		stopped      bool
		chanStop     chan struct{}
		keyInput     *KeyInput
		ranking      *Ranking
		timer        *time.Timer
		tickTime     time.Duration
		paused       bool
		gameOver     bool
		previewBoard bool
		score        int
		level        int
		deleteLines  int
		ai           *Ai
		aiEnabled    bool
		aiTimer      *time.Timer
	}
)

var (
	boards []Boards

	baseDir string
	logger  log15.Logger
	minos   *Minos
	board   *Board
	view    *View
	engine  *Engine
)
