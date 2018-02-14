package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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
		logger.Error("NewRanking ReadFile", "error", err.Error())
	}

	scoreStrings := strings.Split(string(scoreBytes), ",")
	for index, scoreString := range scoreStrings {
		if index > 8 {
			break
		}
		score, err := strconv.ParseUint(scoreString, 10, 64)
		if err != nil {
			logger.Error("NewRanking ParseUint", "error", err.Error())
			score = 0
		}
		ranking.scores[index] = score
	}

	return ranking
}

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

func (ranking *Ranking) InsertScore(newScore uint64) {
	for index, score := range ranking.scores {
		if newScore > score {
			ranking.slideScores(index)
			ranking.scores[index] = newScore
			return
		}
	}
}

func (ranking *Ranking) slideScores(index int) {
	for i := len(ranking.scores) - 1; i > index; i-- {
		ranking.scores[i] = ranking.scores[i-1]
	}
}
