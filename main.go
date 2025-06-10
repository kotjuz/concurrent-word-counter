package main

import (
	"bufio"
	"encoding/json"
	"flag"
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

	resultMap = make(map[string]int)

	wg sync.WaitGroup
	mu sync.Mutex
)

func main() {
	maxWorkers := flag.Int("workers", 5, "Number of concurrent files being processed")
	fileNumber := flag.Int("filenum", 10, "Number of files to generate")
	flag.Parse()

	semaphore := make(chan struct{}, *maxWorkers)
	err := GenerateFiles(wordList, *fileNumber)
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
		semaphore <- struct{}{}
		go func(path string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			processFile(path)
		}(filepath.Join(dirPath, file.Name()))

	}
	wg.Wait()

	err = saveToJson()

	if err != nil {
		fmt.Println(err)
	}

}

func processFile(filePath string) error {
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
