package main

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
  "strings"
  "sync"
  "time"
)

// Worker struct
type Worker struct {
  id int
}

// Start worker
func (worker Worker) Start(fileChannel chan string, waiter *sync.WaitGroup) {
  go func() {
    for {
      fmt.Printf("Worker #%d waiting ...\n", worker.id)
      select {
      case fileName := <-fileChannel:
        fmt.Printf("Worker #%d counting file "+fileName+"\n", worker.id)
        count := countWordsInFile(fileName)
        fmt.Println(fileName + ": " + strconv.Itoa(count))

      case <-time.After(5 * time.Second):
        fmt.Printf("Worker #%d timeout after 5 seconds of inactivity. Quiting ...\n", worker.id)
        waiter.Done()
        return
      }
    }
  }()
}

func countWordsInFile(fileName string) int {
  // read file
  file, err := os.Open("./files/" + fileName)
  checkFile(err)

  reader := bufio.NewReader(file)

  content, err := ioutil.ReadAll(reader)
  if err != nil {
    fmt.Println("error reading file " + fileName)
  }

  contentStr := string(content)
  words := strings.Fields(contentStr)

  file.Close()
  return len(words)
}
