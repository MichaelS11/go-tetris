package main

import (
	"log"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	blankColor      = termbox.ColorBlack
	boardXOffset    = 4
	boardYOffset    = 2
	rankingFileName = "/tetris.db"

	// MinoPreview is for the preview mino
	MinoPreview MinoType = iota
	// MinoCurrent is for the current mino
	MinoCurrent = iota
	// MinoDrop is for the drop mino
	MinoDrop = iota
)

type (
	// MinoType is the type of mino
	MinoType int
	// MinoBlocks is the blocks of the mino
	MinoBlocks [][]termbox.Attribute
	// MinoRotation is the rotation of the mino
	MinoRotation [4]MinoBlocks

	// Mino is a mino
	Mino struct {
		x            int
		y            int
		length       int
		rotation     int
		minoRotation MinoRotation
	}

	// Minos is a bag of minos
	Minos struct {
		minoBag  [7]MinoRotation
		bagRand  []int
		bagIndex int
	}

	// Board is the Tetris board
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

	// Boards holds all the premade boards
	Boards struct {
		colors   [][]termbox.Attribute
		rotation [][]int
	}

	// KeyInput is the key input engine
	KeyInput struct {
		stopped      bool
		chanStop     chan struct{}
		chanKeyInput chan *termbox.Event
	}

	// View is the display engine
	View struct {
	}

	// Ai is the AI engine
	Ai struct {
		queue    *[]rune
		newQueue *[]rune
		index    int
	}

	// Ranking holds the ranking scores
	Ranking struct {
		scores []uint64
	}

	// Engine is the Tetirs game engine
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
	logger  *log.Logger
	minos   *Minos
	board   *Board
	view    *View
	engine  *Engine
)
