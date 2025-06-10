package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	wordList = []string{
		"apple", "banana", "orange", "grape", "peach",
		"dog", "cat", "mouse", "rabbit", "fox",
		"book", "pen", "notebook", "paper", "pencil",
	}
	maxWorkers = 5

	resultMap = make(map[string]int)

	wg        sync.WaitGroup
	mu        sync.Mutex
	semaphore chan struct{}
)

func main() {
	err := GenerateFiles(wordList)
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := os.ReadDir(dirPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		wg.Add(1)
		go processFile(filepath.Join(dirPath, file.Name()))
	}
	wg.Wait()

	err = saveToJson()

	if err != nil {
		fmt.Println(err)
	}

}

func processFile(filePath string) error {
	defer wg.Done()

	file, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()

		mu.Lock()
		resultMap[word] += 1
		mu.Unlock()
	}

	err = scanner.Err()
	return err

}

func saveToJson() error {
	file, err := os.Create("results.json")
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(resultMap)
	return err
}
