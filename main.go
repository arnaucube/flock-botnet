package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const version = "0.1-dev"

func main() {
	c.Yellow("projectFlock")
	fmt.Println("---------------")
	c.Cyan("projectFlock initialized")
	c.Purple("https://github.com/arnaucode/projectFlock")
	fmt.Println("version " + version)
	fmt.Println("Reading flockConfig.json file")
	flock := readConfigTokensAndConnect()

	//var flock Flock
	fmt.Println("---------------")
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number")
	options := `
	1 - Manual tweet from flock of bots
	2 - Markov manual text generator
	3 - Markov's Flock Botnet
	0 - Exit script
option to select: `
	for {
		fmt.Print(options)

		option, _ := newcommand.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("selected 1 - Manual tweet from flock of bots")
			optionManualTweetFromFlock(flock)
			break
		case "2":
			fmt.Println("selected 2 - Markov manual text generator")
			c.Yellow("generating markov chains (may take some seconds)")
			text, _ := readTxt("text.txt")
			states := markov.train(text)
			c.Green("markov chains generated")

			optionTweetMarkov(states)
			break
		case "3":
			fmt.Println("selected 3 - Markov Flock Botnet")
			optionMarkovFlockBotnet(flock)
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
