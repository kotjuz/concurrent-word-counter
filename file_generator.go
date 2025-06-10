package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

const dirPath string = "generated_files"

func GenerateFiles(wordList []string) error {
	err := os.MkdirAll(dirPath, 0777)

	if err != nil {
		return errors.New("failed to create directory")
	}

	for i := 1; i <= 10; i++ {
		filename := filepath.Join(dirPath, fmt.Sprintf("file%d.txt", i))
		file, err := os.Create(filename)

		if err != nil {
			return errors.New("failed to create file")
		}

		wordsNumber := rand.Intn(100) + 50

		for j := 1; j <= wordsNumber; j++ {
			randomWord := wordList[rand.Intn(len(wordList))] + "\n"
			_, err = file.Write([]byte(randomWord))

			if err != nil {
				return errors.New("failed to write to file")
			}
		}

	}
	return nil
}
