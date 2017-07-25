package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Worker new
func Worker(channelQuit chan int, workQueue chan string) {
	select {
	case <-channelQuit:
		// add code to shutdown Go routine

	case fileName := <-workQueue:
		count := countWordsInFile(fileName)
		fmt.Println(fileName + ": " + strconv.Itoa(count))

	case <-time.After(5 * time.Second):
		fmt.Println("Worker timeout after 5 seconds of inactivity")
		channelQuit <- -1
		// waited 30 seconds and nothing happened
		// report this and exit
	}
}

func countWordsInFile(fileName string) int {
	// read file
	file, err := os.Open("./files/" + fileName)
	checkFile(err)

	reader := bufio.NewReader(file)

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println("error reading file " + fileName)
		// break
	}

	contentStr := string(content)
	re := regexp.MustCompile(`\r?\n`)
	contentStrSanNewline := strings.Replace(re.ReplaceAllString(contentStr, ":"), "::", " ", -1)
	contentWords := strings.Split(string(contentStrSanNewline), " ")

	file.Close()
	return len(contentWords)
}
