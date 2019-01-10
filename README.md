# Go Tetris

Golang Tetris for console window with optional AI

## Features include

- AI (use i key to toggle)
- Lock delay
- Next piece
- Ghost piece
- Top scores
- Board choices
- Edit boards

## Compile

```
go get github.com/MichaelS11/go-tetris
go install github.com/MichaelS11/go-tetris
```

## Play

Then run the binary created, go-tetris or go-tetris.exe

## Keys start screen

| Key | Action |
| --- | --- |
| &larr; | previous board |
| &rarr; | next board |
| spacebar | start game |
| ctrl e | edit board |
| q | quit |

## Keys during game

| Key | Action |
| --- | --- |
| &larr; | left move |
| &rarr; | right move |
| &darr; | soft drop |
| &uarr; | hard drop |
| spacebar | hard drop |
| z | left rotate |
| x | right rotate |
| p | pause |
| q | quit |
| i | toggle AI |

## Keys edit mode

| Key | Action |
| --- | --- |
| &larr; | move cursor left |
| &rarr; | move cursor right |
| &darr; | move cursor down |
| &uarr; | move cursor up |
| z | rotate left |
| x | rotate right |
| c | cyan block - I |
| b | blue block - J |
| w | white block - L |
| e | yellow block - O |
| g | green block - S |
| a | magenta block - T |
| r | red block - Z |
| f | free block |
| ctrl b | change board size |
| ctrl s | save board |
| ctrl n | save board as new |
| ctrl k | delete board |
| ctrl o | empty board |
| ctrl q | quit edit mode |

## Screenshots

![alt text](https://raw.githubusercontent.com/MichaelS11/tetris/master/screenshots/tetris.png "Go Tetris")

![alt text](https://raw.githubusercontent.com/MichaelS11/tetris/master/screenshots/heart.png "Golang Tetris Heart")

![alt text](https://raw.githubusercontent.com/MichaelS11/tetris/master/screenshots/editmode.png "Edit Mode Peace Symbol")

![alt text](https://raw.githubusercontent.com/MichaelS11/tetris/master/screenshots/highscores.png "Tetris High Scores")

## To do

* Improve AI speed (slow on large boards)
