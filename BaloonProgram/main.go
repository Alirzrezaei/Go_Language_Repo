package main

import (
	"fmt"
	"os"
	"strings"
)

type Baloon struct {
	position int
	height   int
}

const MY_FILE = "input 1.txt"

func fileReadingToString() []string {
	//data, error = os.ReadFile(MY_FILE)
	var myData, myError = os.ReadFile(MY_FILE)
	if myError != nil {
		fmt.Println("Error in reading file!")
		panic(myError)
	}
	return strings.Split(string(myData), "\n")
}

func commandchecker(baloon *Baloon, command string, value int) {
	switch command {
	case "forward":
		baloon.position += value
	case "up":
		baloon.height += value
	case "down":
		baloon.height -= value
	default:
		fmt.Println("Not valid command provided!")
	}
}

func main() {
	var stringData = fileReadingToString()

	for _, line := range stringData {
		if len(line) != 0 {
			fmt.Println("line: ", line)
			continue
		}
	}

}
