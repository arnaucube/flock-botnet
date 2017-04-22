package main

import (
	"io/ioutil"
	"strings"
)

func readTxt(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		//Do something
	}
	dataClean := strings.Replace(string(data), "\n", " ", -1)
	content := string(dataClean)
	return content, err
}
