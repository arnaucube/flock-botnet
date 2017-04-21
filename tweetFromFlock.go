package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func postTweet(client *twitter.Client, text string) {
	tweet, httpResp, err := client.Statuses.Update(text, nil)
	if err != nil {
		fmt.Println(err)
	}
	if httpResp.Status != "200 OK" {
		c.Red("error: " + httpResp.Status)
		c.Purple("maybe twitter has blocked the account, CTRL+C, wait 15 minutes and try again")
	}
	fmt.Print("tweet posted: ")
	c.Green(tweet.Text)
}
func tweetFromFlock(flock Flock, text string) {
	fmt.Println("Starting to publish tweet: " + text)
	fmt.Println(strconv.Itoa(len(flock.Clients)))
	for i := 0; i < len(flock.Clients); i++ {
		fmt.Print("tweeting from: ")
		c.Cyan("@" + flock.ScreenNames[i])
		postTweet(flock.Clients[i], text)
		//waitTime(1)
	}

}

func optionTweetFromFlock(flock Flock) {

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
