package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const version = "0.1-dev"

func main() {
	c.Yellow("twFlock")
	fmt.Println("---------------")
	c.Cyan("twFlock initialized")
	c.Purple("https://github.com/arnaucode/twFlock")
	fmt.Println("version " + version)
	fmt.Println("Reading flockConfig.json file")
	//flock := readConfigTokensAndConnect()
	var flock Flock
	fmt.Println("---------------")
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number")
	options := `
	1 - Tweet from flock of bots
	2 - Markov
	0 - Exit script
option to select: `
	for {
		fmt.Print(options)

		option, _ := newcommand.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("selected 1 - Tweet from flock of bots")
			optionTweetFromFlock(flock)
			break
		case "2":
			fmt.Println("selected 2 - Markov")

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

			states := markov.train("the", "text.txt")
			generatedText := markov.generateText(states, firstWord, count)
			c.Green(generatedText)
			break
		case "0":
			fmt.Println("selected 0 - exit script")
			os.Exit(3)
			break
		default:
			fmt.Println("Invalid option")
			break
		}
	}
}
