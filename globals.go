package main

import (
	"log"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	blankColor       = termbox.ColorBlack
	boardXOffset     = 4
	boardYOffset     = 2
	aiTickDivider    = 8
	rankingFileName  = "/go-tetris.db"
	settingsFileName = "/go-tetris.json"

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
		fullLinesY   []bool
	}

	// Boards holds all the boards
	Boards struct {
		name     string
		colors   [][]termbox.Attribute
		rotation [][]int
	}

	// BoardsJSON is for JSON format of boards
	BoardsJSON struct {
		Name     string
		Mino     [][]string
		Rotation [][]int
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
		editMode     bool
	}

	// Edit is the board edit mode
	Edit struct {
		x         int
		y         int
		moved     bool
		boardSize bool
		width     int
		height    int
		saved     bool
	}

	// Settings is the JSON load/save file
	Settings struct {
		Boards []BoardsJSON
	}
)

var (
	baseDir string
	logger  *log.Logger
	minos   *Minos
	board   *Board
	view    *View
	engine  *Engine
	edit    *Edit

	boards            []Boards
	numInternalBoards int
)
