package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=======================")

	// iterate through files in directory /files
	textFiles, _ := ioutil.ReadDir("./files")
	for _, f := range textFiles {
		fileName := f.Name()

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
		fmt.Println(fileName + ": " + strconv.Itoa(len(contentWords)))
		file.Close()
	}

	fmt.Println("=======================")
}

func checkFile(e error) {
	if e != nil {
		panic(e)
	}
}
