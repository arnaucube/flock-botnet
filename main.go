package main

import (
	"bufio"
	"fmt"
	"os"
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
	flock := readConfigTokensAndConnect()
	fmt.Println("---------------")
	newcommand := bufio.NewReader(os.Stdin)
	fmt.Print("Please select command number")
	options := `
	1 - Tweet from flock of bots
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
