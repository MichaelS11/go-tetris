package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// NewRanking create a new ranking
func NewRanking() *Ranking {
	ranking := &Ranking{
		scores: make([]uint64, 9),
	}

	if _, err := os.Stat(baseDir + rankingFileName); os.IsNotExist(err) {
		for i := 0; i < 9; i++ {
			ranking.scores[i] = 0
		}
		return ranking
	}

	scoreBytes, err := ioutil.ReadFile(baseDir + rankingFileName)
	if err != nil {
		logger.Println("NewRanking ReadFile error:", err)
	}

	scoreStrings := strings.Split(string(scoreBytes), ",")
	for index, scoreString := range scoreStrings {
		if index > 8 {
			break
		}
		score, err := strconv.ParseUint(scoreString, 10, 64)
		if err != nil {
			logger.Println("NewRanking ParseUint error:", err)
			score = 0
		}
		ranking.scores[index] = score
	}

	return ranking
}

// Save saves the rankings to a file
func (ranking *Ranking) Save() {
	var buffer bytes.Buffer

	for i := 0; i < 9; i++ {
		if i != 0 {
			buffer.WriteRune(',')
		}
		buffer.WriteString(strconv.FormatUint(ranking.scores[i], 10))
	}

	ioutil.WriteFile(baseDir+rankingFileName, buffer.Bytes(), 0644)
}

// InsertScore inserts a score into the rankings
func (ranking *Ranking) InsertScore(newScore uint64) {
	for index, score := range ranking.scores {
		if newScore > score {
			ranking.slideScores(index)
			ranking.scores[index] = newScore
			return
		}
	}
}

// slideScores slides the scores down to make room for a new score
func (ranking *Ranking) slideScores(index int) {
	for i := len(ranking.scores) - 1; i > index; i-- {
		ranking.scores[i] = ranking.scores[i-1]
	}
}
