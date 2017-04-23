package main

import (
	"fmt"
	"strconv"
	"time"
)

func waitTime(minutes int) {
	//wait to avoid the twitter api limitation
	timeToSleep := time.Duration(minutes) * time.Minute
	fmt.Println("waiting " + strconv.Itoa(minutes) + " min to avoid twitter api limitation")
	fmt.Println(time.Now().Local())
	time.Sleep(timeToSleep)
}
func waitSeconds(seconds int) {
	//wait to avoid the twitter api limitation
	timeToSleep := time.Duration(seconds) * time.Second
	fmt.Println("waiting " + strconv.Itoa(seconds) + " seconds")
	fmt.Println(time.Now().Local())
	time.Sleep(timeToSleep)
}
