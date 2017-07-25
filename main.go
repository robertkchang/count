package main

import (
  "fmt"
  "io/ioutil"
  "sync"
)

func main() {
  fmt.Println("=======================")

  fileChannel := make(chan string, 1)

  waiter := &sync.WaitGroup{}
  waiter.Add(10)

  for workerIdx := 0; workerIdx < 10; workerIdx++ {
    worker := Worker{id: workerIdx}

    fmt.Printf("Starting worker #%d\n", worker.id)
    worker.Start(fileChannel, waiter)
  }

  // iterate through files in directory /files
  textFiles, _ := ioutil.ReadDir("./files")
  for _, f := range textFiles {
    fileChannel <- f.Name()
  }

  waiter.Wait()
  fmt.Println("\n\nAll workers have shutdown!")

  fmt.Println("=======================")
}

func checkFile(e error) {
  if e != nil {
    panic(e)
  }
}
