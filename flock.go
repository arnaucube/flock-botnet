package main

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
)

func replyTweet(client *twitter.Client, text string, inReplyToStatusID int64) {
	tweet, httpResp, err := client.Statuses.Update(text, &twitter.StatusUpdateParams{
		InReplyToStatusID: inReplyToStatusID,
	})
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
