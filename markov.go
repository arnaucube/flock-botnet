package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Markov struct{}

/*type NextState struct {
	Word  string
	Count int
	Prob  float64
}*/
type State struct {
	Word       string
	Count      int
	Prob       float64
	NextStates []State
}

var markov Markov

func addWordToStates(states []State, word string) ([]State, int) {
	iState := -1
	for i := 0; i < len(states); i++ {
		if states[i].Word == word {
			iState = i
		}
	}

	if iState > 0 {
		states[iState].Count++
	} else {
		var tempState State
		tempState.Word = word
		tempState.Count = 1

		states = append(states, tempState)
		iState = len(states) - 1
	}
	//fmt.Println(iState)
	return states, iState
}

func countWordsProb(words []string) {
	var states []State
	totalWordsCount := 0
	//count words
	for i := 0; i < len(words); i++ {
		var iState int
		totalWordsCount++
		states, iState = addWordToStates(states, words[i])
		if iState != len(words)-1 {
			states[iState].NextStates, _ = addWordToStates(states[iState].NextStates, words[iState+1])
		}
	}

	//count prob
	for i := 0; i < len(states); i++ {
		states[i].Prob = (float64(states[i].Count) / float64(totalWordsCount) * 100)
		for j := 0; j < len(states[i].NextStates); j++ {
			states[i].NextStates[j].Prob = (float64(states[i].NextStates[j].Count) / float64(totalWordsCount) * 100)
		}
	}
	fmt.Println("totalWordsCount: " + strconv.Itoa(totalWordsCount))
	fmt.Println(states)
}

func textToWords(text string) []string {
	s := strings.Split(text, " ")
	words := s
	return words
}

func readText(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		//Do something
	}
	content := string(data)
	return content, err
}

func (markov Markov) train(firstWord string, path string) string {
	text, _ := readText(path)
	words := textToWords(text)
	countWordsProb(words)
	return text
}
