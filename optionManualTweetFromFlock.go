package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func optionManualTweetFromFlock(flock Flock) {

	fmt.Print("entry tweet content: ")
	newcommand := bufio.NewReader(os.Stdin)
	text, _ := newcommand.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Print("tweet content: ")
	c.Purple(text)

	c.Red("Are you sure? [y/n]")
	newcommand = bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		tweetFromFlock(flock, text)
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
