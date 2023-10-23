package intermediate

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/sonnyb378/goroutineChallenges/pkg/customerror"
)

func WordCountInFile(dir string) ([]string, error) {
	// Challenge 2: Concurrent Word Count
	// Write a Go program that reads multiple text files concurrently, counts the number of
	// words in each file, and reports the results.

	// Input:
	// Three text files with content.

	// Desired Output:
	// File 1: 100 words
	// File 2: 85 words
	// File 3: 120 words

	wg := sync.WaitGroup{}

	fileContentChannel := make(chan map[string]int)

	openDir, _ := os.Open(dir)
	dirFiles, err := openDir.ReadDir(0)
	if err != nil {
		return []string{}, &customerror.CustomError{
			Message: fmt.Sprintf("%v", err),
		}
	}

	for _, txtFile := range dirFiles {
		timeStart := time.Now()

		fileName := txtFile.Name()
		fileBytes, err := os.ReadFile(dir + "/" + fileName)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", fileName, err)
			continue // Continue to the next file
		}

		wg.Add(1)
		go WordCount(&wg, fileContentChannel, fileBytes, fileName)

		fmt.Printf("Execution time for %v: %v (microseconds)\n", fileName, time.Since(timeStart).Microseconds())
	}

	go func() {
		wg.Wait()
		close(fileContentChannel)
	}()

	results := ReadChannel(fileContentChannel)

	if len(results) > 0 {
		return results, nil
	}

	return []string{}, &customerror.CustomError{
		Message: "No files found in given folder!",
	}
}

func ReadChannel(fileContentChannel chan map[string]int) []string {
	results := []string{}
	for content := range fileContentChannel {
		for k, v := range content {
			results = append(results, fmt.Sprintf("File \"%v\": %v words", k, v))
		}
	}
	return results
}

func WordCount(wg *sync.WaitGroup, fileContent chan map[string]int, fileBytes []byte, fileName string) {
	defer wg.Done()

	words := []string{}

	// exclude whitespace in match
	re := regexp.MustCompile(`[^\s]+\b`)

	if re.MatchString(string(fileBytes)) {
		words = re.FindAllString(string(fileBytes), -1)
	}

	data := make(map[string]int)
	data[fileName] = len(words)

	fileContent <- data
}
