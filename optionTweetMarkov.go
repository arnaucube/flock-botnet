package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func optionTweetMarkov(states []State) {

	fmt.Print("entry the first word: ")
	newcommand := bufio.NewReader(os.Stdin)
	firstWord, _ := newcommand.ReadString('\n')
	firstWord = strings.TrimSpace(firstWord)
	fmt.Print("first word: ")
	c.Purple(firstWord)

	c.Red("how many words you want on the text?")
	newcommand = bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	fmt.Print("Number of words on text to generate: ")
	c.Purple(answer)
	count, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("incorrect entry, need a positive number")
	}

	c.Red("how many sentences you want to generate?")
	newcommand = bufio.NewReader(os.Stdin)
	answer, _ = newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	fmt.Print("Number of sentences to generate: ")
	c.Purple(answer)
	sentences, err := strconv.Atoi(answer)
	if err != nil {
		fmt.Println("incorrect entry, need a positive number")
	}

	fmt.Println("generating text")
	for i := 0; i < sentences; i++ {
		generatedText := markov.generateText(states, firstWord, count)
		c.Green(generatedText)
	}
}
