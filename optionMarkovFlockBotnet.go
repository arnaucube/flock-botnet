package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func isRT(tweet *twitter.Tweet) bool {
	tweetWords := strings.Split(tweet.Text, " ")
	for i := 0; i < len(tweetWords); i++ {
		if tweetWords[i] == "RT" {
			return true
		}
	}
	return false
}
func isFromBot(flock Flock, tweet *twitter.Tweet) bool {
	for i := 0; i < len(flock.ScreenNames); i++ {
		if flock.ScreenNames[i] == tweet.User.ScreenName {
			return true
		}
	}
	return false
}
func generateMarkovResponse(states []State, word string) string {
	generatedText := markov.generateText(states, word, 15)
	return generatedText
}
func processTweet(states []State, flockUser *twitter.Client, botScreenName string, keywords []string, tweet *twitter.Tweet) {
	c.Yellow("bot @" + botScreenName + " - New tweet detected:")
	fmt.Println(tweet.Text)

	tweetWords := strings.Split(tweet.Text, " ")
	generatedText := "word no exist on the memory"
	for i := 0; i < len(tweetWords) && generatedText == "word no exist on the memory"; i++ {
		fmt.Println(strconv.Itoa(i) + " - " + tweetWords[i])
		generatedText = generateMarkovResponse(states, tweetWords[i])
	}
	c.Yellow("bot @" + botScreenName + " posting response")
	fmt.Println(tweet.ID)
	replyTweet(flockUser, "@"+tweet.User.ScreenName+" "+generatedText, tweet.ID)
	waitTime(1)
}
func startStreaming(states []State, flock Flock, flockUser *twitter.Client, botScreenName string, keywords []string) {
	// Convenience Demux demultiplexed stream messages
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if isRT(tweet) == false && isFromBot(flock, tweet) == false {
			processTweet(states, flockUser, botScreenName, keywords, tweet)
		}
	}
	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}
	demux.Event = func(event *twitter.Event) {
		fmt.Printf("%#v\n", event)
	}

	fmt.Println("Starting Stream...")
	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         keywords,
		StallWarnings: twitter.Bool(true),
	}
	stream, err := flockUser.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}
	// Receive messages until stopped or stream quits
	demux.HandleChan(stream.Messages)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	/*ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()*/
}
func optionMarkovFlockBotnet(flock Flock) {
	c.Green("Starting Markov's Flock botnet")
	fmt.Println("generating Markov chains")
	inputText, _ := readTxt("text.txt")
	states := markov.train(inputText)

	//getting the keywords
	c.Purple("entry words to stream tweets (separated by comma): ")
	newcommand := bufio.NewReader(os.Stdin)
	text, _ := newcommand.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.Replace(text, " ", "", -1)
	keywords := strings.Split(text, ",")
	c.Purple("total keywords: " + strconv.Itoa(len(keywords)))
	fmt.Print("keywords to follow: ")
	fmt.Println(keywords)

	c.Red("Are you sure? [y/n]")
	newcommand = bufio.NewReader(os.Stdin)
	answer, _ := newcommand.ReadString('\n')
	answer = strings.TrimSpace(answer)
	switch answer {
	case "y":
		fmt.Println("ok, you are sure")
		for i := 0; i < len(flock.Clients); i++ {
			go startStreaming(states, flock, flock.Clients[i], flock.ScreenNames[i], keywords)
			waitSeconds(35)
		}
		break
	default:
		fmt.Println("Operation cancelled")
		break
	}
}
