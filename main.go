package main

import "fmt"

var wordList = []string{
	"apple", "banana", "orange", "grape", "peach",
	"dog", "cat", "mouse", "rabbit", "fox",
	"book", "pen", "notebook", "paper", "pencil",
}

func main() {
	err := GenerateFiles(wordList)
	if err != nil {
		fmt.Println(err)
		return
	}
}
